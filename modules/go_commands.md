go mod init 
    the command initializaes and writes a new go.mod file in the current directory 
    In effect creating a new module rooted at the current directory 
    It accepts one optional argument, which is the module path for the new module

go mod tidy 
    This command ensures that the go.mod file matches the source code in the module 
    It adds any missing module requirements necessary to build the current module's packages and dependencies
    and it removes requirements on modules that don't provide any relevant packages

go run <file>
    Compiles and runs the program 
    Internally it compiles your code and builds an executable binary in a temporary location, launches that temp exe-file and finally cleans it when your app exits
    Is usefull for testing small programs during the initial development stage of your application 

go build 
    Compiles the packages named by the import paths, along with their dependencies into an executable
    The executable gets created in the current source directory

go install 
    Compile and move the executable to $GOPATH/bin
    run this executalbe from any path on the terminal 

go get
    Resolves its command-line arguments to packages at specific module version
    updates go.mod to require those versions
    Downloads source code into the module cache 
    We can use this command to add a dependency for a package or upgrade it to its latest version ``` go get example.com/pkg```
    Or we can downgrade a package to a specific version  ```go get example.com/pkg@1.2.3```