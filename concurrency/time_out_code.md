- Most interactive program must return a resopnse within a certain amount of time

Blocking timeout in select can be achieved by using After function of time package

syntax
```
func After(d Duratio) <- chan Time
```

The After function waits for d duration to finish and then it returns the current time on a channel