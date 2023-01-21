# Go tutorials


<details>
    <summary>Table of Content</summary>

  - [Tutorials](#tutorials)
  - [Commands](#commands)
    - go mod init - [Initialize a new module](#initialize-a-new-module)
    - go mod tidy - [Add missing dependencies or remove unused ones](#add-missing-dependencies-or-remove-unused-ones)
    - go mod edit - [Edit and format go.mod file](#edit-and-format-gomod-file)
    - go mod vendor - [Copy dependencies with mod vendor](#copy-dependencies-with-mod-vendor)
    - go run - [Run the app](#run-the-app)
    - go test - [Run tests](#run-tests)
    - go build - [Build](#build)
    - go install - [Install](#install)
    - go fmt - [Format files](#format-files)
    - go fix - [Update packages to use new APIs](#update-packages-to-use-new-apis)
    - go vet - [Get mistakes with go vet](#get-mistakes-with-go-vet)
    - go clean - [Remove files and clean cache](#remove-files-and-clean-cache)
    - go list - [List packages and modules](#list-packages-and-modules)
    - go doc - [Generate documentation](#generate-documentation)
    - [Debugging](#debugging)
    - [Linting & fixing common problems](#linting--fixing-common-problems)
        - [golint (deprecated)](#golint-deprecated)
        - [revive](#revive)
        - [go vet](#go-vet)
  - [The language](#the-language)
    - [Exported name](#exported-name)
    - [:= operator](#-operator)
    - [Multiple return values](#multiple-return-values)
    - [Named results](#named-results)
    - [Arrays and slices](#arrays-and-slices)
      - [Arrays](#arrays)
      - [Slices](#slices)
    - [The init function](#the-init-function)
  - [Random](#random)
    - [go:embed](#goembed)
        - [Simple use](#simple-use)
        - [Use conditionally](#use-conditionally)
        - [Gotchas](#gotchas)
    - [go:generate](#gogenerate)
        - [Alternative - using -ldflags](#alternative---using--ldflags)
        - [go:generate use](#gogenerate-use)
        - [Links](#links)
    - [Build Constraints](#build-constraints)
    - [flag.Func](#flagfunc)
  - [Testing and linting in pipeline](#testing-and-linting-in-pipeline)
  - [Other resources to check](#other-resources-to-check)

</details>

## Tutorials

Follow along for go.dev [tutorials](https://go.dev/doc/tutorial/):
- [Getting started with go](https://go.dev/doc/tutorial/getting-started)
- [Create Go modules series](https://go.dev/doc/tutorial/create-module)
- [RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin). Tutorial's README [here](./3-restful-api).

Follow along for [freeCodeCamp tutorials](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org). 
Tutorial's README [here](./4-freeCodeCamp-Learn-Go-Programming).

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

### Copy dependencies with mod vendor

[here](https://pkg.go.dev/cmd/go#hdr-Make_vendored_copy_of_dependencies) and [here](https://go.dev/ref/mod#go-mod-vendor)

> go mod vendor [-e] [-v] [-o outdir]

- `[-v]`
  - print the names of vendored modules and packages to standard error
- `[-e]`
  - causes vendor to attempt to proceed despite errors encountered while loading packages
- `[-o]`
  - causes vendor to create the vendor directory at the given path instead of "vendor"

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

### Format files

[here](https://pkg.go.dev/cmd/go#hdr-Gofmt__reformat__package_sources)

> go fmt [-n] [-x] [packages]

- `[-n]`
  - prints commands that would be executed
- `[-x]`
  - prints commands as they are executed
- `[-mod]`
  - sets which module download mode to use: `readonly` or `vendor`

### Update packages to use new APIs

> go fix [-fix list] [packages]

- `[-fix]`
  - sets a comma-separated list of fixes to run. The default is all known fixes.

### Get mistakes with go vet

[here](https://pkg.go.dev/cmd/go#hdr-Report_likely_mistakes_in_packages)

> go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]

- `[-n]`
  - prints commands that would be executed
- `[-x]`
  - prints commands as they are executed
- `[-vettool=prog]`
  - selects a different analysis tool with alternative or additional checks

Example:
- the 'shadow' analyzer can be built and run:
> go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
> go vet -vettool=$(which shadow)

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

### Generate documentation

[here](https://pkg.go.dev/cmd/go#hdr-Show_documentation_for_package_or_symbol)

> go doc [doc flags] [package|[package.]symbol[.methodOrField]]

Example:
- documentation for formatting files
> go doc cmd/gofmt
- show all documentation for the package
> go doc -all

### Debugging

The standard debugger for Go applications is [Delve](https://github.com/go-delve/delve) - [commands](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md)

> go install github.com/go-delve/delve/cmd/dlv@latest

Check it was correctly installed with `dlv version` or with a path `~/go/bin/dlv`.

Using Delve:
> dlv debug main.go
> break bp1 main.main:3
> condition bp1 i == 2
> continue
> next
> step
> stepout
> restart
> exit

> print <expr>
> set <variable> = <value>
> locals
> whatis <expr>

### Linting & fixing common problems

#### golint (deprecated)

- applies the rules based on [Effective Go](https://go.dev/doc/effective_go) and a collection of comments from [code reviews](https://github.com/golang/go/wiki/CodeReviewComments).
- can't be configured

#### revive

[here](https://github.com/mgechev/revive)

- provides support for controlling which rules are applied

> go install github.com/mgechev/revive@latest
> revive

- you can disable or enable a rule in code:
```
// revive:disable:exported
// revive:enable:exported
``` 

- the configuration file for this linter is `revive.toml` 
    - options [here](https://github.com/mgechev/revive#configuration)
    - recoomented configuration [here](https://github.com/mgechev/revive#recommended-configuration)
- run the linter with configuration file
> revive -config revive.toml

#### go vet

> go vet
> go vet main.go
> go vet -json main.go

> go vet -assign=false
> go vet -assign

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

## Random

### go:embed

- [documentation](https://pkg.go.dev/embed) and [article link](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)
- include the contents of arbitrary files and directories with `//go:embed FILENAME(S)` followed by 
a `string` or `[]byte` variable for one file or `embed.FS` for a group of files
- patterns like `files/*.html` will also work (but not `**/*.html` recursive globbing)

#### Simple use

```go
package main

import (
    _ "embed"
    "fmt"
    "strings"
)

var (
    Version string = strings.TrimSpace(version)
    //go:embed version.txt
    version string
)

func main() {
    fmt.Printf("Version %q\n", Version)
}
```

#### Use conditionally

Include version information conditionally based on whether a build tag is passed to the go tools:
```go
// version_dev.go file:

//go:build !prod
// +build !prod

package main

var version string = "dev"
```

Run:
```bash
$ go run .
Version "dev"
```

```go
// version_prod.go file:

//go:build prod
// +build prod

package main

import (
    _ "embed"
)

//go:embed version.txt
var version string
```

Run:
```bash
$ go run -tags prod .
Version "0.0.1"
```

#### Gotchas

- need to import the `embed` package in any file that uses an embed directive
- can only use `//go:embed` for variables at the package level, not within functions or methods
- when you include a directory, it won’t include files that start with `.` or `_`, but if you use a wildcard, 
like `dir/*`, it will include all files that match, even if they start with `.` or `_`
    - use `//go:embed dir` or `//go:embed dir/*.ext` for security reasons (otherwise you will accidentally include Mac OS's `.DS_Store`)
- for security reasons, Go also won’t follow symbolic links or go up a directory when embedding

### go:generate

- [article link](https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/)
- automate the running of tools to generate source code before compilation

#### Alternative - using -ldflags

Can use `-ldflags` to override string values:
```go
package main

import (
	"fmt"
)

var VersionString = "unset"

func main() {
	fmt.Println("Version:", VersionString)
}
```

Run:
```bash
$ go run main.go
Version: unset

$ go run -ldflags '-X main.VersionString=1.0' main.go
Version: 1.0

# Can use git commit hash:
$ go run -ldflags "-X main.VersionString=`git rev-parse HEAD`" main.go
Version: db5c7db9fe3b632407e9f0da5f7982180c438929

# Need single quote if there is a space
$ go run -ldflags "-X 'main.VersionString=1.0 (beta)'" main.go
Version: 1.0 (beta)

# Ca use the standard Unix date command
$ go run -ldflags "-X 'main.VersionString=`date`'" main.go
Version: Sun Nov 27 16:42:10 EST 2016
```

#### go:generate use

```go
package project

//go:generate echo Hello, Go Generate!

func Add(x, y int) int {
	return x + y
}
```

Run:
```bash
$ go generate
Hello, Go Generate!
```

Can be used with tools like [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer), 
[jsonenums](https://github.com/campoy/jsonenums) and [schematyper](https://github.com/idubinskiy/schematyper).

#### Links

- [Eli Bendersky's website](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate/)
- [Reducing boilerplate with go generate](https://blog.gopheracademy.com/advent-2015/reducing-boilerplate-with-go-generate/)

### Build Constraints

- [documentation](https://pkg.go.dev/go/build#hdr-Build_Constraints), [go help buildconstraint](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- condition under which a file should be included in the package

### flag.Func

- [article link](https://blog.carlmjohnson.net/post/2020/add-func/)

## Testing and linting in pipeline

Project [here](https://github.com/mariamihai/test-go-github-actions) for GitHub Actions use.

## Other resources to check
- [go.dev - Effective Go](https://go.dev/doc/effective_go)
- [Go by example](https://gobyexample.com/)
- [freeCodeCamp.org](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org)
- [Udemy - Building web applications with go - intermediate level](https://www.udemy.com/course/building-web-applications-with-go-intermediate-level)
- [Go lang by example](https://golangbyexample.com/golang-comprehensive-tutorial/)
- [Go go-to guide](https://yourbasic.org/golang/)
- [GitHub - Go Courses](https://github.com/golang/go/wiki/Courses)
- [Golang Programs](https://www.golangprograms.com/go-language.html)
- [Programiz - Getting started](https://www.programiz.com/golang/getting-started)
- [go.dev - Modules blog series](https://go.dev/blog/using-go-modules)
- [go.dev - Maps](https://go.dev/blog/maps)
- [go.dev - The blank identifier](https://go.dev/doc/effective_go#blank)
- [go.dev - Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [go.dev - Developing and publishing modules](https://go.dev/doc/modules/developing)
- [go.dev - Getting started with multi-module workspaces](https://go.dev/doc/tutorial/workspaces)
  - needs go 1.18 or later
- [go.dev - Getting started with generics](https://go.dev/doc/tutorial/generics)
  - needs go 1.18 or later
- [go.dev - Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)
  - needs go 1.18 or later
- [Defer, panic and recover](https://go.dev/blog/defer-panic-and-recover)
- [Defer](https://go.dev/doc/effective_go#defer)
- [Slice](https://zetcode.com/golang/slice/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/index.html)
- [Lets talk about logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
- [testing.T type](https://pkg.go.dev/testing#T)
- [Test_packages)](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [Sorting a slice of a given struct by multiple fields](https://www.linkedin.com/pulse/golang-sorting-slice-given-struct-multiple-fields-tiago-melo/)
- [Skip tests in go](https://blog.dharnitski.com/2019/04/29/skip-tests-in-go/)
- [White box testing](https://devmethodologies.blogspot.com/2013/11/unit-tests-white-box-testing.html)
- [Switches in Go](https://gobyexample.com/switch)
- [Map string interface](https://bitfieldconsulting.com/golang/map-string-interface)
- [Methods in Go](https://go101.org/article/method.html)
- [Method Sets - pointer vs value receiver](https://stackoverflow.com/questions/33587227/method-sets-pointer-vs-value-receiver)
- [Difference between T and T* method sets in Go](https://gronskiy.com/posts/2020-04-golang-pointer-vs-value-methods/)
- [go test command](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [go build command](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
- [go install command](https://go.dev/ref/mod#go-install)
- [go clean command](https://pkg.go.dev/cmd/go#hdr-Remove_object_files_and_cached_files)
- [go list command](https://pkg.go.dev/cmd/go#hdr-List_packages_or_modules)
- [go env command](https://pkg.go.dev/cmd/go#hdr-Print_Go_environment_information)
- [go get command](https://go.dev/ref/mod#go-get)
- [How to write go code](https://go.dev/doc/code)
- [Go GUI projects](https://github.com/go-graphics/go-gui-projects)
- [Dave Cheney blog](https://dave.cheney.net/)
- [Reflect pkg](https://pkg.go.dev/reflect)
- [Reflection in Golang](https://www.geeksforgeeks.org/reflection-in-golang/)
- [Reflection in Go](https://golangbot.com/reflection/)
- [Reflection - golangprograms](https://www.golangprograms.com/reflection-in-golang.html)


- [Go and Proxy Servers](https://eli.thegreenplace.net/2022/go-and-proxy-servers-part-1-http-proxies/)
- [Serving static files and web apps](https://eli.thegreenplace.net/2022/serving-static-files-and-web-apps-in-go/)


Exercises and challenges:
- [Codewars](https://www.codewars.com/)
- [Kattis](https://www.kattis.com/)
- [Gophercises](https://gophercises.com/)
- [Golangr](https://golangr.com/exercises/)
- [TutorialEdge](https://tutorialedge.net/challenges/go/)
- [GitHub - Plutov](https://github.com/plutov/practice-go)
- [GitHub - inancgumus](https://github.com/inancgumus/learngo)
- [GitHub - RajaSrinivasan](https://github.com/RajaSrinivasan/assignments)
- [exercism](https://exercism.org/)
- [HackerRank](https://www.hackerrank.com/)
- [Go Bootcamp](https://www.golangbootcamp.com/book/exercises)
