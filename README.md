# Project Description

## Overview

Simple Task is a multi-language collection of small programming challenges and reference implementations. The repository is organized around standalone tasks that exercise algorithmic thinking, data transformation, parsing, scheduling, search, and basic state-machine logic. Each task is described in the main README with an example input and expected output, and implemented in one or more programming languages.

The project is useful as a practice workspace for solving coding prompts, comparing implementations across languages, and gradually expanding a catalog of reusable solutions.

## Purpose

The main goal of this project is to keep programming exercises simple, readable, and easy to run. Instead of building one large application, the repository stores many independent tasks. This makes it practical to:

- Practice problem solving in small, focused increments.
- Add the same solution in several languages.
- Compare language syntax, data structures, and standard-library choices.
- Keep completed and planned tasks visible in one central task list.
- Use examples as quick smoke tests for each implementation.

## Repository Structure

The repository is grouped by language:

- `golang/` contains Go implementations, usually with each task in its own `task_##/main.go` directory.
- `python/` contains Python implementations as individual `task_##.py` files.
- `ruby/` contains Ruby implementations for tasks that have been started in Ruby.
- `README.md` contains the full task catalog, examples, and links to available implementations.

Some README links are marked as `TBD`, which indicates that a task prompt exists but an implementation for that language has not yet been added.

## Task Style

Most tasks are intentionally compact and self-contained. They usually include:

- A short natural-language problem statement.
- One or more examples.
- A function-oriented solution.
- A simple executable entry point or print statement for quick manual verification.

The tasks cover a broad range of beginner-to-intermediate programming topics, including:

- String and character processing.
- Array and list transformations.
- Sorting and grouping.
- Greedy scheduling.
- Grid traversal.
- State-machine validation.
- Parsing structured text such as CSV.
- Mathematical sequences and number classification.
- Search and combinatorics.

## Current Implementation Coverage

The repository currently has the strongest coverage in Go and Python. Go implementations are present for many of the earlier tasks and selected later tasks. Python implementations cover a substantial subset of the catalog. Ruby support has started with an implementation for the mechanical keyboard assembly task.

The README acts as the source of truth for which tasks are available in each language and which are still planned.

## How to Use the Project

To explore the project, start with `README.md` and choose a task from the list. Each task entry links to the implementation files that exist for that prompt.

Typical workflows include:

- Reading a prompt and attempting a solution before opening the existing implementation.
- Running an implementation directly from its language folder.
- Adding a missing implementation for a `TBD` language entry.
- Improving an existing solution while preserving the example behavior.
- Adding clearer examples or edge cases to the README.

## Contribution Guidelines

When adding or updating a task, keep the change small and consistent with the existing structure:

- Use the next appropriate task number from the README.
- Place the solution in the matching language directory.
- Keep each task independent from unrelated tasks.
- Include a direct example or minimal executable output where practical.
- Update README links when a new implementation is added.
- Prefer clear, readable code over overly clever shortcuts.

If a task has different reasonable interpretations, document the chosen behavior near the solution or in the README so future implementations can stay consistent.

## Project Status

Simple Task is an evolving practice repository. The task catalog already includes many prompts, but implementation coverage is intentionally incomplete. That makes the project a good place to continue adding language support, refining existing solutions, and experimenting with different approaches to the same problem.
