# Ex005

Learning about http and routing.

## Code

The example code can be found in the [web.go](web.go) file.

## Learnings

- one can do routing with http.HandleFunc
- http.HandleFunc is a function that takes a path and a function as arguments
- the function is called when the path is requested
- the function is passed the request and response objects
- the function can then do whatever it wants with the request and response
  objects
- longer paths are handled first
- request url paths are automatically sanitized (multiple slashes, dots, etc.)
