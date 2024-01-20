From writing and reading a file, to writing and reading an HTTP response, to processing databas operations, all processes are based on I/O 
The io standard library aims to abstract these existing processes and tasks, which appear in countless places, into a shared interface 

io interfaces 
- The Reader interface
- The Writer interface 

-> The io Reader interface 
Basically provides the input function 

``` 
type Reader interface {
    Read(p [] byte) (n int, err error)
}
```

-> io Writer interface 
basically provides a generic way to output data 
```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
