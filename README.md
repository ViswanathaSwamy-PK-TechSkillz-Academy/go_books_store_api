# Books Store Web API and Web Application in Go lang

I am learning Web API, and Web Application in Go Language from different Video Courses, Books, and Websites

## Reference(s)

> 1. <https://go.dev/learn/>
> 1. <https://shell.cloud.google.com/?walkthrough_tutorial_url=https%3A%2F%2Fraw.githubusercontent.com%2Fgolang%2Ftour%2Fmaster%2Ftutorial%2Fweb-service-gin.md&pli=1&show=ide&environment_deployment=ide>
> 1. <https://app.pluralsight.com/library/courses/go-building-web-services-applications/table-of-contents>

## Setup

```powershell

```

## Understanding Http Requests / Responses / Methods / Status Codes

> 1. Discussion and Demo
> 1. <https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages>
> 1. <https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET>
> 1. <https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200>

## REST (Representational State Transfer)

> 1. Discussion and Demo
> 1. <https://en.wikipedia.org/wiki/REST>

## Few Options to develop Web API in go lang

> 1. Discussion and Demo

### 1. **net/http (Default Go HTTP package)**

> 1. `Documentation:` <https://pkg.go.dev/net/http>
> 1. `Overview`: The standard library's `net/http` package is a lightweight, built-in solution that provides everything needed to serve HTTP requests without external dependencies.
> 1. `When to use`: Ideal for small, simple projects or when you want full control over the HTTP server without relying on third-party libraries.

### 2. **Gorilla Mux**

> 1. `Documentation:` <https://github.com/gorilla/mux>
> 1. `Overview`: `Gorilla Mux` is a powerful URL router and dispatcher that is more flexible than `net/http`. It allows for complex routing, including variables in URLs, regular expressions, and subrouters.
> 1. `When to use`: Suitable for more complex routing needs while maintaining a balance between simplicity and flexibility. It’s a widely used and trusted library.

### 3. **Gin**

> 1. `Documentation:` <https://gin-gonic.com/docs/>
> 1. `Overview`: Gin is a lightweight and fast web framework that includes a router, middleware, and useful utilities for building APIs. It is known for its performance due to its low memory footprint.
> 1. `When to use`: Use Gin if you need high performance, want minimalistic code, and require middleware support like logging, error handling, and more. It's popular for RESTful APIs.

## What we will be using?

> 1. Discussion and Demo
> 1. [net/http (Default Go HTTP package)](https://pkg.go.dev/net/http)
