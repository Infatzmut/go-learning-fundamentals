Go code is grouped into packages, and packages are grouped into modules 
A module specifies the dependencias needed to run our code, including the Go version and athe set of other modules it requires in the go.mod file 

A colleciton of Go source code becomes a module when there's a valid go.mod file in its root directory 
It consists of the module declaration, the minium compatible version  for Go, and the dependencies for the imported third-party packages 

If go mod tidy fails because it can not find the package, we can use the replace directive 
```
    go mod edit -replace <path_to_replace> <module_location>
```