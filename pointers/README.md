# Memory Management

- Memory alloaction & deallocation happens automatically

## new()
- Allocate Memory but no INIT
- You will get a memory address
- zeroed storage

## make()
- Alloacte memory and INIT
- you will get a memory address
- non-zeroed storage

## GC - garbage collection
- GC happens automatically -> package runtime
- default GOGC = 100
- GOGC = off
- disable the garbage collection entirely.
- NUMCPU() return no of logical CPUs usable by the current process.


