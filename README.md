## Build and Installation

There are two ways to build and use the project, depending on whether you need a local binary or want to install it into your `$PATH`.

---

### Option 1. Local build (using `go build`)

Suitable for development, experiments, and code exploration.

```bash
go build -o bin/slider ./cmd/slider
```

The binary will be created in the bin/ directory (which is listed in .gitignore).

If needed, you can manually copy it to a directory that is included in your $PATH:
```bash
cp bin/slider ~/bin/
```

### Option 2. Install using go install

Suitable for installing the utility system-wide.

```bash
go install ./cmd/slider
```

The binary will be installed into $GOBIN or, if it is not set, into $GOPATH/bin.
Make sure this directory is included in your $PATH.