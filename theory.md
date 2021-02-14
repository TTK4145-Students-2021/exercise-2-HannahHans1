Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> An atomic operation must either complete fully or not at all. When an atomic operation reads or modifies a resource, no other operation will be able to read or change it.

### What is a critical section?
> The sections of the code where the shared variables can be accessed.

### What is the difference between race conditions and data races?
> A race condition occurs when the result of an operation depends on the ordering of events. A data race is a situation where two threads access the same variable, and at least one of the operation tries to modify the variable.

### What are the differences between semaphores, binary semaphores, and mutexes?
> A semaphore is a variable/data type assosiated with a set of atomic operation which is used to control access to a shared variable. signal() increments the semaphore by one, and wait() decrements it by one. A semaphore cannot have a negative value, so if wait() is called on a zero-valued semaphore the function will wait until another operation signal()s the semaphore. A binary semaphore is a semaphore which can only take the values 0 and 1. A mutex is another way to control access to a shared variable. As opposed to the semaphore, only the operation that locked a resource can unlock it. ("**mut**ual **ex**clusion") 

### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> Channels are used to pass messages between goroutines. Mailboxes is also used for message passing, but a mailbox can store different kinds of messages and chooses between them. A queue is a linear data structure and can be either LIFO, FIFO or a priority queue.

### List some advantages of using message passing over lock-based synchronization primitives.
> - The programmer does not need to keep track of locks and protection.
  - Avoid race condotion

### List some advantages of using lock-based synchronization primitives over message passing.
> - Faster than message passing
