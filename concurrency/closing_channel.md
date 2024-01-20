closing a channel means that no more data can be sent to that channel 
It is generally closed when there's no more data to be sent 
We can use the inbuilt close() function for the operation

syntax `v, ok := <-ch`

if ok is true, this means that the channel is open
if ok is false, this means that the channel is closed and there are no more values to receive

