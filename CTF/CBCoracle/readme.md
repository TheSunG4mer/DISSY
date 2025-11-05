# Padding Oracles

When designing systems using cryptography it is very important to avoid leaking information about secret parameters, such as keys and message plaintexts.
One potent example of how relatively little information may be exploited with disastrous consequences is the case of CBC padding oracles.

When using AES in CBC mode the plaintext must be padded to line up with a whole number of blocks.
One approach to solving this is PKCS#7 padding.
If the last block is missing a single byte, a byte with the value 0x01 will be added, if two bytes are needed the bytes 0x02,0x02 will be added, and so on with 0x03,0x03,0x03 for three bytes...

If a system leaks whether the padding of a block is valid this may be used to extract the plaintext.
This may occur if a server either returns an error message depending on the padding, or may be possible to learn by a timing attack if the server bails-out-fast when the padding is found to be invalid.
Recall for CBC mode if you want to alter what a bit decrypts to you can change the bit in the same position in the previous block.
For extra details and helpful diagrams see [wikipedia](https://en.wikipedia.org/wiki/Padding_oracle_attack).

The plaintext can be found by trying to make the last byte of the block decrypt to the value 0x1, this can be detected as the server will not return a padding error.
Finding this value will at most take 255 attempts.
Using what you have learned you may then attempt to modify the last two bytes so they decrypt to 0x02,0x02.
This allows you to learn a whole block of plaintext using at most 16*255 = 4080 queries.

## The system

The server you will be attacking has a ciphertext encrypted using AES in CBC mode available here [[ciphertext]](http://chals.syssec.dk:14000/flag).

It helpfully also has the option to submit ciphertexts on [[submit]](http://chals.syssec.dk:14000/submitdata), which will return an internal server error with a different message depending on whether an error has ocurred. (With the `unpad` function throwing such an exception if it fails due to incorrect padding.)

## Your goal
Your goal is to decrypt the ciphertext by using the responses given when you submit modified ciphertexts. 
The plaintext you recover will contain the flag.
You can divide your attack into two steps:
    
1. Build a padding oracle, this could be a function which given a ciphertext queries http://chals.syssec.dk:14000/submitdata and uses this to determine a boolean return value which is true if the padding was valid.
2. Using your padding oracle you can then implement the full attack to recover the plaintext. 


## Files
- `server.py` contains parts of the server code showing how your requests are handled.
- `upload.py` contains code for sending a ciphertext to the server in a post request.

## What to hand in

This exercise is one of the two options for the second hand-in.

If you choose to hand in this exercise you must complete this challenge and write a report covering the three following points:

1. Describe your attack in detail, including all the actions you took to attack the system and any relevant code you might have used.
Include the flags which you recovered from the server.

2. Describe why your attack worked. Explain how it was possible for your attack to circumvent the security measures of the system.

3. Suggest one or more ways the system might be improved to prevent your attack.



