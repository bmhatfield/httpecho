# httpecho
A small utility for HTTP diagnostics.

## Purpose
This is a small utility for debugging HTTP interactions by printing information about the requests it receives, or by returning specific response codes depending on the request format.

## Features
- Respond to any request, with any method, and dump the full contents of the request, including headers and body, using `httputil.DumpRequest()`, and then return a 200.
- Force a response code to requests (any method) on the path `/code/{DESIRED_RESPONSE_CODE}` (eg `curl -i localhost:8080/code/401`). This is useful for testing API clients, load balancers, or proxies.

## Further documentation
- It's a short utility, so don't be afraid to peek at the code.
- The command-line supports `-help`:

```
Usage of httpecho:
  -address="": the address to bind to. Default is 0.0.0.0
  -port="8080": the local port to bind to
```
