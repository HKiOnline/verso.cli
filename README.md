<img src="docs/attachments/verso_logo_banner.svg" />

# verso.cli

Simple [CHANGELOG](https://keepachangelog.com) [semver](https://semver.org/) extractor CLI. If you'd like to use verso as a Go library module, have a look at the [hkionline/verso](https://github.com/hkionline/verso) -project.

## Installation

At the moment you need to have Go lang installed. Please make sure you have your GOPATH environmental variable set and exported as well. If you have these requirements fullfilled you can install Verso CLI with the following command:

```bash
go install github.com/hkionline/verso.cli/cmd/verso@latest
```

After Installation you should now have the verso command available from your favorite terminal emulator.

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

1.3.0
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
