from Crypto.Util.Padding import pad,unpad
from Crypto.Cipher import AES
import hashlib 
import random
import os
import socket


p = 0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF
G = 2

SERVER = "chals.syssec.dk"
ALICE_PORT = 13000
BOB_PORT = 13001


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



# Create fake keypair
E_pk, E_sk = generate_DH_pubkey(p, G)

# Establish connection with Alice and Bob
alice_sock = socket.create_connection((SERVER, ALICE_PORT))
bob_sock = socket.create_connection((SERVER, BOB_PORT))


# Initiate communication with Alice
alice_first = alice_sock.recv(1024).decode()
# print(f'Alice intercepter pubKey: {alice_first}')
alice_pk = int(alice_first.split(": ")[1].split("\n")[0],16)
# print(alice_pk)

formated_pk = hex(E_pk)[2:]
alice_sock.sendall(f"{formated_pk}\n".encode())

alice_shared_key = compute_DH_shared_secret(p, alice_pk, E_sk)


# Initiate communication with Bob

bob_first = bob_sock.recv(1024).decode()
# print(bob_first)
bob_sock.sendall(f"{formated_pk}\n".encode())
bob_second = bob_sock.recv(1024).decode()
# print(bob_second)
bob_pk = int(bob_second.split(": ")[1].split("\n")[0],16)

bob_shared_key = compute_DH_shared_secret(p, bob_pk, E_sk)

# Begin communication:

from_alice_raw_message_1 = alice_sock.recv(1024).decode()
print(from_alice_raw_message_1)
alice_decrypted_message_1 = decrypt_message(alice_shared_key, from_alice_raw_message_1.split(": ")[1].split("\n")[0])
print(alice_decrypted_message_1)

to_bob_encrypted_message_1 = encrypt_message(bob_shared_key, alice_decrypted_message_1)
bob_sock.sendall(f"{to_bob_encrypted_message_1}\n".encode())

from_bob_raw_message_2 = bob_sock.recv(1024).decode()
print(from_bob_raw_message_2)
bob_decrypted_message_2 = decrypt_message(bob_shared_key, from_bob_raw_message_2.split(": ")[1].split("\n")[0])
print(bob_decrypted_message_2)

to_alice_encrypted_message_2 = encrypt_message(alice_shared_key, bob_decrypted_message_2)
alice_sock.sendall(f"{to_alice_encrypted_message_2}\n".encode())


from_alice_raw_message_3 = alice_sock.recv(1024).decode()
print(from_alice_raw_message_3)
alice_decrypted_message_3 = decrypt_message(alice_shared_key, from_alice_raw_message_3.split(": ")[1].split("\n")[0])
print(alice_decrypted_message_3)

to_bob_encrypted_message_3 = encrypt_message(bob_shared_key, alice_decrypted_message_3)
bob_sock.sendall(f"{to_bob_encrypted_message_3}\n".encode())

from_bob_raw_message_4 = bob_sock.recv(1024).decode()
print(from_bob_raw_message_4)
bob_decrypted_message_4 = decrypt_message(bob_shared_key, from_bob_raw_message_4.split(": ")[1].split("\n")[0])
print(bob_decrypted_message_4)

# to_alice_encrypted_message_2 = encrypt_message(alice_shared_key, bob_decrypted_message_2)
# alice_sock.sendall(f"{to_alice_encrypted_message_2}\n".encode())
