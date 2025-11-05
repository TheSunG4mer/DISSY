# Cracking bad randomness

When using randomness to secure systems it is important to ensure that random values we generate are actually as unpredictable as we need them to be. Many randomness generators are sufficiently "random looking" for non-cryptographic purposes, but fall far short of the mark if used for cryptography.

This can be for a number of reasons, in some cases the number of possible seeds may be small enough to be exhaustively searched. In other cases it may be possible to predict the next "random" value, given a long enough sequence of previous outputs.

## The system

In this challenge you will be targeting the password reset mechanism of a website.
When a password reset is requested the website will send an email with a password reset token.
This token is the hexadecimal encoding of 128bits of randomness produced by the standard python randomness library. 
You have been given a zip file containing 128 emails with reset tokens.
The emails have been generated in sequence with email0 being produced first, followed by email1 and so on.
All tokens have been generated in sequence from the same instance of the python random number generator.

The python random number generator uses the [Mersenne Twister](https://en.wikipedia.org/wiki/Mersenne_Twister), which naturally outputs 32 bits at a time.
Given sufficent outputs it is possible to predict future outputs of the Mersenne Twister, helpfully others have already implemented this: [randcrack](https://github.com/tna0y/Python-random-module-cracker).

## Your goal
Your goal is to successfully reset the password of the admin by predicting the reset token which will be sent to them in the next password reset email.

[[password reset website]](http://chals.syssec.dk:8081/resetpassword)

## What to hand in

This exercise is one of the three options for the first hand-in.

If you choose to hand in this exercise you must complete this challenge and write a report covering the three following points:

1. Describe your attack in detail, including all the actions you took to attack the system and any relevant code you might have used.
Include the flags which you recovered from the server.

2. Describe why your attack worked. Explain how it was possible for your attack to circumvent the security measures of the system.

3. Suggest one or more ways the system might be improved to prevent your attack.