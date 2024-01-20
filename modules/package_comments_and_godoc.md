Godoc 

Parses Go Source code - including comments - and produces documentation as HTML or plain text 
The result is documentation tightly coupled with the code it documents 

The convention is simple: To document a type, variable, constant, function, or even a package 
    Write a regular comment directly preceding its declaration, with no intervening blank line 
    Use a blank comment to break your comment into multiple paragraphs 

If you have lengthy comments for the package, the convention is to put the comments in a file in your package called doc.go

```
go doc PACKAGE_NAME
go doc PACKAGE_NAME.IDENTIFIER_NAME
```