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