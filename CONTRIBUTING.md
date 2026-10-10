# WoW Habits Contributing Guide <!-- omit from toc -->

Thanks for taking the time to contribute to WoW Habits!

This is a small project maintained by one person. Issues and pull requests are
welcome.

- [Contributing to WoW Habits](#contributing-to-wow-habits)
  - [Prepare your environment](#prepare-your-environment)
  - [Fork and clone WoW Habits](#fork-and-clone-wow-habits)
  - [Using the Makefile](#using-the-makefile)
  - [Install required tools](#install-required-tools)
  - [Building binaries](#building-binaries)
- [Architecture](#architecture)
- [Codebase structure](#codebase-structure)
  - [Go Format and lint](#go-format-and-lint)
  - [PR Approval and Merging](#pr-approval-and-merging)

## Contributing to WoW Habits

This section describes how to make a contribution to WoW Habits.
These instructions are geared towards using a Linux development machine, which is required for doing development on the WoW Habits backend.

### Prepare your environment

Install [Go](https://go.dev/doc/install) 1.21 or later. The exact version used by the
project is set in `go.mod` and downloaded automatically by the `go` command.

### Fork and clone WoW Habits

First, you need to fork this project on GitHub.
Then clone your fork locally:
```shell
$ git clone git@github.com:<you>/wowhabits
$ cd ./wowhabits/
```

### Using the Makefile

WoW Habits uses a Makefile for common actions such as compiling WoW Habits and linting.

You can list available actions by using:
```shell
$ make help

Usage:
  make <target>

...output...
```

### Install required tools

Makefile allow you to install needed development tools (e.g. the linter):
```shell
$ make install-all
```

### Building binaries

To build WoW Habits binaries, you can run `make build-all`.
Built binaries will be placed in the `bin/` directory.
You can manually test to verify that WoW Habits is working by running the binaries.

## Architecture

Description about architecture in the repository is found [here](./docs/ARCHITECTURE.md).

## Codebase structure

Description about important directories in the repository is found [here](./docs/CODE_STRUCTURE.md).

### Go Format and lint

All code changes must pass `make lint-all`.

### PR Approval and Merging

The maintainer reviews every PR and merges it once CI passes and review comments are addressed.
