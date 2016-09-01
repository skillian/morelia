# morelia
Python implementation written in Go

## Why?
1.  No more reference counting:  a Python implementation written in Go does not need reference counting because the Go runtime handles the memory management!  In addition to cleaning up the code, this also will speed up multi-threading because modifying reference counting does not need to use interlocked instructions that really hurt performance under contention.
2.  The infamous GIL:  The CPython implementation of Python has a global interpreter lock that limits the interpreter to executing only one system thread of Python code at a time (non-Python threads running in wrapped libraries won't necessarily hold this GIL).  The only reason that I am yet aware of for having the GIL is to ensure memory is properly managed.  Because of Go's garbage collection, this should be more easily worked around.
3.  Goroutines.  I am curious if it will be feasible or not to implement Python's asyncio library with goroutines.  Hopefully, we'll see.

## How?
So far with plagiarism, of course!  I'm basing the implementation very heavily upon the CPython implementation, just removing the reference counting and the "ob_type" struct member and using Go's interfaces instead.

## What works:
Nothing!  I'm still writing the Python runtime implementation.  So far, I have a few tests that compare Ints using a "CPython-like" Compare function.
