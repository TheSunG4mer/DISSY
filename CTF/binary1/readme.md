# Binary 1

The programming language C provides a variety of library functions for manipulating strings. One such example is the `gets(s)` function which reads a line from stdin and stores it in the string `s`.
This function does not check whether the line given as input actually fits within the length of the string, and will simply continue writing the bytes after where `s` is stored in memory.

In C the local variables for a function are stored together on the stack. Overwriting other local variables may allow you to alter the behaviour of the program.
There are many resources available online explaining the layout of the call stack in C.

The system you are going to attack has a server running a simple C program which takes a username as input.
It then checks whether the admin flag is 0, concluding that the user is not the admin if this is the case.

To connect to the sever you can use netcat as follows:
```
nc chals.syssec.dk 13334
```

## Your goal
Your goal is to provide an input causing the server to treat you as if you were the admin.

## Files
- `chal.c` contains the server code.
- `challenge` is the executable binary obtained by compiling `chal.c`. (This is included if you want to reverse the binary yourself. Note that the flag is redacted, the binary with the real flag is running on the server. So despite being compiled with the same flags, there might be slight differences in offsets in the two files.)
- `example.py` contains sample code for connecting and sending data to the server using the pwntools python package.
- 
## What to hand in

This exercise is one of the three options for the first hand-in.

If you choose to hand in this exercise you must complete the challenges "Binary 1" and "Binary 2", writing a report covering the three following points:

1. Describe your attack in detail, including all the actions you took to attack the system and any relevant code you might have used.
Include the flags which you recovered from the server.

2. Describe why your attack worked. Explain how it was possible for your attack to circumvent the security measures of the system.

3. Suggest one or more ways the system might be improved to prevent your attack.