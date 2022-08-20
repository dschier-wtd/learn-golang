# Ex007

Learning about http and the default multiplexer.

## Code

The example code can be found in the [web.go](web.go) file.

## Learnings

- go http has a default mux (DefaultServeMux.)
- the default mux is used when no other mux is provided
- it should be avoided in production
- it is stored in the net/http global variable
