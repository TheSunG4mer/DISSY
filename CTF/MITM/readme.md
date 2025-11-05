# I'm looking at the man-in-the-middle
As you have seen in the TA session exercises "man-in-the-middle" attacks are a major concern when trying to design an authenticated key exchange protocol.

Now it is your turn to be the person in the middle!

## The system

Alice and Bob are trying to talk to each other, their mutual friend Charlie who has taken a course on cryptography told them it was a good idea to do something called a Diffie-Hellman key exchange to keep their communication secret.
Charlie was on a skiing holiday during the authenticity week of the cryptography course, and did therefore not learn why signatures are important for security.

Alice and Bob have agreed to generate a key to communicate by following this protocol:

1. Alice generates chooses a secret exponent $a$ and computes $G^{a}$, she then sends $G^{a}$ to Bob.

2. Bob then generates his own secret exponent $b$ and computes $G^{b}$, sending this back to Alice.

3. Both parties may now compute $G^{ab} = (G^a)^b = (G^b)^a$, and hash this to obtain a key $K = H(G^{ab})$.
The key $K$ is then used to encrypt the subsequent messages using AES in CBC mode.

You have been given the code of Alice and Bob.

Alice can be found by connecting to 
```
nc chals.syssec.dk 13000
```
and Bob can be found at
```
nc chals.syssec.dk 13001
```

## Your goal
The AKE protocol above does not ensure that Alice and Bob are actually talking to each other.
Your goal is to act as an intermediary between Alice and Bob making them think they are talking to each other.
If you do this successfully you will be able to decrypt and forward the messages sent between them.
The flag you are trying to obtain will be sent somewhere in their subsequent communication.

## Files
- `alice.py` and `bob.py` provide the source code Alice and Bob are using to communicate. 
It may be helpful to reuse parts of their code when playing a role in the key exchange.

## What to hand in

This exercise is one of the two options for the second hand-in.

If you choose to hand in this exercise you must complete this challenge and write a report covering the three following points:

1. Describe your attack in detail, including all the actions you took to attack the system and any relevant code you might have used.
Include the flags which you recovered from the server.

2. Describe why your attack worked. Explain how it was possible for your attack to circumvent the security measures of the system.

3. Suggest one or more ways the system might be improved to prevent your attack.



