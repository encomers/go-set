# Contributing to go-set

Thank you for your interest in contributing to **go-set**. Contributions are welcome and appreciated.

## How to Contribute

### 1. Fork the Repository

Create your own fork of the repository and clone it locally:

```bash
git clone https://github.com/encomers/go-set.git
cd go-set
```

### 2. Create a Branch

Always create a new branch for your changes:

```bash
git checkout -b feature/your-feature-name
```

### 3. Make Changes

* Keep the code simple and readable
* Follow existing code style and structure
* Avoid unnecessary abstractions
* Ensure no external dependencies are introduced

### 4. Write Tests

* Add tests for any new functionality
* Update existing tests if behavior changes
* Make sure all tests pass:

```bash
go test ./...
```

### 5. Run Benchmarks (Optional but Recommended)

If your changes affect performance, include benchmark results:

```bash
go test -bench=. -benchmem
```

### 6. Commit Your Changes

Use clear and descriptive commit messages:

```bash
git commit -m "Add: new feature X"
git commit -m "Fix: issue with Y"
```

### 7. Open a Pull Request

* Describe what you changed and why
* Include benchmarks if relevant
* Keep PRs focused and minimal

---

## Code Guidelines

* Use Go 1.23+ features where appropriate
* Prefer idiomatic Go
* Avoid breaking public API unless absolutely necessary
* Maintain backward compatibility when possible

---

## Project Principles

* **No external dependencies**
* **High performance**
* **Minimal allocations**
* **Simple and predictable API**

---

## Reporting Issues

If you find a bug or want to suggest a feature:

1. Check existing issues first
2. Provide a clear description
3. Include reproducible examples if possible

---

## Questions

If you have questions, feel free to open an issue.

---

Thank you for contributing 🚀
