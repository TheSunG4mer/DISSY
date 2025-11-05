x = bytes([1,2,3,4])
print(x)
print(x[2])
x = x[:2] + bytes([10]) + x[3:]
print(x)