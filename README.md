# coco-filer

coco-filer is a Go library for easy management of file operations such as saving, moving, and handling metadata. It is especially useful for web applications or backend services that deal with file uploads and file management.

## Features
- Encapsulation of file operations (e.g., saving, moving)
- Management of file metadata (name, path, size, modification time, file mode, MIME type)
- Configurable default directory for file operations

## Installation

Add the module to your project:

```
go get github.com/a-digi/coco-filer
```

## Usage

```go
package main

import (
    "github.com/a-digi/coco-filer/filer"
)

func main() {
    fm := filer.NewFileManager("") // Uses the default directory "data/uploads"
    // Perform file operations with fm ...
}
```

## Structure
- `filer/file_manager.go`: Main file for file operations and metadata
- `filer/move.go`: Functions for moving files
- `filer/multipart.go`: Support for multipart uploads

## License

See [LICENSE](LICENSE) for more information.
