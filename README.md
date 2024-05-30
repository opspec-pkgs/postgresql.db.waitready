[![Build Status](https://github.com/opspec-pkgs/postgresql.db.waitready/workflows/build/badge.svg?branch=main)](https://github.com/opspec-pkgs/postgresql.db.waitready/actions?query=workflow%3Abuild+branch%3Amain)

<img src="icon.svg" alt="icon" height="100px">

# Problem statement

Pings a PostgreSQL database until it responds or fails by timing out. Useful to know when a DB has finished launching.

# Example usage

## Visualize

```shell
opctl ui github.com/opspec-pkgs/postgresql.db.waitready#1.0.0
```

## Run

```
opctl run github.com/opspec-pkgs/postgresql.db.waitready#1.0.0
```

## Compose

```yaml
op:
  ref: github.com/opspec-pkgs/postgresql.db.waitready#1.0.0
  inputs:
    dbPassword:  # 👈 required; provide a value
    dbSchema:  # 👈 required; provide a value
    dbUsername:  # 👈 required; provide a value
  ## uncomment to override defaults
  #   dbHostname: "postgresql-db"
  #   dbPort: 5432
  #   timeout: 60
```

# Support

join us on
[![Slack](https://img.shields.io/badge/slack-opctl-E01563.svg)](https://join.slack.com/t/opctl/shared_invite/zt-51zodvjn-Ul_UXfkhqYLWZPQTvNPp5w)
or
[open an issue](https://github.com/opspec-pkgs/postgresql.db.waitready/issues)

# Releases

releases are versioned according to
[![semver 2.0.0](https://img.shields.io/badge/semver-2.0.0-brightgreen.svg)](http://semver.org/spec/v2.0.0.html)
and [tagged](https://git-scm.com/book/en/v2/Git-Basics-Tagging); see
[CHANGELOG.md](CHANGELOG.md) for release notes

# Contributing

see
[project/CONTRIBUTING.md](https://github.com/opspec-pkgs/project/blob/main/CONTRIBUTING.md)
