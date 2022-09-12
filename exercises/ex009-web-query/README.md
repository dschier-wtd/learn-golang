# Ex009

Learning about http and query strings.

## Code

The example code can be found in the [web.go](web.go) file.

## Run

Now you can pass ids and change the output.

```bash
$ curl localhost:4000/article/view?id=1
Display a specific snippet with ID 1...

$ curl localhost:4000/article/view?id=2
Display a specific snippet with ID 2...

$ curl localhost:4000/article/view?id=123
Display a specific snippet with ID 123...

$ curl localhost:4000/article/view?id=foo
404 page not found

$ curl localhost:4000/article/view?id=0
404 page not found
```

## Learnings

- use the proper mux for initialization x.X
- strconv can be used to convert strings, but the syntax is somewhat ...
- that I am not sure about "id, err :="
- you can use fmt to format strings properly and use variables
- using interfaces feels weird, too
- https://www.alexedwards.net/blog/interfaces-explained
-
