Channels are a means throguh which different go routines communicate 
"Do not communicate by sharing memory, instead, share memory by communicating" - Rob Pike 
Share memory by communicating - Go-routines and channels 

The communication is bidirections by default, meaning tha tyou can send an recieve values from the same channel 
By default, channels sends and recieve until the other side is ready 
This allows go-routines to synchronize without explicit locks or condition variables 

```
    var c chan string 
    c := make(chan string)
```

Channel Operation 
- Sending a value 
- Receiving a value 
- Closing a channel 
- Querying Buffer of a channel 
- Querying length of a channel 

``` 
Sneding a value 
ch <- v
```

<- is a channel send operator 
This operator is used to send a value to the channel 
v must be a value which is assignable to the element type of channel ch 

```
Recieving a value 
val := <- ch
```
<- is a channel receibing operator 
this is used to receive a value from a channel 
val is a variable in which the read data from the channel will be stored 
```
Close Operation 
close(ch)
```
close() is a built-in function 
The argument of a close function call must be a channel value 


```
Querying buffer of a channel 
cap(ch)
```
cap() is a built-in function 
return an integer denoting the buffer of the specified channel 

```
Querying length of a channel 
len(ch)
```

len() is a built-in function 
return an integer denoting the length of the specified channel 