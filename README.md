# 🛠️ Pralex - A Flutter Git Workflow CLI

Pralex is a custom command-line interface written in Go to simplify common tasks in Flutter development and Git workflows. Designed to boost productivity, it wraps common `git`, `fvm`, and `vscode` commands into intuitive CLI operations.

---

## ✨ Features

- 🚀 `pralex init` – Initialize a new Flutter project with FVM
- 🔁 `pralex set <version>` – Switch Flutter SDK using FVM
- 📦 `pralex clone <repo_url> [branch]` – Clone Git repositories easily
- 🔀 `pralex change [-b] <branch>` – Checkout or create new Git branches
- ⬆️ `pralex push [branch]` – Push commits (with optional target branch)
- 💻 `pralex open` – Open current directory in VS Code

---

## ⚙️ Installation

### ✅ Go 1.21+ required

You can install the CLI globally using:

```bash
go install github.com/PrathameshKalekar/flutter-cli@latest
