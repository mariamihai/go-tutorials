# Go tutorials

Follow along for go.dev [tutorials](https://go.dev/doc/tutorial/).

## Commands

### Initialize a new module

[here](https://go.dev/ref/mod#go-mod-init)

> go mod init [module-path]
> go mod init <prefix>/<descriptive-text>

- creates a new `go.mod` file in current directory
- the file must not exist already

### Add missing dependencies or remove unused ones

[here](https://go.dev/ref/mod#go-mod-tidy)

> go mod tidy [-e] [-v] [-go=version] [-compat=version]

- [-e] - continue with errors
- [-v] - info for removed modules
- [-go] - sets version, enable or disable module graph pruning and module loading based on version
  (changes from go 1.17)
- [-compat] - set the version to be checked for compatibility

### Edit and format go.mod file

[here](https://go.dev/ref/mod#go-mod-edit)

> go mod edit [editing flags] [-fmt|-print|-json] [go.mod]

Example:
> go mod edit -replace random-prefix/greetings=../greeting

- [editing flags]
  - -module, -go=version flag sets the expected Go language version, -require=path@version / -droprequire=path, 
  -exclude=path@version / -dropexclude=path@version, -replace=old[@v]=new[@v], -dropreplace=old[@v], -retract=version / -dropretract=version
  - can have multiple editing flags
- [-fmt]
  - format
  - implicit as part of other `go mod edit` commands
- [-print]
  - print as text instead of writing to disk
- [-json]
  - returns a json instead of writing to disk
  
### Run the app

[here](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)

> go run [build flags] [-exec xprog] package [arguments...]

- compile and run the `main` go package
- single package: `go run .`
- with path: `go run my/cmd`

## The language

### Exported name

- a function named starting with a capital letter can be called by a function not in the same package
- if calling a function from a different package with a lower letter you get an error: _cannot refer to unexported name xxxxx_

### := operator

- used for declaring and initializing a variable

```go
message := fmt.Sprintf("Hi, %v. Welcome!", name)

// Same as
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

### Multiple return values

[here](https://go.dev/doc/effective_go#multiple-returns)

```
// Definition
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}

// Calling the function
message, err := greetings.Hello("")
```

### Named results

[here](https://go.dev/doc/effective_go#named-results)

- named results are initialized

```
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

## Other resources to check
- [go.dev - Effective Go](https://go.dev/doc/effective_go)
- [Go lang by example](https://golangbyexample.com/golang-comprehensive-tutorial/)
- [Go by example](https://gobyexample.com/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/index.html)
- [GitHub - Go Courses](https://github.com/golang/go/wiki/Courses)
- [Lets talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
- [Defer, panic and recover](https://go.dev/blog/defer-panic-and-recover)
- [Defer](https://go.dev/doc/effective_go#defer)
