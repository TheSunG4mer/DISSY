from Crypto.Cipher import AES
from Crypto.Util.Padding import pad,unpad
import hashlib 
import random
import os
# Secret messages which alice and bob will send to each other. You need to recover these by attacking the handshake!
from messages import alice1, bob1, alice2, bob2

# Publicly agreed on Diffie Hellman parameter (https://www.rfc-editor.org/rfc/rfc2409#page-22)
# 1024 bit prime
p = 0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF
G = 2

# perform DH key exchange
def generate_DH_pubkey(p,G):
    your_sk = random.randint(2,(p-1)//2)
    your_pubkey = pow(G,your_sk,p)
    return your_pubkey, your_sk

def compute_DH_shared_secret(p,their_pubkey, your_sk):
    # Check for clearly broken public keys
    if their_pubkey < 3 or their_pubkey > p-2:
        print("your pubkey looks really suspicious o.0")
        exit()
    # compute shared secret and convert to bytes
    ss_int = pow(their_pubkey, your_sk, p)
    ss = hashlib.md5(ss_int.to_bytes(128,"big")).digest()
    return ss

# Encrypt a ciphertext with AES-CBC using the shared secret
def encrypt_message(ss, message):
    iv = os.urandom(16)
    aes = AES.new(ss,AES.MODE_CBC,iv)
    m_padded = pad(message,16)
    ct = aes.encrypt(m_padded)
    payload = iv.hex() + ":" + ct.hex()
    return payload

# Decrypt a ciphertext with AES-CBC using the shared secret
def decrypt_message(ss,ciphertext):
    iv, ct = ciphertext.split(':')
    iv = bytes.fromhex(iv)
    ct = bytes.fromhex(ct)
    aes = AES.new(ss,AES.MODE_CBC,iv)
    ct = aes.decrypt(ct)
    try:
        ct = unpad(ct,16)
    except:
        # Handle unpad errors silently to avoid padding oracle attacks
        return b'\x00'*16
    return ct


# Alice and bob perform a key exchange by sharing their pubkeys
alice_pk, alice_sk = generate_DH_pubkey(p,G)
print(f"Alice pubkey: {hex(alice_pk)[2:]}")
bob_pubkey = int(input(f"Bob pubkey:"),16)
ss = compute_DH_shared_secret(p,bob_pubkey,alice_sk)

# With a share dsecret, they can securely exchange encrypted messages
alice1_enc = encrypt_message(ss,alice1)
print(f'alice_to_bob_1: {alice1_enc}')

bob1_enc = input("bob_to_alice_1:")
bob1_dec = decrypt_message(ss,bob1_enc)

if not bob1_dec == bob1:
    print("That doesn't sound like you Bob! We've been compromised???")
    print("I'm out of here!")
    exit()

alice2_enc = encrypt_message(ss,alice2)
print(f'alice_to_bob_2: {alice2_enc}')

bob2_enc = input("bob_to_alice_2:")
bob2_dec = decrypt_message(ss,bob2_enc)

if not bob2_dec == bob2:
    print("That doesn't sound like you Bob! We've been compromised???")
    print("I'm out of here!")
    exit()

print("Thanks! Great talking to you")