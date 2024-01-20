Workflow

Design and code the packages that the module will include 
Commit code to your repository using conventions that ensure it's available to others via Go tools
Publish the module to make it discoverable by developers
Over time, revise the module with version that use a version numbering convention that signals each version's stability and backward compatibility 

Desing and Development 

When you're designing a module's public API, try to kkeps its functionality focused and discrete
Ensure backwar compatibility


Decentralized publishing 
In Go, you publish your module by tagging its code in your repository to make it available for other developers to use 
Developers use Go tools(including the go get command) to download your module's code to compile with 

Package discovery 
After you've published your module and some has fetched it with Go tools, it will become visible on the Go package discovery site at pkg.go.dev 
There, developers can search the site to find it and read its documentation

Versioning 
As you revise and improve your module over time, you assign version numbers, designed to signal each version's stability and backward compatibility 
Your indicate a module's version number by tagging the modules source in the repository with a number 

