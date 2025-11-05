
import sys,os
from tqdm import tqdm
import requests

TargetUrl = "http://chals.syssec.dk:14000/submitdata"

# ciphertext = "932a113dee60f69b3448e9a6b1d75e8e869f4746c6a1d71b4ef62e73c8e813efdd89f2248694f486850e53a7c71ad10b9659ab0a9318f4335aa28160d1d85b27"


ciphertext = "111f9f1f87684efd39abecfa9cf58f9c73bba9dfa4b69561ca0f02596f27981141fb2676af5a5f2bcdcca6fc8eab2b8a65d382277965c2d6937a3c23c1561a95"


IV = bytes.fromhex(ciphertext[:32])
ciphertext = ciphertext[32:]


BLOCK_SIZE = 128
BYTE_NB = BLOCK_SIZE//8

# key = '0123456789abcdef'
# IV = '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'

import json
def timerequests(message):
    payload = {}
    payload["submitted_data"] = message
    resp = requests.post(TargetUrl, data=payload)
    timetaken = resp.elapsed.total_seconds()
    # print("doing")
    # print(resp.status_code)
    return resp.status_code == 200


def determine(payload):
    if timerequests(payload):
        return True
    else:
        return False
# def determine(payload):
#     if timerequests(payload) > 1:
#         return True
#     else:
#         return False

# Determine if the message is encrypted with valid PKCS7 padding
def oracle(encrypted):
    # print(type(encrypted))
    res = determine(encrypted.hex().zfill(32))
    # print(f"result = {res}")
    return res



##########################################
# Padding Oracle Attack Proof of Concept #
##########################################

def poc(encrypted):
    block_number = len(encrypted)//BYTE_NB
    decrypted = bytes()
    # Go through each block
    for i in range(block_number, 0, -1):
        current_encrypted_block = encrypted[(i-1)*BYTE_NB:(i)*BYTE_NB]

        # At the first encrypted block, use the initialization vector if it is known
        if(i == 1):
            previous_encrypted_block = IV
        else:
            previous_encrypted_block = encrypted[(i-2)*BYTE_NB:(i-1)*BYTE_NB]
 
        bruteforce_block = previous_encrypted_block
        current_decrypted_block = bytearray(b'\x00'*16)
        padding = 0

        # Go through each byte of the block
        for j in tqdm(range(BYTE_NB, 0, -1)):
            padding += 1

            # Bruteforce byte value
            for value in range(0,256):
                bruteforce_block = bytearray(bruteforce_block)
                bruteforce_block[j-1] = (bruteforce_block[j-1] + 1) % 256
                joined_encrypted_block = bytes(bruteforce_block) + current_encrypted_block

                # Ask the oracle
                if(oracle(joined_encrypted_block)):
                    current_decrypted_block[-padding] = bruteforce_block[-padding] ^ previous_encrypted_block[-padding] ^ padding
                    print(f'current progress: {current_decrypted_block}')
                    
                    # Prepare newly found byte values
                    for k in range(1, padding+1):
                        bruteforce_block[-k] = padding+1 ^ current_decrypted_block[-k] ^ previous_encrypted_block[-k]

                    break

        decrypted = bytes(current_decrypted_block) + bytes(decrypted)
        print(f'current progress: {decrypted}')
    return decrypted[:-decrypted[-1]]  # Padding removal

#### Script ####

usage = """
Usage:
  python3 poracle_exploit.py <message>         decrypts and displays the message
  python3 poracle_exploit.py -o <hex code>     displays oracle answer
Cryptographic parameters can be changed in settings.py
"""

if __name__ == '__main__':
    print(poc(bytes.fromhex(ciphertext)).decode("ascii"))
    # if len(sys.argv) == 2 : #chiffrement
    #     if len(sys.argv[1])%16!=0:       # code size security
    #         print(usage)
    #     else:
    #         print("Decrypted message: ", poc(bytes.fromhex(sys.argv[1])).decode("ascii"))
    # elif len(sys.argv) == 3 and sys.argv[1] == '-o' : #oracle
    #     print(oracle(bytes.fromhex(sys.argv[2])))
    # else:
    #     print(usage)