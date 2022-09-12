# Ex008

Learning about http and the methods.

## Code

The example code can be found in the [web.go](web.go) file.

## Run

```bash
# Run the app
$ go run .

# Post works
$ curl -X POST localhost:4000/article/create
HTTP/1.1 200 OK
Date: Sun, 21 Aug 2022 14:24:17 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8
Create a new article...

# GET, PUT, DELETE do not work
$ curl -i -X GET localhost:4000/article/create
HTTP/1.1 405 Method Not Allowed
Date: Sun, 21 Aug 2022 14:24:24 GMT
Content-Length: 18
Content-Type: text/plain; charset=utf-8

Method not allowed
```

## Learnings

- go's http package does not support routing based on methods
- http status codes can be provided with WriteHeader() and Write()
- WriteHeader() can be called only once per response
- WriteHeader() must be called before Write()
- Header must be set before WriteHeader()
- Error() is a short handle for WriteHeader(CODE) and Write(MESSAGE)
- go's http package provides constants for status codes and methods
- all constants can be found [here](https://pkg.go.dev/net/http/#pkg-constants)
- Go will automatically set three system-generated headers Date, Content-Length
  and Content-Type
