from base64 import encode
from flask import Flask, send_from_directory, jsonify, render_template, request, make_response, g
import os, json
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
import time
app = Flask(__name__)


# Takes a plaintext bytestring, and encrypts it to a hex encoded ciphertext using aes_cbc
def aes_cbc_encrypt(plaintext, key):
    iv = os.urandom(16)
    cipher = AES.new(key, AES.MODE_CBC, iv=iv)
    plaintext_padded = pad(plaintext, 16)
    ciphertext = iv + cipher.encrypt(plaintext_padded)
    cipher_hex = ciphertext.hex()
    return cipher_hex


# Takes a hex encoded IV + Ciphertext, and decrypts it to a bytestring using aes_cbc
def aes_cbc_decrypt(ciphertext, key):
    iv, ct = ciphertext[:32], ciphertext[32:]
    cipher = AES.new(key, AES.MODE_CBC, iv=bytes.fromhex(iv))
    plaintext_padded = cipher.decrypt(bytes.fromhex(ct))
    plaintext = unpad(plaintext_padded, 16)
    return plaintext


@app.route('/')
def index():
    return render_template('index.html')

@app.route('/store')
def store():
    return render_template('store.html')

# New and improved padding oracle resistant input handling!
@app.route('/submitdata', methods=['POST'])
def submitdata():
    submitted_ct = request.form['submitted_data']
    print(f'submitted value is {submitted_ct}')
    # Decrypt encrypted data
    submitted_pt = aes_cbc_decrypt(submitted_ct, key)
    # do something with this data
    return send_from_directory('.', 'received.html')

# Military grade security lets us even show what other people submitted, without leaking what the data they stored was!
# To prove how confident i am in this system, i'll even share the encrypted flag i submitted to the service earlier!
@app.route('/flag')
def flag():
    return render_template('flag.html', hexflag = enc_flag)


if __name__ == "__main__":
    key = os.urandom(16)
    with open("flag.txt", "rb") as f:
        flag = f.read().strip()

    enc_flag = aes_cbc_encrypt(flag, key)
    app.run(host='0.0.0.0', port=os.getenv('PORT'))