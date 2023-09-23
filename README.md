# Get it done

A CLI tool that lets you save your most used CLI commands easily.

## Install
```sh
# Clone repo
git clone https://github.com/Reljod/getitdone.git && cd getitdone

# Install binary via go
go install

# Check usage
getitdone -h
```

## Usage

#### Save last command
```bash
getitdone save "<name>"
```

#### Save command
```bash
getitdone save "<name>" "node server.js"  # or
getitdone save "<name>" 'cat hello.txt | grep "world"'
```

#### List saved commands
```bash
getitdone ls

# <name> - <command>
# command1 => node server.js
# command2 => cat hello.txt | grep "world"
```

#### Use command
```bash
getitdone command1
# Runs: node server.js
```
