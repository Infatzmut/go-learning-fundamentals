A channel that needs a receiver as soon as a message is emitted to the channel 
We do not declare a capacity 
Have some capacity to hold data
On a buffered channel
    - Sending to a channel, blocks the go-routine, only if the buffer is full 
    - Receiving from a channel blocks only when the channel is empty

Syntax `c:= make(chan <data_type>, capacity)`
``` c:= make(chan int, 10)```

Length of an unbuffered channel 
Builtin len() functin can be used to get the length of a channel 
The length of a channel is the number of elements that are already there in the channel 
So, length represents the number of elements queued in the buffer of the channel 
Length of a channel is always less than or equal to the capacity of the channel (length <= capacity>)
Length of unbuffered is always zero 
