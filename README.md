# Envforcer

Requires that environment variables must be set.

## Overview

`envforcer` returns a non-zero exit code if it fails to validate the env.

If any required environment variables are not set, `envforcer` returns an error
(a non-zero exit code).

```
$ envforcer FOO BAR && echo "Hello, $FOO $BAR"
env FOO is not set
env BAR is not set
```

However, if required environment variables are set correctly, `envforcer`
returns successfully (zero exit code).

```
$ export FOO="Tiny"
$ export BAR="Dancer"
$ envforcer FOO BAR && echo "Hello, $FOO $BAR"
Hello, Tiny Dancer
```

## Installation

Install from source via `go get github.com/nathanleiby/envforcer`, or download a release on the [releases](https://github.com/Clever/envforcer/releases) page.

## Usage

Specify the required env via a file and/or command line arguments.

**(1) command line arguments.**

This verifies that `FOO` and `BAR` are both set.

```bash
envforcer FOO BAR 
```

**(2) file (YAML/JSON)**

This verifies that `FOO` and `BAR` are both set.

```bash
envforcer -file=required.yml
```

```yaml
# required.yml
env:
    - FOO
    - BAR
```

**(3) file (YAML/JSON) and command line arguments**

This verifies that `FOO`, `BAR`, and `BAZ` are set.

```bash
envforcer -file=required.yml BAZ
```

## Testing

```
make test
```

## Rolling an official release

The release process requires a cross-compilation toolchain.
[`gox`](https://github.com/mitchellh/gox) can install the toolchain with one command: `gox -build-toolchain`.
From there you can build release tarballs for different OS and architecture combinations with `make release`.

Official releases are listed on the [releases](https://github.com/nathanleiby/envforcer/releases) page.

