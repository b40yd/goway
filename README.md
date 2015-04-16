# Go Way In Go
[![wercker status](https://app.wercker.com/status/04753d15a770eedafaa8fcfcfd1cd8e9/s/master "wercker status")](https://app.wercker.com/project/bykey/04753d15a770eedafaa8fcfcfd1cd8e9)
### Goway:

*Martini* is a powerful package for quickly writing modular web applications/services in Golang.

*Goway* it's an web framework,The *martini* framework code to do some optimization.


### Usage
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

        gm.run()
        //gm.RunOnAddr("0.0.0.0:8080")
    }

```
### Config
In Goway, is configure web app config file.It mainly to inform how the app works,it's an JSON data struct.

```go
    {
      // App version
      "version":"0.0.1",
      // Application debugging information
      // false and true
      "debug":true,
      // Logging
      // E_ALL|E_ERROR|E_WARNING|E_STRICT|E_NOTICE
      "logger":"E_ALL",
      // App run environment
      // development|testing|product
      "env":"development",
      // Setting static directory path
      // Directory to the current app running directory to the root directory
      "staticPath": "/public",
      // HTTP Server ip address
      "httpServer":"0.0.0.0",
      // HTTP Server port
      "serverPort":"8080"
    }
```
### Routing
In Goway, a route is an HTTP method paired with a URL-matching pattern. Each route can take one or more handler methods:

```go
    gm.Get("/", func() {
      // show something
    })

    gm.Patch("/", func() {
      // update something
    })

    gm.Post("/", func() {
      // create something
    })

    gm.Put("/", func() {
      // replace something
    })

    gm.Delete("/", func() {
      // destroy something
    })

    gm.Options("/", func() {
      // http options
    })

    // You can also create routes for static files
    pwd, _ := os.Getwd()
    gm.Static("/public", pwd)

```
### Logger
In Goway, a logger is an HTTP request after call debug infomation

```go
    // Tlogs is map[int]string data struct
    var Logs = []Tlogs
    var log = Goway.Tlogs{}
    // E_ERROR | E_WARNING | E_NOTICE | E_STRICT
    log[E_ERROR] = "the is error!"
    Logs = append(Logs,log)
    gm.Logger.Use(Logs)
    //OR
    gm.Logger.Error("the is error!")
    gm.Logger.Notice("the is notice!")
    gm.Logger.Warning("the is warning!")
    gm.Logger.Strict("the is strict!")
    //use params
    sid := 110
    gm.Logger.Error("the is error! Sessionid: %d", sid)
    gm.Logger.Notice("the is notice! Sessionid: %d", sid)
    gm.Logger.Warning("the is warning! Sessionid: %d", sid)
    gm.Logger.Strict("the is strict! Sessionid: %d", sid)
```
### Other example
*   read test file
    [example/test.go](example/test.go)

### About Goway
Inspired by *Express*(Nodejs) and *Sinatra*(Ruby) and *Martini*(Go) and *Symfony*(PHP).
This framework is simple enough, and the use of modular programming, this is a way I like it very much.
Subsequent functional may not continue like as *Martini*,*Goway* will learn other good characteristics of the web framework.


### License
Go Way is released under the GPLV3 license:
    [License](LICENSE)
