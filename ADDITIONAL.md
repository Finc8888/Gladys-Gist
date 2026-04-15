# Additional information
- In Go's servemux, longer URL patterns always take precendece over shorter one
- http.Handle() and http.HandleFunc() allow you register routes without declaring a servemux like this
```bash
http.handleFunc("/", home)
err := http.ListenAndServe(":4000", nil)
```
But behind the scenes, these functions register their routes with DefaultServeMux.
But DefaultServeMux is a global variable, any package can access it and register a route - including any third-party packages. It could be to expose a malicious handler to the web. So use own locally-scoped servemux instead
- It's only possible to call w.WriteHandler() once per response. If you dont't call w.WriteHandler() explicitly, then the first call to w.Write() will automatically send a 200 OK.
- The pattern of passing http.ResponseWriter to other functions is super-common in Go. In practice, it's quite rare to use the w.Write() and w.WriteHeader() methods directly.
- Go will attempt to set the correct Content-Type by content sniffing the response body with the http.DetectContentType() function. If the function can't guess the content type, Go will fall back to setting the header `Content-type: application/octet-stream` But this fuction can't distinguish JSON from plain text. So, by default, JSON responses will be sent with a `Content-Type: text/plain; charset=utf-8` header. To prevent this it needs to set the correct header manually like so:
```go
w.Header().Set("Content-Type", "application/json")
````
- When you're using the Set(), Add(), Del(), Get() and Values() methods on the header map, the header name will always be canonicalization using the `textproto.CanonicalMimeHeaderKey()` function. If you need to avoid this canonicalization behavior you can exit the underlying header map directly:
```go
w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
```
If a HTTP/2 connection is being used, Go will always automatically convert the header names and values to lowercase for you as per the HTTP/2 secifications.

- The Del() method doesn't remove system-generated headers. To suppress these, you need to access the underlying header map directly and set the value to nil.
```go
w.Header()["Date"] = nil
```
- There's no single right way to structure web application in Go. But there is [popular](https://github.com/thockin/go-build-template) - and [tried-and-tested](https://peter.bourgon.org/go-best-practices-2016/) approach.

- The `cmd` directory contains the application-specific code for the executable applications in the project. For now we'll have just one executable application - the web application - which lives under the `cmd/web` dirrectory.

- The `internal` directory contains the ancillary non-application-specific code used in the project. We use it to hold potentially reusable code like validation helplers and the SQL db models for the project. Any packages under `internal` cannot be imported by code outside of our project. This is useful because it prevents other codebases from importing and relying on the packages in our `internal` directory.

- The `ui` directory contains the user-interface assets used by the web application. Specifically, tthe `ui/html` dirrectory contains static files(like CSS and images).
