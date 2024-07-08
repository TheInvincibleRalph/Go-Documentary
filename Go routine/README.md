## Deadlock

In an operating system, a deadlock occurs when a process or thread enters a waiting state because a requested system resource is held by another waiting process, which in turn is waiting for another resource held by another waiting process. If a process remains indefinitely unable to change its state because resources requested by it are being used by another process that itself is waiting, then the system is said to be in a deadlock.




## Coffman conditions (MHNC)

A deadlock situation on a resource can arise only if all of the following conditions occur simultaneously in a system:

1. Mutual exclusion: At least one resource must be held in a non-shareable mode (we are assuming that one resource could have multiple instances); that is, only one process at a time can use the resource. Otherwise, the processes would not be prevented from using the resource when necessary. Only one process can use the resource at any given instant of time.

2. Hold and wait or resource holding: a process is currently holding at least one resource and requesting additional resources which are being held by other processes.

3. No preemption: a resource can be released only voluntarily by the process holding it.

4. Circular wait: each process must be waiting for a resource which is being held by another process, which in turn is waiting for the first process to release the resource. In general, there is a set of waiting processes, P = {P1, P2, ..., PN}, such that P1 is waiting for a resource held by P2, P2 is waiting for a resource held by P3 and so on until PN is waiting for a resource held by P1.

## Detection

Under the deadlock detection, deadlocks are allowed to occur. Then the state of the system is examined to detect that a deadlock has occurred and subsequently it is corrected. An algorithm is employed that tracks resource allocation and process states, it rolls back and restarts one or more of the processes in order to remove the detected deadlock. Detecting a deadlock that has already occurred is easily possible since the resources that each process has locked and/or currently requested are known to the resource scheduler of the operating system.

After a deadlock is detected, it can be corrected by using one of the following methods:

1. Process termination: one or more processes involved in the deadlock may be aborted. One could choose to abort all competing processes involved in the deadlock. This ensures that deadlock is resolved with certainty and speed. But the expense is high as partial computations will be lost. Or, one could choose to abort one process at a time until the deadlock is resolved. This approach has a high overhead because after each abort an algorithm must determine whether the system is still in deadlock. Several factors must be considered while choosing a candidate for termination, such as priority and age of the process.

2. Resource preemption: resources allocated to various processes may be successively preempted and allocated to other processes until the deadlock is broken.

---[Wikipedia](https://en.wikipedia.org/wiki/Deadlock)


