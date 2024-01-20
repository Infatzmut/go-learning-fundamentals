In programming, file handling essentially means working with it, sucha as getting metadata information, creating new files, or simply reading and writing data to and from a file 
In Golang, the API for file handling is well-knitted into the standard architecture and we can do it using the standard libraries

- File handling libraries 
The os package provides an API interface for file handling which is uniform across all operating systems 
It provides functionality such as creation, deletion, opening a file, modifying its permissions and so on 
The io package provides interfaces for basic i/o primitives and wraps them into easy-to-use public interfaces
The filepath package that would provide functions to parse construct file paths 
We also use the fmt package to format I/O with functions to read and write to the standard input and output 


- Construction file paths 
-> Using the filepath package 
-> It provides functions to parse and construct file paths in a way that is portable between operating systems 
-> dir/file on Linux vs dir\file on Windows, for example 

--> Summary 
    - Join method that constucts paths in a portable way 
    - Dir and Base method that split a path to the directory and the file 
    - IsAbs method to check if a path is absolute or not 
    - Ext method to get the extension from a filename 
