# README writer

Write a README that helps someone decide whether a project is useful and get started. It leads with the value, keeps installation and usage brief, and leaves implementation details to other documents.

The tone is direct, opinionated, and conversational. Every paragraph has to earn the reader's time.

## How to install it

In Claude Code, run:

```
/plugin marketplace add KellyBennett/ClaudeSkills
/plugin install readme-writer@kbennett-skills
```

If you've already added the marketplace, update it with `/plugin marketplace update kbennett-skills` before installing.

## How to use it

Ask it to write a README, rewrite one that's too dense, or review one for unnecessary detail:

```
/readme-writer Write a README for this project.
/readme-writer Rewrite README.md for someone trying the app for the first time.
/readme-writer Review README.md. What can we cut?
```

READMEs follow a simple order: project name and introduction, installation, usage with realistic examples, and brief contribution guidance when useful.
