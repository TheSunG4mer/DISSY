import random

def generate_password_reset_token():
    token = hex(random.getrandbits(128))[2:].zfill(128//4)
    return token