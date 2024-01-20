In Go language, panic is just like an exception, it also arises at runtime
Panic means an unexpected condition arises in your Go program due to which the execution of your program is terminated
There are few scenarios that can cause panic while working with channels such as 
- Sending to a channel after it has been closed
- Closing an already closed channel 

