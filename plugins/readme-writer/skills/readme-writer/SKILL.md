---
name: readme-writer
description: Write or rewrite a project README for human readers with little time. Use when the user asks to create, improve, shorten, or review a README. Focus on value, installation, first use, and brief optional contribution guidance. Do not use for specifications, agent instructions, or implementation documentation.
---

# README writer

Describe the value, explain the user's part, and leave the implementation to other documents.

Write for a human who might click away at any moment. They want to know what the project does, whether it solves their problem, and how to get started.

## Voice

Use the qualities the user associates with DHH: direct, opinionated, conversational, and confident about the project's choices. Explain its philosophy in terms of what the user gets. Use plain words, concrete benefits, and short paragraphs. Skip marketing claims, corporate language, and cleverness for its own sake. Don't manufacture opinions or claims the project doesn't support.

## Find the user's path

Read the existing README and enough project material to establish what the product does, its intended users, the supported installation path, and normal first use. Source code and internal documents can establish facts without becoming README content.

Use the project's actual commands, interface labels, prerequisites, and examples. Don't invent features, installation steps, compatibility claims, or contribution policies. Ask a focused question when a missing fact prevents correct instructions. Don't run installations or change the product merely to write its README.

Choose one primary reader: the person using the product. For a developer library, that person may be a programmer, but still needs no knowledge of the project's internals. Explain necessary unfamiliar terms where they first appear.

## Structure

Use these sections in this order. Omit contribution guidance when it doesn't help. Keep required notices, such as license or attribution text, without expanding them into unrelated documentation.

### Project name

Start with `# <Project name>` and a very high-level introduction. Say what the project is, what problem it solves, and its general philosophy. Show a concrete benefit or recognizable situation. Lead with the user's outcome. Do not open with frameworks, architecture, protocols, or a list of components.

### How to install it

Give just enough instructions to get the project installed through its normal supported path. Include only prerequisites and setup choices that the reader must act on. Put steps in execution order and make commands usable as written, explaining any values the reader must replace. Avoid a catalog of alternative setups when one recommended path is enough.

### How to use it

Briefly explain how to interact with the product. Start with the first useful action and give a few realistic examples that introduce common tasks. Show what the reader should expect when that helps them recognize success. Include details only when they help someone get acquainted with using it correctly. Link to deeper usage documentation only where the reader has a reason to follow it.

### How to contribute

Keep this brief and welcoming, or omit it. Explain the simplest useful way to report a problem or contribute. Link to existing contributor guidance for longer setup or process instructions. Don't turn contribution into an obstacle course.

## Content rules

- Include only information the reader needs to understand the value, install the product, or use it correctly. A necessary user action earns space; the difficulty of implementing a feature does not.
- Never include implementation details, development history, internal decision records, agent instructions, review rules, test evidence, or audit reports. Put those in appropriate separate documents when requested, rather than expanding the assignment by default.
- Do not mix user, maintainer, reviewer, and agent audiences. Keep internal processes out when they require no user action. Do not make users repeat instructions the product already supplies to agents.
- Do not require users to manage internals. When a file location, setting, or storage behavior requires a user decision, explain only their action and its consequence, without teaching the underlying machinery.
- Select better content before polishing sentences. Simpler wording does not rescue an unnecessary paragraph. Do not add material to demonstrate how much work went into the product.

## Final edit

Check every paragraph against both questions:

1. "Does the reader who might be about to click away from this page really need to know this information?"
2. "Does the reader need this to use the product correctly?"

The introduction earns its place by helping the reader decide whether the product solves their problem. Installation and usage earn theirs by getting the reader to a useful result. Optional contribution guidance earns its place by making a contribution easy.

Cut anything that cannot justify its place. Check that a newcomer can follow the supported installation and first-use path without guessing, and that examples match the actual product. Keep consequential limitations the reader needs for correct use. Brevity must not make the instructions incomplete.

When asked to write or rewrite, deliver the finished README or edit the requested file. When asked only to review, identify the specific content to cut or change without silently rewriting it. Keep any explanation of your edits outside the README.
