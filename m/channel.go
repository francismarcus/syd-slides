ch := make(chan int)

ch <- 18  // Send to channel
v := <-ch // Receive from channel.
