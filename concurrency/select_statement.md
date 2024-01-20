The select statement in Go looks like a switch statement but for channels 
The select statement lets a go-routine wait on multiple communication 

Select blocks until any of the case statements are ready 
If multiple case statements are ready, then it selects one at random and proceeds 

syntax 
```
select {
    case channel_send_or_receive:
      // Do something
    case channel_send_or_receive:
      // Do something
    default:
      // (optional)
}
```

The select statement lets a go routine wait on multiple communication operations

Select along with channels and go routines becomes a very powerful tool for managin synchronization and concurrency 
