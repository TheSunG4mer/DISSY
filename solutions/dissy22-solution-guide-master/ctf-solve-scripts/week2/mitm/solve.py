import os
import socket
import random
import hashlib
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad

# 1024-bit MODP group from RFC 2409
PRIME_P = 0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE65381FFFFFFFFFFFFFFFF
GENERATOR_G = 2

SERVER = "chals.syssec.dk"
ALICE_PORT = 13000
BOB_PORT = 13001


def dh_generate_keypair(p, g):
    """Return (public, private) DH keypair (random small private like original)."""
    priv = random.randint(2, (p - 1) // 2)
    pub = pow(g, priv, p)
    return pub, priv


def dh_shared_key_md5(p, their_pub, my_priv):
    """Compute shared secret, then derive 128-bit key using MD5 (preserve original)."""
    if their_pub < 3 or their_pub > p - 2:
        # original printed and exited; keep that behavior to avoid changing protocol flow
        print("your pubkey looks really suspicious o.0")
        exit()
    shared_int = pow(their_pub, my_priv, p)
    return hashlib.md5(shared_int.to_bytes(128, "big")).digest()


def aes_cbc_encrypt(key, data_bytes):
    iv = os.urandom(16)
    cipher = AES.new(key, AES.MODE_CBC, iv)
    ct = cipher.encrypt(pad(data_bytes, 16))
    return iv.hex() + ":" + ct.hex()


def aes_cbc_decrypt(key, payload):
    iv_hex, ct_hex = payload.split(":")
    iv = bytes.fromhex(iv_hex)
    ct = bytes.fromhex(ct_hex)
    cipher = AES.new(key, AES.MODE_CBC, iv)
    pt = cipher.decrypt(ct)
    try:
        return unpad(pt, 16)
    except Exception:
        # keep original behavior: hide padding errors and return 16 NUL bytes
        return b"\x00" * 16


# --- open sockets (same ordering as original) ---
alice_sock = socket.create_connection((SERVER, ALICE_PORT))
bob_sock = socket.create_connection((SERVER, BOB_PORT))

# Intercept Alice's first message (contains Alice's public key)
alice_first = alice_sock.recv(1024).decode()
alice_pub = int(alice_first.split(": ")[1].split("\n")[0], 16)
print(f'Alice intercepter pubKey: {alice_pub}')

# Create forged DH keypair for MITM
forged_pub, forged_priv = dh_generate_keypair(PRIME_P, GENERATOR_G)
formatted_forged = hex(forged_pub)[2:]
print(f'Formatted fake pubKey: {formatted_forged}')

# Compute the shared secret that Alice will derive for our forged key (we impersonate Bob)
alice_ss = dh_shared_key_md5(PRIME_P, alice_pub, forged_priv)

# Send forged public to Alice (impersonating Bob)
alice_sock.sendall(f'{formatted_forged}\n'.encode())

# Receive Alice's encrypted message (with the shared secret she computed using our forged pub)
alice_response = alice_sock.recv(1024).decode()
print(f"Alice server response: {alice_response}")

# Bob initial request (discarded by original)
bob_resp_1 = bob_sock.recv(1024).decode()

# Send forged public to Bob (impersonating Alice)
bob_sock.sendall(f'{formatted_forged}\n'.encode())

# Now intercept Bob's public key
bob_resp_2 = bob_sock.recv(1024).decode()
bob_pub = int(bob_resp_2.split(": ")[1].split("\n")[0], 16)
print(f'Bob intercepted pubKey: {bob_pub}')

# Compute the shared secret Bob will derive (we impersonate Alice)
bob_ss = dh_shared_key_md5(PRIME_P, bob_pub, forged_priv)

# Extract ciphertext Alice->Bob from the earlier alice_response and decrypt it using alice_ss
alice_ct = alice_response.split(": ")[1].split("\n")[0]
alice_plain = aes_cbc_decrypt(alice_ss, alice_ct)
print(f"Decrypted message from Alice to Bob: {alice_plain}")

# Re-encrypt that plaintext under Bob's shared secret and forward to Bob
forward_to_bob = aes_cbc_encrypt(bob_ss, alice_plain)
print(f"Encrypted message from Bob to Alice: {forward_to_bob}")
bob_sock.sendall(f'{forward_to_bob}\n'.encode())

# Intercept Bob's response to Alice
bob_response_3 = bob_sock.recv(1024).decode()
print(f"Bob server response: {bob_response_3}")

bob_ct = bob_response_3.split(": ")[1].split("\n")[0]
bob_plain = aes_cbc_decrypt(bob_ss, bob_ct)
print(f"Decrypted message from Bob to Alice: {bob_plain}")

# Re-encrypt under Alice's secret and forward to Alice
reply_to_alice = aes_cbc_encrypt(alice_ss, bob_plain)
print(f"Encrypted message from Alice to Bob: {reply_to_alice}")
alice_sock.sendall(f'{reply_to_alice}\n'.encode())

# Intercept Alice's next response
alice_response_2 = alice_sock.recv(1024).decode()
print(f"Alice server response: {alice_response_2}")

alice_ct2 = alice_response_2.split(": ")[1].split("\n")[0]
alice_plain2 = aes_cbc_decrypt(alice_ss, alice_ct2)
print(f"Decrypted message from Alice to Bob ->: {alice_plain2}")

# Re-encrypt for Bob and forward
forward2_to_bob = aes_cbc_encrypt(bob_ss, alice_plain2)
print(f"Encrypted message from Bob to Alice: {forward2_to_bob}")
bob_sock.sendall(f'{forward2_to_bob}\n'.encode())

# Intercept Bob's following response
bob_response_4 = bob_sock.recv(1024).decode()
print(f"Bob server response: {bob_response_4}")

bob_ct2 = bob_response_4.split(": ")[1].split("\n")[0]
bob_plain2 = aes_cbc_decrypt(bob_ss, bob_ct2)
print(f"Decrypted message from Bob to Alice: {bob_plain2}")

# Close sockets (same as original)
alice_sock.close()
bob_sock.close()
