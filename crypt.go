package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"


	// "encoding/hex"

)

func Encrypt(key, data []byte) ([]byte, error) {
    blockCipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(blockCipher)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = rand.Read(nonce); err != nil {
        return nil, err
    }

    ciphertext := gcm.Seal(nonce, nonce, data, nil)

    return ciphertext, nil
}

func Decrypt(key, data []byte) ([]byte, error) {
    blockCipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(blockCipher)
    if err != nil {
        return nil, err
    }

    nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }

    return plaintext, nil
}

func GenerateKey(password string) ([]byte) {
	bPassword := []byte(password)

	key := make([]byte, 32)
	for i := 0; i < len(password); i++ {
		if i < 32 {
			key[i] = bPassword[i]
		} 
	}

    return key
}

// func main() {
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Printf("ciphertext: %s\n", )

//     plaintext, err := Decrypt(key, ciphertext)
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Printf("plaintext: %s\n", plaintext)
// }