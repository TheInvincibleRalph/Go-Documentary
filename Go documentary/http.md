In Go, an `http.Request` provides various pieces of information about the incoming HTTP request. Here's a breakdown of what you can access:

### 1. **Request Method**
   - `r.Method` (e.g., `GET`, `POST`, `PUT`, etc.)
   - This tells you the HTTP method used for the request.

### 2. **Request URL**
   - `r.URL.Path` gives the path part of the URL (e.g., `/users/profile`).
   - `r.URL.Query()` returns query parameters as a `url.Values` map (e.g., `?id=123`).

### 3. **Headers**
   - `r.Header` contains the request headers as a map, where each key corresponds to a header field.
   - Example: `r.Header.Get("Content-Type")` gets the `Content-Type` header.

### 4. **Body**
   - `r.Body` is an `io.ReadCloser`, allowing you to read the body of the request (commonly for `POST` and `PUT` requests).
   - Example: `body, err := ioutil.ReadAll(r.Body)` reads the entire body into a byte slice.

### 5. **Form Data**
   - `r.Form` and `r.PostForm` are used to retrieve form data (from URL-encoded forms).
   - Use `r.ParseForm()` to parse URL query parameters and form data.
   - `r.PostFormValue("name")` can access form values directly.

### 6. **Cookies**
   - `r.Cookies()` returns all the cookies sent by the client.
   - `r.Cookie("session_id")` retrieves a specific cookie by name.

### 7. **Remote Address**
   - `r.RemoteAddr` provides the IP address and port of the client making the request (e.g., `192.168.1.1:54321`).

### 8. **Request Protocol**
   - `r.Proto` gives the protocol used (e.g., `HTTP/1.1`, `HTTP/2.0`).

### 9. **Host**
   - `r.Host` returns the host part of the URL (e.g., `example.com`).

### 10. **TLS Information (if applicable)**
   - `r.TLS` contains information about the TLS connection (e.g., certificates), if the request was made over HTTPS.

### 11. **Request Context**
   - `r.Context()` provides access to a `context.Context`, which can be used for managing request-scoped values, deadlines, cancellations, etc.

### 12. **User Agent**
   - `r.UserAgent()` provides the user agent string (e.g., browser or client making the request).

### Example of accessing some of this data in a handler:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Method:", r.Method)               // Request method (GET, POST, etc.)
    fmt.Println("URL Path:", r.URL.Path)           // Request URL path
    fmt.Println("User-Agent:", r.UserAgent())      // User agent string
    fmt.Println("Remote Address:", r.RemoteAddr)   // Client's IP address
    fmt.Println("Query Params:", r.URL.Query())    // Query parameters
    cookie, _ := r.Cookie("session_id")            // Accessing a specific cookie
    fmt.Println("Session ID:", cookie.Value)
}
```

This allows you to get almost all relevant information about the incoming HTTP request in Go.

### Creating Go Clients amd the use of Pointers

The reason we use a pointer (`&http.Client{}`) to create an `http.Client` in Go is mainly due to how Go handles struct types and pointers, as well as the intended use of the `http.Client` for making multiple requests. Here are the key reasons:

### 1. **Shared State and Reusability**
   - The `http.Client` is designed to be **reused** across multiple requests. When you use a pointer (`&http.Client{}`), it allows you to **share the same client** instance across multiple requests, maintaining state like connection pooling and configurations such as timeouts or custom transport.
   - If you created the client as a value (i.e., `client := http.Client{}`), you'd end up creating a **new instance every time** you use it, losing the advantages of connection pooling and reuse, which are important for efficiency in network operations.

### 2. **Efficiency (Connection Pooling)**
   - The `http.Client` maintains **persistent connections** (connection pooling) to the server for performance. This helps avoid creating a new TCP connection for every request.
   - When using a pointer, you can reuse the same `http.Client` instance and benefit from the pooled connections, which reduces the overhead of opening and closing connections frequently.

### 3. **Mutation of Struct Fields**
   - If you need to modify certain fields of the `http.Client` (e.g., `Timeout`, `Transport`, or setting up a custom `RoundTripper`), using a pointer ensures that you're working with the **same instance** of the client across different parts of your code.
   - Without a pointer, changes to the struct would only affect a copy, and those changes wouldn't be reflected in subsequent uses of the client.

### 4. **Idiomatic Go Usage**
   - In Go, it's common practice to use pointers for objects that you intend to modify or that have internal state, like the `http.Client`. This is idiomatic because it allows you to modify the client and pass the same instance around without unnecessary copying.
   
### Example

```go
// Correct usage (pointer)
client := &http.Client{
    Timeout: 10 * time.Second,  // Setting custom timeout
}

// Reuse the same client for multiple requests
resp, err := client.Get("http://example.com")
```

In this example, using a pointer means the same `client` instance is reused, benefiting from any internal optimizations like connection reuse.

### Why not use a value directly?
If you were to create a value like this:

```go
client := http.Client{}
```

- Each time you pass `client` around, Go would **copy** the entire struct, potentially leading to inefficiencies (especially if fields like connection pooling need to be shared).
- Changes made to one copy of the client wouldn't reflect in other parts of the program where you expected the same `client` instance.

### Conclusion
Using a pointer to create an `http.Client` allows for **efficient reuse**, supports **connection pooling**, and ensures that any modifications to the client (like setting timeouts or custom transports) are reflected across the program. This makes it the standard and recommended approach in Go for HTTP clients.