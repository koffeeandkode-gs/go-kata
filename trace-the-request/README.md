# GO-12. Trace the Request

Build HTTP middleware that assigns a unique ID to every request so the request can be followed through the application logs.

## Requirements

1. Create an HTTP server using only the Go standard library.
2. Implement middleware that generates a non-empty, unique request ID for every incoming request.
3. Store the request ID in the request's `context.Context`.
4. Make the request ID available to downstream HTTP handlers through the context.
5. Include the request ID in log messages written by both the middleware and a handler.
6. Return the request ID to the client in an `X-Request-ID` response header.
7. Keep the context key private and strongly typed to avoid collisions.

## Suggested behavior

Add a simple endpoint such as `GET /hello`. A request to the endpoint should:

- receive a newly generated request ID;
- log the HTTP method, path, and request ID;
- let the handler retrieve and log the same ID; and
- include that ID in the response headers.

Make several requests and confirm that each request gets a different ID while every log entry for a single request uses the same ID.

## Constraints

- Use only the standard library.
- Keep request-ID generation and logging concerns in middleware rather than duplicating them in every handler.
- Do not use a package-level mutable variable without making it safe for concurrent requests.

## Run and verify

Once you have added your implementation, run it from the repository root:

```sh
go run ./trace-the-request
```

Then call the endpoint from another terminal:

```sh
curl -i http://localhost:8080/hello
```

The response should contain an `X-Request-ID` header, and the server logs should show that same value.

## Expected output

```text
HTTP/1.1 200 OK
X-Request-Id: <request-id>

Hello World
```

The server logs should use the same `<request-id>` in both the middleware and handler.
