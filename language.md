
Message passing seems like a better choice than shared variables. With several elevators, buttons, and lights, I think it will be problematic to keep track of all locks. We need some kind of message passing for the networking part, so why not use the same technique to move data as well. Also, it seems easier to prevent errors to propagate when memory is not shared.

Go uses channels for message passing and has libraries for networking. There is also a lot of useful information on how to use Go, both in general and to this specific project. So far, I also think that the Go syntax is intuitive and quite similar to languages I already know.


