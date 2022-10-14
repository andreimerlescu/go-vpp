# VPP

Version Plus Plus [![Build & Test VPP](https://github.com/andreimerlescu/go-vpp/actions/workflows/go.yaml/badge.svg)](https://github.com/andreimerlescu/go-vpp/actions/workflows/go.yaml)

## About

The `vpp` utility is small. It's designed to accept up to 2 CLI arguments but only 1 is required. 
The first argument is responsible for what the __CURRENT VERSION__ is of your project and when you run
`vpp CURRENT_VERSION` the `STDOUT` will be the next PATCH release. Otherwise, if the next intended
release is a major release, then run `vpp CURRENT_VERSION major`. 

## Installation

This package is intended to be installed on your system and should be accessible via your path.

```bash
go install github.com/andreimerlescu/go-vpp@latest
```

## Usage

Ensure the `vpp` package is in your path before using. 

```text
./vpp VERSION [bumpType]
         ^        ^ 
         |        |
         |        --- acceptable values "major", "minor" or "patch" (default)
         |
         ------------- the CURRENT version that you wish to bump
```

| Acceptable Format   | Example Value | Command     | Output  | 
|---------------------|---------------|-------------|---------|
| `Major`             | `1`           | `vpp 1`     | `1.0.1` | 
| `Major.Minor`       | `1.1`         | `vpp 1.1`   | `1.1.1` |
| `Major.Minor.Patch` | `1.1.1`       | `vpp 1.1.1` | `1.1.2` |

## License

This project uses the [MIT License](LICENSE).

## Contributors

1. [@andreimerlescu](https://github.com/andreimerlescu)

