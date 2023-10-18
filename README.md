# 100DaysofCode-Go
100 days of code challenge, in Golang


## Basics

1. Packages
  - Go applications are organised in packages.
  - A package is a collection of source files located in the same directory.
  - All source files in a directory must share the same package name.
2. Functions
  - Functions accept zero or more parameters
  - Function returns use the `return` keyword.
  ```Golang
  package greeting
  // A public function 
  func Hello (name string) string {
    return hi(name)
    }
  // A private function
  func hi(name string) string {
    return "Hi " + name
    }
  ```
  - Multi-line comments are inserted using `/*` and `*/`
  - Single-line comments are inserted using `//`
3. Variables
  - Go variables are statically-typed, similar to Rust.
  - Variables are defined by explicitly specifying a type 
    - var test int
  - Can also be defined using an initializer, the compiler will assign the var type
    - test := 10
  - Once declared, vars can be assigned with `=`
  - A var cannot change type in its lifetime.
    ```Golang
    count := 100 // Create a var and assign initial value
    count = 101 // Change the value to 101 
    count = "test" // This will fail while compiling due to type change
    ```
  - Constants are similar to variables, but are immutable during the lifetime.
    ```Golang
    const Age = 21 // Defines a numeric constant 
    ```
