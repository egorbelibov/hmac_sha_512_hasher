# HMAC-SHA-15 Hasher

Simple `Go` script that hashes a (message & key) with the **HMAC-SHA-512** functions and generates a **base16 & 64 output**. 

## Usage
Run the script with:
```console
go run main.go
```
Then, when asked, provide a `message` & `key`. 

By default, your input is invisble, but if you want to see it you can run the script with an additional argument: 
```console
go run main.go show-input
```
---
## Sample Output
With `show-input` flag:
```
===========================
HMAC Hasher to Base 16 & 64
===========================
Message: message
Key: key
===========================
==> Base 16: 
e477384d7ca229dd1426e64b63ebf2d36ebd6d7e669a6735424e72ea6c01d3f8b56eb39c36d8232f5427999b8d1a3f9cd1128fc69f4d75b434216810fa367e98
==> Base 64: 
5Hc4TXyiKd0UJuZLY+vy0269bX5mmmc1Qk5y6mwB0/i1brOcNtgjL1QnmZuNGj+c0RKPxp9NdbQ0IWgQ+jZ+mA==
```