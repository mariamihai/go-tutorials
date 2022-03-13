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

- `[-e]` - continue with errors
- `[-v]` - info for removed modules
- `[-go]` - sets version, enable or disable module graph pruning and module loading based on version
  (changes from go 1.17)
- `[-compat]` - set the version to be checked for compatibility

### Edit and format go.mod file

[here](https://go.dev/ref/mod#go-mod-edit)

> go mod edit [editing flags] [-fmt|-print|-json] [go.mod]

Example:
> go mod edit -replace random-prefix/greetings=../greeting

- `[editing flags]`
  - `-module`, `-go=version`, `-require=path@version` / `-droprequire=path`, `-exclude=path@version` / `-dropexclude=path@version`, 
  `-replace=old[@v]=new[@v]`, `-dropreplace=old[@v]`, `-retract=version` / `-dropretract=version`
  - can have multiple editing flags
- `[-fmt]`
  - format
  - implicit as part of other `go mod edit` commands
- `[-print]`
  - print as text instead of writing to disk
- `[-json]`
  - returns a json instead of writing to disk
  
### Run the app

[here](https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program)

> go run [build flags] [-exec xprog] package [arguments...]

- compile and run the `main` go package
- single package: `go run .`
- with path: `go run my/cmd`

### Run tests

[here](https://pkg.go.dev/cmd/go#hdr-Test_packages)

> go test [build/test flags] [packages] [build/test flags & test binary flags]

- run tests for files matching the filename pattern `*_test.go` and functions named `TestXxx` (after `Test` there needs 
to be an upper case - exported name)
- these files can contain tests, benchmark or example functions

More info:
> go help testfunc

### Build

[here](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)

> go build [-o output] [build flags] [packages]

- compiles the packages, along with their dependencies, doesn't install the results
- ignores files that end in `_test.go`

### Install

[here](https://go.dev/ref/mod#go-install)

> go install [build flags] [packages]

- compiles and installs the packages
- [compile and install](https://go.dev/doc/tutorial/compile-install)

### Remove files and clean cache

- [here](https://pkg.go.dev/cmd/go#hdr-Remove_object_files_and_cached_files)

> go clean [clean flags] [build flags] [packages]

Example:
> go clean -i <package-name>
 
- remove the corresponding installed archive or binary added by `go install`

### List packages and modules

[here](https://pkg.go.dev/cmd/go#hdr-List_packages_or_modules)

> go list [-f format] [-json] [-m] [list flags] [build flags] [packages]

Example:
- the installation path after compiling `caller` into an executable
> go list -f '{{.Target}}'

- an array will all available `.go` files under current path
> go list -f '{{.GoFiles}}'

- print the package data in JSON format
> go list -json

## The language

### Exported name

- a function named starting with a capital letter can be called by a function not in the same package
- if calling a function from a different package with a lower letter you get an error: _cannot refer to unexported name xxxxx_

### := operator

- used for declaring and initializing a variable

```
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

### Arrays and slices

[here](https://go.dev/blog/slices-intro)

#### Arrays

- the array's size is fixed
- the length is part of its type (`[4]int` is distinct and incompatible with `[3]int`)
- can be indexed - `s[n]` will return the nth element, starting from 0

```
var a [4]int
a[0] = 1

i := a[0] // i = 1
```

- arrays don't need to be initialized explicitly
- the arrays are values, the variable is not a pointer to the first element but represents the entire array
- assigning an array value makes a copy of the content

```
b := [2]string{"Penn", "Teller"}

// or
b := [...]string{"Penn", "Teller"}
```

#### Slices

```
letters := []string{"a", "b", "c"}

// or
var s[]byte
// Using func make([]T, len, cap) []T
s = make([]byte, 5, 5) // s = []byte{0, 0, 0, 0, 0}

// or
s = make([]byte, 5)
```

- can form a slice from "slicing" an existing slice or array

```
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}

// b[1:4] == []byte{'o', 'l', 'a'}

// b[:2] = []byte{'g', 'o'}
// b[2:] = []byte{'l', 'a', 'n', 'g'}
// b[:] = b

// Create slíce from array
golang := [6]byte{'g', 'o', 'l', 'a', 'n', 'g'}
s := golang[:]
```

- `b[1:4]` shares the same storage as `b`
- doesn't copy the slice's data, but it creates a new slice value that points to the original array

```
d := [byte]{'r', 'o', 'a', 'd'}
e := d[2:]

e[1] = 'm'

// d = [byte]{'r', 'o', 'a', 'm'}
// e = [byte]{'a', 'm'}
```

- getting the length (the number of elements referred to by the slice) and the capacity (the number of elements in the
  underlying array) with `len(s)` and `cap(s)`

- can't grow a slice beyond its capacity and can't index outside the bounds - these cause runtime panic
- slices can't be re-sliced below 0

- copy data from a source slice to a destination, returning the number of elements copied
> func copy (dest, src []T) int

- append the elements `x` to the end of the slice `s`; it will take care of increasing the size of the slice if needed
> func copy (s []T, x ...T) []T

Example:
```
a := make([]int, 1) // a = []int{0}

a = appent(a, 1, 2, 3) // a = []int{0, 1, 2, 3}
```

- append one slice to another with `...` to expand the second slice
```
a := []int{1, 2, 3}
b := []int{4, 5, 6}

a = append(a, b...) // a = append(a, b[0], b[1], b[2])
```

### The init function

- the `init` functions are automatically executed at program setup
- used for initializations that can't be done by declarations, verifications of program state
- first the imported packages are initialized, then the global variables are initialized, then the `init` function is called
- each file can have multiple init `functions`

## Other resources to check
- [go.dev - Effective Go](https://go.dev/doc/effective_go)
- [Go lang by example](https://golangbyexample.com/golang-comprehensive-tutorial/)
- [Go by example](https://gobyexample.com/)
- [Go go-to guide](https://yourbasic.org/golang/)
- [GitHub - Go Courses](https://github.com/golang/go/wiki/Courses)
- [Programiz - Getting started](https://www.programiz.com/golang/getting-started)
- [go.dev - Modules blog series](https://go.dev/blog/using-go-modules)
- [go.dev - Maps](https://go.dev/blog/maps)
- [go.dev - The blank identifier](https://go.dev/doc/effective_go#blank)
- [go.dev - Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [go.dev - Developing and publishing modules](https://go.dev/doc/modules/developing)
- [Defer, panic and recover](https://go.dev/blog/defer-panic-and-recover)
- [Defer](https://go.dev/doc/effective_go#defer)
- [Slice](https://zetcode.com/golang/slice/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/index.html)
- [Lets talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
- [testing.T type](https://pkg.go.dev/testing#T)
- [Test_packages)](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [Skip tests in go](https://blog.dharnitski.com/2019/04/29/skip-tests-in-go/)
- [Switches in Go](https://gobyexample.com/switch)
- [go test command](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [go build command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [go install command](https://go.dev/ref/mod#go-install)
- [go clean command](https://pkg.go.dev/cmd/go#hdr-Remove_object_files_and_cached_files)
- [go list command](https://pkg.go.dev/cmd/go#hdr-List_packages_or_modules)
- [go env command](https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information)
- [go get command](https://go.dev/ref/mod#go-get)
