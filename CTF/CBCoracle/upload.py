from time import time
import requests


TargetUrl = "http://chals.syssec.dk:14000/submitdata"

ciffertext = "111f9f1f87684efd39abecfa9cf58f9c73bba9dfa4b69561ca0f02596f27981141fb2676af5a5f2bcdcca6fc8eab2b8a65d382277965c2d6937a3c23c1561a95"

# Here's the basic script i use when i need to back up data!
def upload_data(encrypted_bytes):
    encrypted_data = encrypted_bytes.hex()
    payload = {}
    payload["submitted_data"] = encrypted_data
    resp = requests.post(TargetUrl, data=payload)
    return resp.status_code

blocks = []
for i in range(4):
    blocks.append(list(bytes.fromhex(ciffertext[32 * i : 32 * (i + 1)])))
# print(type(blocks[0][0]))

def test_block_for_correct_padding(block1, block2):
    combined_data = block1 + block2
    byte_data = bytes(combined_data)
    return upload_data(byte_data) == 200

def XOR_lists(L1, L2):
    return [x1 ^ x2 for x1, x2 in zip(L1, L2)]

def make_padding_list(n):
    return [0] * (16 - n + 1) + [n] * (n - 1)

# print(XOR_lists([0,1,2], [3,5,7]))

# for i in range(1, 17):
#     print(make_padding_list(i))

answer = []

for padding_block, data_block in zip(blocks[:3], blocks[1:]):
    answer_block = [0] * 16
    for position in range(15, -1, -1):
        padding_number = 16 - position
        pre_padding = XOR_lists(make_padding_list(padding_number), answer_block)
        # print(f"Current pre padding: {pre_padding}.")
        for guess in range(256):
            pre_padding[position] = guess
            current_padding_block = XOR_lists(padding_block, pre_padding)
            if test_block_for_correct_padding(current_padding_block, data_block):
                answer_block[position] = guess ^ padding_number
                print(f"Found valid entry {bytes([guess ^ padding_number])} at position {position}")
                
        # else:
        #     print("Went through all options, and nothing worked")
        #     break
    answer += answer_block
    print(bytes(answer_block))
    
print(f"Final answer: {bytes(answer)}")
    



# encrypted_data = bytes.fromhex("ab"*32)

# for x in range(256):
#     n = upload_data(bytes([0] * 15 + [x] + [0] * 16))
#     if n == 200:
#         print(x)






# def XOR_bytes(b1, b2):
#     return bytes(a ^ b for a, b in zip(b1, b2))
# print(len(ciffertext) / 2)

# plaintext = ""

# cifferbytes = bytes.fromhex(ciffertext)

# print(cifferbytes)

# while len(cifferbytes) > 16:
#     new_block = bytes.fromhex("0" * 32)
#     # print(new_block)
#     for i in range(1, 17):
#         padding = bytes([0] * (16 - i + 1) + [i] * (i - 1))
#         padding = XOR_bytes(padding, new_block)
        
        
#         front = bytes([0] * (len(cifferbytes) - 32))
#         back = bytes([0] * 16)
#         padding = b''.join([front, padding, back])
#         base = XOR_bytes(cifferbytes, padding)
#         # print(padding)
        
        
#         for x in range(256):
#             guess_xor = bytes([0] * (len(cifferbytes) - 32 + 16 - i) + [x] + [0] * (16 + i - 1))
#             n = upload_data(XOR_bytes(base, guess_xor))
#             if n == 200:
#                 new_block = new_block[:16-i] + XOR_bytes(bytes([x]), bytes([i])) + new_block[16-i+1:]
#                 print("Found new char :", new_block)
#                 break
#         else:
#             print("Found no match...")
#     plaintext = str(new_block) + plaintext
#     cifferbytes = cifferbytes[:-16]
#     print("Plaintext now looks like:", plaintext)
# print(plaintext)