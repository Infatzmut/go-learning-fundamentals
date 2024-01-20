We can find out the number of logical processors using the runtime.Numcpus method 
Go routines can be considered as a lightweight application-level thread that has a separate independent execution
The go runtime has its onw scheduler taht will multiplex the go-routines on the OS level threads in the go runtime
It schedules an arbitrary number of go-routines onto an arbitrary number of OS threads(m:n multiplexing)

Go scheduler is a cooperative scheduler
is a style of scheduling in which the OS never interups a running process to initate a context switch from one process to another 
Processes must voluntarily yield control periodically or when logically blocked on a resource 
Of course, theere are some specific check points where go-routine can yield its execution to other go-routine. These are called context switches

Example of context switching 
 - Functions call 
 - Garbage Collection 
 - Network Calls 
 - Channel operations 
 - On using go keyword 

Go routines vs Threads 
-Go rotuine are cheaper 
Go-routines are multiplexed to a fewer number of OS threads 
The context switching time of go routines is much faster 
Go routines communicate using channels 