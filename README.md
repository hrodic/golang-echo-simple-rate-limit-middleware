# golang-echo-simple-rate-limit-middleware
Rate limit middleware for golang echo framework

## Requirements

* go 1.12
* github.com/labstack/echo v3.3
* https://godoc.org/golang.org/x/time/rate

## Usage

As any other middleware:

```
e := echo.New()
e.Use(RateLimitWithConfig(RateLimitConfig{
  Limit: 2,
  Burst: 2,
}))
```

run your server and try out the middleware easily with curl

```
while true; do curl https://localhost:8080 --insecure; done
```

responses

```
...
...
...
HTTP/2 200 
access-control-allow-origin: 
content-type: application/json; charset=UTF-8
vary: Accept-Encoding
vary: Origin
x-content-type-options: nosniff
x-frame-options: SAMEORIGIN
x-request-id: eVdqddyz8XHFJ2UgDjlOb7QbfY23Hfyi
x-xss-protection: 1; mode=block
content-length: 32
date: Sat, 17 Aug 2019 13:59:49 GMT

{"application":"","version":""}
HTTP/2 429 
content-type: application/json; charset=UTF-8
content-length: 32
date: Sat, 17 Aug 2019 13:59:49 GMT

{"message":"Too Many Requests"}
```
