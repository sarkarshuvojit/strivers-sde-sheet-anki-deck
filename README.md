# striver sde sheet as ANKI Decks

## Overview

## Project Structure

```
.
├── assets                 # Storage for JSON snapshots
├── cmd
│   ├── create-deck        # Transform snapshots into tagged Anki decks
│   └── download-snapshot  # Fetch new DSA study plan 
├── pkg
│   ├── types              # Type definitions
│   └── utils              # Utility functions
├── go.mod
├── go.sum
└── Makefile
```

## Prerequisites

- Go 1.20+
- Make (optional, for build scripts)

## Installation

```bash
git clone https://github.com/yourusername/dsa-study-planner.git
cd dsa-study-planner
go mod download
```

## Usage

### Download Snapshot

```bash
go run cmd/download-snapshot/main.go \
    --path ./path-to-output.json
```

### Create CSV files from snapshot that can be imported on Anki

```bash
go run cmd/create-deck/main.go \
    --path ./assets/snapshot-zero.json \
    --output-dir-path outputfolder \
    --create-dir # if you want to auto-create outputfolder
```



## Usage

```
Usage of /tmp/go-build3581604133/b001/exe/main:
  -create-dir
        Automatically create the output directory if it does not exist
  -output-dir-path string
        Destination directory path for output files (defaults to a timestamped folder) (default "assets/decks-2024-12-13_20-29-02")
  -path string
        Path to the source snapshot JSON file to be processed (default "./assets/snapshot-zero.json")
```
