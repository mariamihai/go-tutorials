# freeCodeCamp.org - Learn Go Programming

Follow along for Michael Van Sickle's [tutorial](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org).

An overview of all the tutorials can be found [here](../../..).

<details>
    <summary>Table of Content</summary>

  - [Introduction](#introduction)
  - [Setting Up a Development Environment](#setting-up-a-development-environment)
    - [Setting up in Ubuntu](#setting-up-in-ubuntu)
    - [Important directories in own workspace](#important-directories-in-own-workspace)
  - [Variables](#variables)
  - [Primitives](#primitives)
  - [Constants](#constants)
  - [Arrays and Slices](#arrays-and-slices)
  - [Maps and Structs](#maps-and-structs)
  - [If and Switch Statements](#if-and-switch-statements)
  - [Looping](#looping)
  - [Defer, Panic, and Recover](#defer-panic-and-recover)
  - [Pointers](#pointers)
    - [Creating pointers](#creating-pointers)
    - [Dereferencing pointers](#dereferencing-pointers)
      - [Pointer arithmetic](#pointer-arithmetic)
    - [The new function](#the-new-function)
    - [Working with nil](#working-with-nil)
    - [Types with internal pointers](#types-with-internal-pointers)
  - [Functions](#functions)
    - [Parameters](#parameters)
      - [Passing by value](#passing-by-value)
      - [Passing by reference](#passing-by-reference)
      - [Variadic Parameters](#variadic-parameters)
    - [Return values](#return-values)
      - [Named return](#named-return)      
      - [Return multiple results](#return-multiple-results)
    - [Anonymous functions](#anonymous-functions)
    - [Methods](#methods)
  - [Interfaces](#interfaces)
    - [Basics](#basics)
    - [Composing interfaces](#composing-interfaces)
    - [Type conversion](#type-conversion)
        - [Type empty interface](#type-empty-interface)
        - [Type switches](#type-switches)
    - [Implementing with values vs pointers](#implementing-with-values-vs-pointers)
    - [Best practices](#best-practices)
  - [Goroutines](#goroutines)
  - [Channels](#channels)
  - [Other resources to check](#other-resources-to-check)

</details>

## Introduction

- developed at Google by Robert Griesemar, Rob Pike and Ken Thompson

Go:
- is strong and statically typed
- has an excellent community
- stands out through: simplicity, fast compile times, garbage collected (you can manage the memory, but you don't have to), 
built-in concurrency, compile to standalone binaries (go runtime, go dependency libraries are compiled in here - you get easier 
version management at runtime)

## Setting Up a Development Environment

### Setting up in Ubuntu

It is not necessary to set `GOROOT` if it is the default path.

Can have a compound `GOPATH` to differentiate between own code and third party libraries.  All segments of the `GOPATH` 
are searched for source code.

> cd ~
> vim .bashrc

    # Not necessary if default location
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOROOT/bin

    # Used by go get for storing third party libraries
    export GOPATH=/home/username/golib
    export PATH=$PATH:$GOPATH/bin
    # Own code here
    export GOPATH=$GOPATH:/home/username/code

> source ~/.bashrc

### Important directories in own workspace

A workspace in go needs a `/src` directory.

Binaries will be put in `/bin` directory.

`pkg` - for intermediary binaries that don't need to be recompiled every time

## Variables

## Primitives

## Constants

## Arrays and Slices

## Maps and Structs

## If and Switch Statements

## Looping

## Defer, Panic, and Recover

## Pointers

### Creating pointers

`a` and `b` are value types and `b` is a copy of the value from `a` (they don't point to the same memory location):

```
a := 1
b := 2
fmt.Println(a, b) // 1 2
a = 3
fmt.Println(a, b) // 3 2
```

Change `b` into a pointer:

```
var a int = 1
// b is a pointer to an integer and points to a
var b *int = &a
// *b - dereferencing operator
fmt.Println(a, *b, &a, b) // 1 1 0xc000018030 0xc000018030
a = 2
fmt.Println(a, *b, &a, b) // 2 2 0xc000018030 0xc000018030
*b = 3
fmt.Println(a, *b, &a, b) // 3 3 0xc000018030 0xc000018030
```

### Dereferencing pointers

On previous example `*` in `*b` is a dereferencing operator. At runtime finds the pointer, the memory location and returns 
the value stored there.

Both `a` and `*b` point to the same data, they point to the same memory location.

#### Pointer arithmetic

```
a := [3]int{1, 2, 3}
b := &a[0]
c := &a[1]
fmt.Printf("%v %p %p\n", a, b, c) // [1 2 3] 0xc000016018 0xc000016020
```

Difference in memory between `b` and `c` is 2 bytes, coming from the `a` array which contains integers, 2 bytes apart.

Go doesn't allow pointer arithmetic (`c := &a[0] + 4`).

To do this, go has [unsafe package](https://pkg.go.dev/unsafe). This contains operations that the go runtime will not check.

### The new function

You can work with the pointer without caring where the underlying data is stored just pointing to where it is:

```
func main() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	// Can use new() to initialize to 0
	// ms = new(myStruct) // &{0}
	fmt.Println(*ms)
}

type myStruct struct {
	foo int
}
```

### Working with nil

A pointer not initialized it is initialized by go:

```
var ms *myStruct
fmt.Println(*ms) // <nil>
```

### Types with internal pointers

Get to the underlying field by dereferencing the ms pointer:

```
func main() {
	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 1
	fmt.Println((*ms).foo) // 1
	
	// Works the same as (*ms).foo because of the limitations put on pointers
	// The pointer ms doesn't have a field foo, it points to a structure that does have a field foo but this is syntactic sugar
	ms.foo = 1
	fmt.Println(ms.foo) // 1
}

type myStruct struct {
	foo int
}
```

Arrays are independent of each other:

```
a := [3]int{1, 2, 3}
b := a
fmt.Println(a, b) // [1 2 3] [1 2 3]
a[1] = 22
fmt.Println(a, b) // [1 22 3] [1 2 3]
```

Slices are not:

```
a := []int{1, 2, 3}
b := a
fmt.Println(a, b) // [1 2 3] [1 2 3]
a[1] = 22
fmt.Println(a, b) // [1 22 3] [1 22 3]
```

Same behavior for maps:

```
a := map[string]string{"foo": "bar", "baz": "buz"}
b := a
fmt.Println(a, b) // map[baz:buz foo:bar] map[baz:buz foo:bar]

a["foo"] = "qux"
fmt.Println(a, b) // map[baz:buz foo:qux] map[baz:buz foo:qux]
```

## Functions

### Parameters

```
func main() {
    sayGreeting("Hello", "Stacey")
}

// syntactic sugar - (param1, param2 same_type_for_both)
// type inferred for first param
func sayGreeting(greeting, name string) {
    //...
}
```

#### Passing by value

A copy of the value is sent to the sayGreeting function.

```
func main() {
    greeting := "Hello"
    name := "Stacey"
    sayGreeting(greeting, name)
    fmt.Println(name) // Stacey
}

func sayGreeting(greeting, name string) {
    name = "Ted"
    fmt.Println(name) // Ted
}
```

#### Passing by reference

The address of name is sent to the sayGreeting function.

Sometimes passing the pointer is more efficient than passing the value. For e.g., when you need to pass a large data structure.

```
func main() {
    greeting := "Hello"
    name := "Stacey"
    sayGreeting(&greeting, &name)
    fmt.Println(name) // Ted
}

func sayGreeting(greeting, name *string) {
    *name = "Ted"
    fmt.Println(name) // Ted
}
```

#### Variadic Parameters

THere can be only one in a function definition, and it is the last parameter.

```
// values is a slice with the values from the arguments
func sum(values ...int) { 
    result := 0
    
    for _, v := range values {
        result += v
    }
    
    //...
}
```

### Return values

#### Return local variable as a pointer

```

func main() {
    s := sum(1, 2, 3, 4, 5)
    fmt.Println(*s)
}

func sum(values ...int) *int { 
    result := 0
    
    for _, v := range values {
        result += v
    }
    
    return &result
}
```

`result` is defined on the execution stack of `sum`. The execution stack is destroyed when the function exits. Go recognizes 
`result` address is returned and so it promotes the variable to the shared memory / heap.

#### Named return

Syntactic sugar for declaring the result variable which is implicitly returned

```
func sum(values ...int) (result int) { 
    //...
    
    // don't need to specify result
    return
}
```

#### Return multiple results

```
func main() {
    d, err := divide(5.0, 0.0)
    
    if err != nil {
        fmt.Println(err)
        return
    } 

    fmt.Println(d)
}

func divide(a, b, float64) (float64, error) {
    if b == 0.0 {
        return 0.0, fmt.Errorf("Cannot divide by 0")
    }
    
    return a / b, nil
}
```

### Anonymous functions

```
func main() {
    func() {
        fmt.Println("abc")
    }() // invoke the anonymous function
    
    var f1 func() = func() {
        fmt.Println("abc")
    }
    f1()
    
    var f2 func(string) = func(a string) {
		fmt.Println(a)
	}
	f2("abc")
}
```

When using outer scope variables, better to pass them as arguments. This way, changes in the outer scope won't influence the inner 
scope. Evident issues when it comes to asynchronous code.

```
func main() {
    for i := 0; i < 5; i++ {
        func(i int) {
            fmt.Println(i)
        }(i)
    }
}
```

Example of anonymous functions with parameters and returns:

```
func main() {
    var divide func (float64, float64) (float64, error)
    
    divide = func(a, b float64) (float64, error) {
        if b == 0.0 {
            return 0.0, fmt.Errorf("Cannot divide by 0")
        } else {
            return a / b, nil
        }
    }
    
    d, err := divide(5.0, 3.0)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(d)
}
```

### Methods

A method is a function executing in a known context (any type). Can create a type for an int and add methods on that type.

```
type greeter struct {
	greeting string
	name     string
}

// Method
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// Method
func (g *greeter) anotherGreet() {
	fmt.Println(g.greeting, g.name)
	g.name = "Modified"
}

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()
	
	g.anotherGreet()
	fmt.Println(g.name) // Modified
}
```

`greeter` is specified as a value type in `greet()` - this is a value receiver (a copy of the type is passed to the function).

`anotherGreet()` receives a pointer (pointer receiver) and it is able to modify 

## Interfaces

### Basics

```
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (n int, err  error) {
	n, err = fmt.Println(string(data))
	return
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, Go!"))
}
```

Usually, one method interface is called the method name + `ER` (`write` => `writER`).

Any type can have a method associated with it and can implement an interface.

```
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func main() {
	// Need to cast an integer to an IntCounter
	myInt := IntCounter(0)

	var inc Incrementer = &myInt

	for i:= 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}
```

### Composing interfaces

Defining a composite interface:
```
type WriterCloser interface {
	Writer
	Closer
}
```

As long as you implement all the methods on the embedded interfaces then you implement the composed interface as well.

```
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello, Go! abcdefghijklmnopqrstuvwxyz"))
	wc.Close()
}

// NewBufferedWriterCloser Basically a constructor
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}


	return n, nil
}

func (bwc BufferedWriterCloser) Close() error {
	// Flush the rest of the buffer, prints the last 8 or fewer characters
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)

		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}

	return nil
}
```

### Type conversion

Can use type conversion to change `wc` from a `WriterCloser` to a `BufferedWriterCloser`. This way, `bwc` has access to 
the internal fields of the `BufferedWriterCloser`.

```
var wc WriterCloser = NewBufferedWriterCloser()

bwc := wc.(*BufferedWriterCloser)
```

Can check if the conversion is failing instead of panicking if the conversion is unsuccessful:

```
// import "io"

var wc WriterCloser = NewBufferedWriterCloser()

r, ok := wc.(io.Reader)
if ok {
    fmt.Println(r)
} else {
    fmt.Println("Conversion failed")
}
```

#### Type empty interface

= interface with no methods

```
type EmptyInterface interface {}

// or:

var myObj interface {} = NewBufferedWriterCloser()
if wc, ok := myObj.(WriterCloser); ok {
    wc.Write([]byte("Hello, Go! abcdefghijklmnopqrstuvwxyz"))
    wc.Close()
}

r, ok := myObj.(io.Reader)
if ok {
    fmt.Println(r)
} else {
    fmt.Println("Conversion failed")
}
```

Anything can be cast into an object that has no methods on it (including primitives).

As there are no methods on an empty interface this usually is an intermediary step and another conversion is needed to be 
able to use that type.

#### Type switches

Commonly used with the empty interface to list the types expected and based on the actual type add the logic of processing them.

```
var i interface {} = 0

switch i.(type) {
case int:
    fmt.Println("i is an integer.")
case string:
    fmt.Println("i is a string.")
default:
    fmt.Println("I don't know what i is.")
}
```

With the object itself:

```
	var i interface {} = 0

	switch v := i.(type) {
	case int:
		fmt.Println("i is an integer - " , v) // i is an integer - 0
	case string:
		fmt.Println("i is a string.")
	default:
		fmt.Println("I don't know what i is.")
	}
```

### Implementing with values vs pointers

The method set of a type determines the methods that can be called on an object of that type.

Every type has a method type associated with it. This can be possible empty.

[Definition](https://go.dev/ref/spec#Method_sets):
```
The method set of a defined type T consists of all methods declared with receiver type T.

The method set of a pointer to a defined type T (where T is neither a pointer nor an interface) is the set of all methods declared with receiver *T or T.

The method set of an interface type is the intersection of the method sets of each type in the interface's type set (the resulting method set is usually just the set of declared methods in the interface).
```

### Best practices

- use many, small interfaces
- don't export interfaces for types that will be consumed 
  - export concrete types when creating a library
  - don't assume you know how people are going to use it, allow them to create an interface that these types will implement (they will implement only what they need)
- do export interfaces for types that will be used by package
  - whoever is using this library can create a concrete type for this interface, containing exactly what you need for the implementation
- design functions and methods to receive interfaces whenever possible

## Goroutines

## Channels

## Other resources to check
- [Autocomplete daemon for Go](https://github.com/nsf/gocode) (not maintained anymore)