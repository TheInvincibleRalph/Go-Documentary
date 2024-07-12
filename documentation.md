In Go, there are several ways to create an instance of a struct. 

```go
type Comment struct {
	Text string `form:"text" json:"text"`
}
```

Here are a few common methods to create an instance of the `Comment` struct:

1. **Using the `new` function:**

```go
cmt := new(Comment)
```
This allocates memory for a `Comment` struct and returns a pointer to it.

2. **Using a composite literal:**

```go
cmt := &Comment{}
```
This creates a `Comment` struct and returns a pointer to it.

3. **Direct initialization (if you don't need a pointer):**

```go
cmt := Comment{}
```
This creates a `Comment` struct value (not a pointer).

4. **Direct initialization with field values:**

```go
cmt := &Comment{
    Text: "Example text",
}
```
This creates a `Comment` struct with the `Text` field set to "Example text" and returns a pointer to it.

5. **Using the `var` keyword:**

```go
var cmt Comment
```
This declares a variable `cmt` of type `Comment` and initializes it with zero values.

6. **Using the `var` keyword with a pointer:**

```go
var cmt *Comment
cmt = &Comment{}
```
This declares a variable `cmt` of type `*Comment` and then assigns it to a new `Comment` instance.

Each method has its own use case depending on whether you need a pointer to the struct or just the struct value.