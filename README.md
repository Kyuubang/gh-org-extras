# 🚀 gh-org-extras

[![GitHub license](https://img.shields.io/github/license/Kyuubang/gh-org-extras)](https://github.com/Kyuubang/gh-org-extras/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/Kyuubang/gh-org-extras.svg)](https://github.com/Kyuubang/gh-org-extras/releases/)
[![Go Version](https://img.shields.io/github/go-mod/go-version/Kyuubang/gh-org-extras)](https://github.com/Kyuubang/gh-org-extras)

> A powerful extension for GitHub CLI that enhances your GitHub organization management capabilities

`gh-org-extras` extends the standard `gh org` command with additional features that make managing GitHub organizations, teams, and members much easier - especially at scale!

## ✨ Features

- 👥 **Member Management** - List, invite, and remove organization members with ease
- 🔄 **Bulk Operations** - Perform actions on multiple users or teams at once
- 👪 **Team Management** - Create, update, and organize your GitHub teams
- 📋 **Detailed Listing** - View comprehensive information about members and teams
- 🔍 **Pagination Support** - Efficiently handle large organizations
- 🔒 **Dry-Run Mode** - Preview changes before applying them

## 📦 Installation

### Prerequisites

- GitHub CLI - [Installation Guide](https://github.com/cli/cli?tab=readme-ov-file#installation)
- Requires GitHub CLI version 2.0.0 or higher

### Installing gh-org-extras

```bash
gh extension install kyuubang/gh-org-extras
```

### Updating

```bash
gh extension upgrade kyuubang/gh-org-extras
```

### Uninstalling

```bash
gh extension remove kyuubang/gh-org-extras
```

## 🔧 Usage

### Basic Commands

```bash
gh org-extras [command]
```

### Member Management

List organization members:
```bash
gh org-extras member list <organization-name>
```

Invite members:
```bash
gh org-extras member invite <organization-name> <username>
```

### Team Management

List teams:
```bash
gh org-extras team list <organization-name>
```

### Bulk remove members from a team

```bash
gh org-extras bulk remove <organization_name> -t <team_slug>
```

## 🔐 Authentication

This extension requires a GitHub token with appropriate permissions:

- For listing members/teams: `read:org` scope
- For managing members/teams: `admin:org` scope

If you encounter permission issues, run:
```bash
gh auth login --scopes admin:org
```

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues and pull requests.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
