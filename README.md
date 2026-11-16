# iscale

An early, work-in-progress image-processing build tool written in Go. The `imgresize` package does multithreaded image resizing, splitting the image into horizontal strips that run in separate goroutines. The `builder` package loads JSON job files (see `testdata/builder/jobs/`). Each job names a tool, an input folder filter, an output path, a filename pattern such as `{inputFileBase}-{#}.jpg`, and options that can be overridden per file or folder. A file cache tracks what has already been built.

There is also a stub HTTP server in `server/`, meant to become a web UI for editing jobs. `main.go` currently just initializes the builder against `testdata/builder`.
