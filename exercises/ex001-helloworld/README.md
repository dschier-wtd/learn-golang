# Ex001

Hello World!

## Code

Create a file with below content:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")
}
```

## Build & Run

Initialize a new project:

```bash
$ go mod init exercise
go: creating new go.mod: module exercise
go: to add module requirements and sums:
        go mod tidy
```

Run without building:

```bash
$ go run hello.go
Hello, World!
```

Build and run:

```bash
$ go build hello.go
$ ./hello
Hello, World!
```

## Learnings

This section summarizes my learnings.

- basic go structure
- build and run code
- the fmt package contains functions for formatting
- the fmt package contains functions for printing
