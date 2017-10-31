# gobase16to64
Convert strings with base-16 chars into Base64 encoding

Mini project to convert base16 strings into Base64.

I decided to make it when I wanted to concatenate two SHA1 hashes but with a 63 char limit.
I figured there'd be fewer collisions if I converted to Base64 encoding first.

I tried to make this as fast as possible, but there's probably plenty of improvements that could be made.
