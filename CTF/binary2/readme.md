# Binary 2
This exercise should only be attempted after completing "Binary 1".

In an attempt to improve the security of the system the code in `chal.c` has been changed.
Now rather than whether the `isadmin` flag is set to 0, the code instead checks whether `isadmin != 1`. This means the code will only think a user is the admin exactly when `isadmin == 1`.

To connect to the sever you can use netcat as follows:
```
nc chals.syssec.dk 13333
```

## Your goal
Your goal is once again to provide an input causing the server to treat you as if you were the admin.

## Files
- `chal.c` contains the server code.
- `challenge` is the executable binary obtained by compiling `chal.c`. (This is included if you want to reverse the binary yourself. Note that the flag is redacted, the binary with the real flag is running on the server. So despite being compiled with the same flags, there might be slight differences in offsets in the two files.)
- `example.py` contains sample code for connecting and sending data to the server using the pwntools python package.

## What to hand in

This exercise is one of the three options for the first hand-in.

If you choose to hand in this exercise you must complete the challenges "Binary 1" and "Binary 2", writing a report covering the three following points:

1. Describe your attack in detail, including all the actions you took to attack the system and any relevant code you might have used.
Include the flags which you recovered from the server.

2. Describe why your attack worked. Explain how it was possible for your attack to circumvent the security measures of the system.

3. Suggest one or more ways the system might be improved to prevent your attack.