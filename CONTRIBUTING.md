# Contributing to Autumn Go

Thank you for your interest in contributing to Autumn Go! This is the unofficial Go SDK for Autumn's subscription and usage tracking API.

## How to contribute

First, set up the project locally using the installation guide below

Create a new branch for your changes
```bash
git checkout -b feature/async-client
# or
git checkout -b fix/error-handling
```

Commit your changes with clear, descriptive messages using our format:
```bash
git commit -m "feat(client): add async client support"
git commit -m "feat(types): add checkout response model"
git commit -m "fix(client): handle timeout errors properly"
git commit -m "feat(autumn): add custom retry logic"
```

Push to your fork
```bash
git push origin your-branch-name
```

Submit a Pull Request
- Go to your fork on GitHub and click "New Pull Request"
- Fill out the PR template completely
- Include tests for new features

## Installation Guide

### Requirements
- Go 1.19+
- Git

### Quickstart

**Step 1: Fork and Clone**
- Click the 'Fork' button at the top right of this repository
- Clone your fork locally: `git clone https://github.com/YOUR-USERNAME/autumn-go.git`

**Step 2: Install dependencies**
```bash
cd autumn-go
go mod tidy
```

**Step 3: Run tests**
```bash
go test ./...
```

### Local Development

**Format code:**
```bash
go fmt ./...
```

**Check for issues:**
```bash
go vet ./...
```

**Run examples:**
```bash
cd examples
go run main.go
```