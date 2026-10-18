"""Policy tests and real-Git integration tests; no model or GitHub doubles."""
import copy
import importlib.util
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

spec = importlib.util.spec_from_file_location('cycle', Path(__file__).parents[1] / 'cycle.py')
cycle = importlib.util.module_from_spec(spec)
spec.loader.exec_module(cycle)


class ReviewPolicyTests(unittest.TestCase):
    def setUp(self):
        self.report = dict(head='a' * 40, squint='Complete Squint review including all units.',
                           magic_tricks='Complete test audit with all case verdicts.', unresolved=[], blockers=[])
        self.pr = dict(state='open', draft=False,
                       head=dict(sha='a' * 40, ref='feature', repo=dict(full_name='owner/repo')),
                       base=dict(ref='main'))

    def test_accepts_complete_matching_review(self):
        self.assertEqual(cycle.validate_report(self.report, 'a' * 40), self.report)

    def test_rejects_missing_skill_or_wrong_commit(self):
        for field in self.report:
            report = copy.deepcopy(self.report)
            del report[field]
            with self.subTest(field=field), self.assertRaises(cycle.Blocked):
                cycle.validate_report(report, 'a' * 40)
        with self.assertRaises(cycle.Blocked):
            cycle.validate_report(self.report, 'b' * 40)

    def test_rejects_invalid_blocker_and_review_types(self):
        for key, value in [('blockers', 'none'), ('unresolved', [None]), ('squint', ''), ('magic_tricks', False)]:
            report = dict(self.report, **{key: value})
            with self.subTest(key=key), self.assertRaises(cycle.Blocked):
                cycle.validate_report(report, 'a' * 40)

    def test_full_reviews_survive_chunking(self):
        self.report['squint'] = 'Full detailed review.\n' * 6000
        body = cycle.review_body(self.report)
        parts = cycle.chunks(body)
        self.assertEqual(''.join(parts), body)
        self.assertTrue(all(len(part) <= 50000 for part in parts))
        self.assertIn(self.report['magic_tricks'], body)

    def test_open_same_repository_head_is_accepted(self):
        self.assertEqual(cycle.validate_pr(self.pr, 'owner/repo', 'a' * 40), self.pr)

    def test_stale_closed_draft_fork_and_base_branch_are_rejected(self):
        variants = []
        for key, value in [('state', 'closed'), ('draft', True)]:
            variants.append(dict(self.pr, **{key: value}))
        fork = copy.deepcopy(self.pr)
        fork['head']['repo']['full_name'] = 'stranger/repo'
        variants.append(fork)
        base = copy.deepcopy(self.pr)
        base['head']['ref'] = 'main'
        variants.append(base)
        stale = copy.deepcopy(self.pr)
        stale['head']['sha'] = 'b' * 40
        variants.append(stale)
        for pr in variants:
            with self.subTest(pr=pr), self.assertRaises(cycle.Blocked):
                cycle.validate_pr(pr, 'owner/repo', 'a' * 40)

    def test_control_files_and_environment_variants_are_protected(self):
        for path in ['.github/workflows/a.yml', 'sub/.env.local', 'AGENTS.md',
                     '.agents/skills/x', '.codex/config.toml', '.git/config']:
            with self.subTest(path=path):
                self.assertTrue(cycle.protected(path))
        self.assertFalse(cycle.protected('src/feature.py'))


class PatchIntegrationTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.workspace = Path(self.temp.name)
        self.repo = self.workspace / 'target'
        self.repo.mkdir()
        (self.workspace / 'result').mkdir()
        self.git('init', '-q')
        (self.repo / 'code.py').write_text('value = 1\n')
        self.git('add', '.')
        self.git('-c', 'user.name=Test', '-c', 'user.email=test@example.com', 'commit', '-qm', 'base')
        self.head = self.git('rev-parse', 'HEAD')

    def git(self, *args):
        return subprocess.check_output(['git', *args], cwd=self.repo, text=True).strip()

    def run_stage(self, stage, head=None):
        return subprocess.run(
            [sys.executable, str(Path(cycle.__file__).resolve()), stage],
            cwd=self.workspace, env=dict(os.environ, HEAD=head or self.head),
            text=True, capture_output=True,
        )

    def test_regular_staged_code_change_is_accepted(self):
        (self.repo / 'code.py').write_text('value = 2\n')
        self.git('add', '.')
        cycle.validate_changes(self.repo, self.head)

    def test_new_control_file_is_rejected(self):
        (self.repo / 'AGENTS.md').write_text('Change instructions\n')
        self.git('add', '.')
        with self.assertRaises(cycle.Blocked):
            cycle.validate_changes(self.repo, self.head)

    def test_symlink_is_rejected(self):
        (self.repo / 'link').symlink_to('code.py')
        self.git('add', '.')
        with self.assertRaises(cycle.Blocked):
            cycle.validate_changes(self.repo, self.head)

    def test_real_patch_round_trip_preserves_new_file(self):
        (self.repo / 'new.py').write_text('answer = 42\n')
        result = self.run_stage('export')
        self.assertEqual(result.returncode, 0, result.stderr)
        self.git('reset', '--hard', self.head)
        result = self.run_stage('apply')
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual((self.repo / 'new.py').read_text(), 'answer = 42\n')
        self.assertEqual(self.git('show', ':new.py'), 'answer = 42')

    def test_cli_patch_round_trip_preserves_final_blank_context(self):
        (self.repo / 'code.py').write_text('value = 1\n\n')
        self.git('add', '.')
        self.git('-c', 'user.name=Test', '-c', 'user.email=test@example.com',
                 'commit', '-qm', 'blank context')
        self.head = self.git('rev-parse', 'HEAD')
        (self.repo / 'code.py').write_text('value = 2\n\n')
        result = self.run_stage('export')
        self.assertEqual(result.returncode, 0, result.stderr)
        self.git('reset', '--hard', self.head)
        result = self.run_stage('apply')
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual((self.repo / 'code.py').read_bytes(), b'value = 2\n\n')

    def test_cli_export_rejects_changed_head_without_emitting_patch(self):
        result = self.run_stage('export', head='0' * 40)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn('Fixer changed HEAD', result.stderr)
        self.assertFalse((self.workspace / 'result/fix.patch').exists())

    def test_cli_apply_rejects_protected_patch(self):
        (self.repo / 'AGENTS.md').write_text('Change instructions\n')
        self.git('add', '.')
        patch = subprocess.check_output(['git', 'diff', '--binary', self.head], cwd=self.repo)
        (self.workspace / 'result/fix.patch').write_bytes(patch)
        self.git('reset', '--hard', self.head)
        result = self.run_stage('apply')
        self.assertNotEqual(result.returncode, 0)
        self.assertIn('protected control/secret files', result.stderr)

    def test_cli_export_and_apply_reject_both_sides_of_protected_renames(self):
        for source, destination in [
            ('AGENTS.md', 'ordinary.md'),
            ('.github/workflows/example.yml', 'ordinary.yml'),
            ('code.py', 'sub/AGENTS.md'),
        ]:
            with self.subTest(source=source, destination=destination):
                source_path = self.repo / source
                source_path.parent.mkdir(parents=True, exist_ok=True)
                source_path.write_text('original content\n')
                self.git('add', '.')
                self.git('-c', 'user.name=Test', '-c', 'user.email=test@example.com',
                         'commit', '-qm', 'rename baseline')
                self.head = self.git('rev-parse', 'HEAD')
                (self.repo / destination).parent.mkdir(parents=True, exist_ok=True)
                self.git('mv', '-f', source, destination)
                patch = subprocess.check_output(
                    ['git', 'diff', '--binary', self.head], cwd=self.repo)
                result = self.run_stage('export')
                self.assertNotEqual(result.returncode, 0)
                self.assertIn('protected control/secret files', result.stderr)
                self.git('reset', '--hard', self.head)
                (self.workspace / 'result/fix.patch').write_bytes(patch)
                result = self.run_stage('apply')
                self.assertNotEqual(result.returncode, 0)
                self.assertIn('protected control/secret files', result.stderr)
                self.git('reset', '--hard', self.head)


if __name__ == '__main__':
    unittest.main()
