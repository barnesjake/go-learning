### run app
`go run .`

### packages
https://pkg.go.dev/

### adding new packages
`go mod tidy`

### Exported names
a function whose name starts with a capital letter can be called by a functioun not in the same package.

### functions
![img.png](images/functions_img.png)`func

### := operator
`:=` operator is a shortcut for declaring and initializing a variable. i.e.
```go
    message := "123"
```

instead of 
```go
    var message string
    message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

### printers
```
If the name starts with Print, it writes to standard output
If the name starts with Fprint, it writes to an io.Writer (possibly to a file, thus the 'f')
If the name starts with Sprint, it writes to a string and returns that string
If the name ends with f, it is a formatted print, that is, it gets a format argument like "%s %d", and formats the output based on that.
If the name ends with ln like Println, it prints a newline after writing
Otherwise it simply prints its arguments using their default formats.
```

### accessing modules
`$ go mod edit -replace example.com/greetings=../greetings`


### testing
run tests:
`go test`

run tests with verbose output
`go test -v`

### build 
`$ go build`
then run `./hello`

### 


### 


### 


### 


