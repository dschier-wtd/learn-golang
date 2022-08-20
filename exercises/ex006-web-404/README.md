# Ex006

Learning about http and notFound handling.

## Code

The example code can be found in the [web.go](web.go) file.

## Learnings

- go routing knows fixed paths and subtree patterns
- fixed paths are not ending with a slash
- subtree patterns are ending with a slash
- subtree patterns translate to `/**` which is a catch-all pattern
- restriction for the `/` pattern must be done in the function
- go knows a http.notFound handler
