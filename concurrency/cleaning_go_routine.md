go-routine-leak 

Whenever you launch a go-routine funciton, you must make sure that it will eventually exit
A go-routine that would never terminate, forever occupies the memry it has reserved. This kin of memory leak is called go-routine leak 

Go routine leak if they end up either blocked forever on I/O like channel communicaion or fall into infinite loops 
