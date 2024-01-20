When to use buffered channels?

- Proper use of buffered channel means that you must handle the case where the channel is block, which might happen due to waiting on sender/receiver

- Buffered channels are useful when you know how many go-routines you have launched, want to limit the nuber of go-routines you will launch, or want to limit the amount of work that is queued up.