 The problem with go-routines was the main go-rouine terminating before the go-routines compoleted or even began their execution 
 To wait for multiple go-routines to finish, we can use a wait group 
 A wait group is a synchronization primitive that allows multiple go-routines to wait for each other 
 This package acts like a counter that blocks execution in a structured way until its internal counter becomes 0

 Syntax 
 ```
 import "sync"
 var wg sync.WaitGroup
 ```

 This have three methods 

 wg.Add -> Indicates the number of available go-routines to wait for
 wg.Wait -> Blocks the execution of code untils the internal counter reduces to 0 
 wg.Done -> This method decreses the interal count parameter in Add method by 1 


