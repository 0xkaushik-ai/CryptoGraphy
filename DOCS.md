- go mode define the module and go version 
- aes is the block cipher package 
- gf.go is the gf2** finite-fiel arithmetic,the foundation byte addition (XOR),xtime and multiplication. Else else calls into this 
- The S-box (and its inverse) — the substitution table, built from the field inverse in gf.go. This is  AES'S non-linearity . 
-  key.go Key expansion turns the 16/24/32 byte key into all the per round keys .
- cipher the round operation and the full block encrypt/decrypt that chains them. 


- mode using the cipher on real data (package mode)
- ctr mode turn the block cipher into a stream cipher so you can encrytp any length,first apperance of the nonce/counter 
- GCM-ctr plus an authentication tag,giving autheticate encryption (Encrypt+tamper detection)