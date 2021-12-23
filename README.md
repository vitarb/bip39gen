# bip39gen

bip39gen is a simple utility for generating crypto secure BIP39 passphrases of arbitrary length.

## Notes
Each word provides 11 bits of entropy (2048 words = 2^11).
You should use at least 64 bits (or 6 words) if you want reasonable protection against brute force attacks.

## Usage
Build the binary:
```bash
make
```

Run it passing desired number of words to be generated (defaults to 12):
```bash
./bip39gen -n 6
```
