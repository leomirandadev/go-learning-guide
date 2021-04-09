# This is the guide of how to learn Golang

**Install golang: [GOLANG OFFICIAL](https://golang.org/)** 

First, you should know the basic characteristics of go.

Go is compiled. If you need test your application before the build you can use:
```shell
  $ go run nameOfMainFile.go
```

For basic build you can use:
```shell
  $ go build nameOfMainFile.go
```

The main file always must have a main function like
```go
  package main // this is the name of package

  func main(){ // func is the prefix of every function on go
    // this is the main function
  }
```

Packages is normally separated by directories and the functions are share with all the files. Like:

```
  project
    └── main.go
    └── anotherFile.go
```

if your main file has the function called "calc" you can access this function on anotherFile.go using:
```go
  // It's anotherFile.go
  package main // this is the name of package

  func anotherFunction(){
    calc() // 
  }
```

# PRACTICAL EXAMPLES OF THE BASE

  - [Types in go](types)
  - [Functions in go](functions)
  - [Pointers in go](pointers)
  - [Interfaces in go](interfaces)
  - [Errors in go](errors)
  - [Tests in go](tests)
  - [Routines in go](routines)
  - [Channels in go](channels) (WRITING)

# USING DAY BY DAY

  - [HTTP communication in go](API)
  - [Data Base communication in go](DB) (WRITING)
  - [The go mod](goMod) (WRITING)

# MORE
  - [Clean Architecture with Golang](https://github.com/leomirandadev/clean-architecture-go)
    - ref: [https://www.youtube.com/watch?v=Yg_ae0UvCv4](https://www.youtube.com/watch?v=Yg_ae0UvCv4)
  - [gRPC with Golang and NodeJS](https://github.com/leomirandadev/grpc-project)
    - ref: [https://www.youtube.com/watch?v=BdzYdN_Zd9Q](https://www.youtube.com/watch?v=BdzYdN_Zd9Q)

  - [Messager Broker NATS](https://github.com/leomirandadev/nats-playground)
    - ref (pt-BR): [https://www.youtube.com/watch?v=NBi0r7QOJSs](https://www.youtube.com/watch?v=NBi0r7QOJSs)