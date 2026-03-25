# Security Policy

## Supported Versions

This project currently supports the latest version available on the `main` branch.

Older versions may not receive security updates.

---

## Reporting a Vulnerability

If you discover a security vulnerability, please report it responsibly.

* Do NOT open a public issue
* Contact the maintainer directly

You can report issues via:

* GitHub private security advisory (preferred)
* Email (if available)

---

## What to Include

Please include as much detail as possible:

* Description of the vulnerability
* Steps to reproduce
* Potential impact
* Suggested fix (if any)

---

## Response Time

The maintainer will try to respond as soon as possible, but no strict SLA is guaranteed.

---

## Scope

This library is a data structure implementation and does not handle:

* Networking
* Cryptography
* External input validation

However, issues like:

* race conditions
* unsafe concurrent access
* panics

are considered relevant and should be reported.

---

## Disclosure Policy

* Vulnerabilities will be fixed before public disclosure
* A changelog entry will mention security fixes when applicable

---

Thank you for helping make this project more secure.
