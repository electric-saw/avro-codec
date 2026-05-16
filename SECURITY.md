# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 2.x     | :white_check_mark: |
| < 2.0   | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please report it responsibly.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report them via email to: **mariobaborsavalente@gmail.com**

You should receive a response within 48 hours. If for some reason you do not, please follow up via email to ensure we received your original message.

### What to Include

Please include the following information in your report:

- Type of vulnerability (e.g., buffer overflow, SQL injection, cross-site scripting)
- Full paths of source file(s) related to the vulnerability
- Location of the affected source code (tag/branch/commit or direct URL)
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### What to Expect

- **Acknowledgment**: We will acknowledge your report within 48 hours
- **Communication**: We will keep you informed of the progress towards a fix
- **Disclosure**: We will coordinate with you on the disclosure timeline
- **Credit**: We will credit you in the security advisory (unless you prefer to remain anonymous)

## Security Best Practices

When using this library, please keep in mind:

### Input Validation

The library has a default `MaxByteSliceSize` of 1 MiB to prevent memory exhaustion attacks from untrusted input. If you're processing untrusted Avro data, ensure this limit is appropriate for your use case:

```go
cfg := avro.Config{
    MaxByteSliceSize: 1024 * 1024, // 1 MiB (default)
}
```

### Schema Validation

Always validate schemas before using them in production:

```go
schema, err := avro.Parse(schemaJSON)
if err != nil {
    // Handle invalid schema
}
```

### Dependency Updates

Keep your dependencies up to date. We use Dependabot to help maintain our dependencies, and we recommend you do the same.

## Security Updates

Security updates will be released as patch versions and announced through:

- GitHub Security Advisories
- Release notes

We recommend watching this repository for releases to stay informed about security updates.
