<img src="docs/attachments/verso_logo_banner.svg" />

# verso

Simple [CHANGELOG]("https://keepachangelog.com") [semver]("https://semver.org/") extractor CLI and libray.

## Usage

### Usage as CLI

When using verso as CLI, it expects to find a CHANGELOG-file from the working directory you are in.

Get the latest version in the CHANGELOG.
```bash
verso latest
```

Get list of versions found from the CHANGELOG.
```bash
verso list
```

### Usage as library package

```go
import "github.com/hkionline/verso/pkg/verso"

changelog, err := verso.Parse("path/to/CHANGELOG.md")

```
