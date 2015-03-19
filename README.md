# Go Way In Go

## Goway:

    *Martini* is a powerful package for quickly writing modular web applications/services in Golang.

    *Goway* it's an web framework,The *martini* framework code to do some optimization.


## Usage
```go
    import (
        "github.com/wackonline/goway"
    )
```     
    Within the main function is to write like this:

```go
    func main() {
        gm := goway.Bootstrap()

        gm.Get("/", func() string {
            return "hello,world"
        })
        
        gm.RunOnAddr("0.0.0.0:8080")
    }

```
## Routing
    In Goway, a route is an HTTP method paired with a URL-matching pattern. Each route can take one or more handler methods:

```go
    m.Get("/", func() {
      // show something
    })

    m.Patch("/", func() {
      // update something
    })

    m.Post("/", func() {
      // create something
    })

    m.Put("/", func() {
      // replace something
    })

    m.Delete("/", func() {
      // destroy something
    })

    m.Options("/", func() {
      // http options
    })

    // You can also create routes for static files
    pwd, _ := os.Getwd()
    gm.Static("/public", pwd)

```
## Other example
*   read test file
    [example/test.go](example/test.go)

## About Goway
    Inspired by *Express*(Nodejs) and *Sinatra*(Ruby) and *Martini*(Go) and *Symfony*(PHP).
    This framework is simple enough, and the use of modular programming, this is a way I like it very much.
    Subsequent functional may not continue like as *Martini*,*Goway* will learn other good characteristics of the web framework.


## License
    Go Way is released under the GPLV3 license:
    [License](https://github.com/wackonline/structrecord/blob/master/LICENSE)