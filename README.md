# go-playground
Learning about [Go (golang)](https://go.dev/)

## Notes
Go has a minimalistic and simple syntax similar to C **except** that it has built in garbage collection where we don't have to manually manage memory (no usages of `maclloc()`, `free()` that coud lead to memory issues like dangling pointers, segfaults, etc.)

Go supports some object-orientated principles: encapsulation, polymorphism, abstraction, and composition, but doesn't have classes or inheritance. 
* Encapsulation: 
    * Capitalized names -> exported (public) e.g. `var ApiKey = 123abc`
    * Lowercase names -> unexported (private) e.g. `var pathUrl = 123abc`
* Polymorphism:
    * Achieved with interfaces ("duck typing" - "if it walks and quacks like a duck, then it must be a duck)
* Abstraction:
    * Achieved with interfaces and unexported (private) fields/methods
* Composition:
    * Achieved with struct embedding
    * "has-a" releationship instead of traditional "is-a"

Go doesn't have classes, but it has struct and methods. It also doesn't have constructors or method overloading. And instead of inheritance, it uses composition (struct embedding, see above).

Go doesn't use an external build tool (like Gradle/Maven). The built in Go toolchain handles dependency resolution and builds with `go mod` and `go build` respectively.

Go also comes with its own standard testing package `testing`. It's good enough for unit tests so we don't have to bring in another library (like JUnit). It runs with `go test`. But, there are 3rd party frameworks that could add more expression and capabilities like Testify (rich asserts/mocks), Ginkgo (behavior-driven development), GoMock (mocking library).
