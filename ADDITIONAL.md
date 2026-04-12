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
