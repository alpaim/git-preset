# Git Preset

A simple CLI tool to switch between different **local** Git user configurations.

## Overview

Git Preset allows you to define and quickly switch between multiple Git user identities (name and email combinations) with a single command.

## Installation

### Build from Source

```bash
# Clone the repository
git clone https://github.com/alpaim/git-preset.git
cd git-preset

# Build the binary
go build

# Move to a directory in your PATH
mv git-preset /usr/local/bin/
```

## Configuration

### Config File Location

The configuration file should be stored at:

```
~/.config/git-preset/config.yaml
```

Create this directory if it doesn't exist:

```bash
mkdir -p ~/.config/git-preset
```

### Sample Config File

Create a `config.yaml` file with your presets:

```yaml
presets:
  work:
    name: "Jane Doe"
    email: "jane.doe@company.com"
  personal:
    name: "Jane Doe"
    email: "jane.personal@example.com"
  opensource:
    name: "Jane Doe"
    email: "jane.github@example.com"
```

## Usage

Once installed and configured, switch between Git identities using:

```bash
git-preset work
```

This will set your **local** Git user name and email to the values defined in the "work" preset.

If you specify a preset that doesn't exist, the tool will show available presets:

```
Preset 'unknown' not found. Available presets: [work personal opensource]
```

## Development

### Requirements

- Go 1.16 or later
- `gopkg.in/yaml.v3` package

### Get Dependencies

```bash
go get gopkg.in/yaml.v3
```

### Build for Development

```bash
go build -o git-preset
```