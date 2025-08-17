<img src="docs/attachments/verso_logo_banner.svg" />

# verso.cli

Simple [CHANGELOG](https://keepachangelog.com) [semver](https://semver.org/) extractor CLI. If you'd like to use verso as a Go library module, have a look at the [hkionline/verso](https://github.com/hkionline/verso) -project.

## Installation

There are couple of ways to install Verso to your system. The project is working on adding more altrenative, easy ways, to install Verso. Below are the currently supported installation methods.

### Using Go Install

You can install Verso easily if you have Go installed in your system. Please make sure you have your GOPATH environmental variable set and exported as well. If you have these requirements fullfilled you can install Verso CLI with the following command:

```bash
go install github.com/hkionline/verso.cli/cmd/verso@latest
```

After Installation you should now have the verso command available from your favorite terminal emulator.

### Downloading A Release Binary

The project makes multiple kinds of binaries available to download directly from the releases page. Navigate to releases, choose the latest release or any of the older releases and dowload the compressed directory of a Verso-release. The directory contains the lisence, changelog and the binary itself. You may manually move the binary to any location you want Verso CLI to run.Make sure to add it to your PATH for ease of use.

*Note* some operating systems, notably MacOS is by default suspicious of any downloaded binary and they are marked as "quarantined". The binary can be "un-quaranined", just follow the instructions in [docs](docs/apple-macos-binary-quarantine.md).

## Usage

Verso CLI expects to find a CHANGELOG-file from the working directory you are in. Alternatively you can pipe a CHANGELOG-file from a path.

### Latest version

Get the latest version in the CHANGELOG.
```bash
> verso latest

0.3.0
```

### List all version

Get list of versions found from the CHANGELOG. Each version will be in its own line.
```bash
> verso list

0.3.0
0.2.5
0.2.4
0.2.3
0.2.2
0.2.1
0.2.0
0.1.0
```

### Bump version

Bump (increment) the latest found version. Assume the latest version is 0.3.0.
```bash
> verso bump major

1.0.0

> verso bump minor

0.4.0

> verso bump patch

0.3.1
```

### Piping a changelog to verso

Pipe a CHANGELOG-file from a path using e.g. cat
```bash
> cat CHANGELOG.md | verso latest 

0.3.0
```

### Bonus: Changelog ready version header

Use verso and sed to produce CHANGELOG.md ready version header. Assume the latest version is 0.3.0.
```bash
> verso bump minor | sed "s/\(.*\)/\## [\1] $(date +%Y-%m-%d)/"

## [0.4.0] 2025-05-21
```
