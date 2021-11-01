package main

import (
	"encoding/json"	
	"os"
	"bufio"
	"encoding/hex"
)

type AssumedRoleCredentials struct {
    AccessKeyId   		string
    SecretAccessKey 	string
	SessionToken 		string
	Expiration 			string
}

type AssumedRoleUser struct {
    AssumedRoleId   	string
    Arn 				string
}

type AssumedRole struct {
    Credentials   		AssumedRoleCredentials 	`json:"Credentials"`
    User 				AssumedRoleUser 		`json:"AssumedRoleUser"`
}

type Profile struct {
	Name				string
	RoleArn 			string					`json:"role_arn"`		
	AssumedRole			AssumedRole				`json:"assumed_role"`
}

type Config struct {
    AccessKeyId   		string					`json:"aws_access_key_id"`
    SecretAccessKey 	string					`json:"aws_secret_access_key"`
	MfaSerial			string					`json:"mfa_serial"`	
	Profiles			[]Profile				`json:"profiles"`
}

func json_EncryptJson(passphrase string, to_encrypt string) ([]byte, error) {
	data := []byte(to_encrypt)
    key := GenerateKey(passphrase)
    ciphertext, err := Encrypt(key, data)
	return ciphertext, err
}

func json_DecryptJson(passphrase string, to_decrypt []byte) ([]byte, error) {
    key := GenerateKey(passphrase)
    cleartext, err := Decrypt(key, to_decrypt)
	return cleartext, err
}

func json_WriteConfig(initDir string, passphrase string, config_json []byte) (int) {
	f, err := os.Create(initDir + "/aws-context.json")
    check(err)
    defer f.Close()

	w := bufio.NewWriter(f)
	encrypted, err := json_EncryptJson(passphrase, string(config_json))
    check(err)

    n4, err := w.WriteString(hex.EncodeToString(encrypted))
	check(err)

	w.Flush()

	return n4
}

func json_ReadConfig(initDir string, passphrase string) (Config) {
	hextext, err := os.ReadFile(initDir + "/aws-context.json")
    check(err)

	ciphertext := make([]byte, hex.DecodedLen(len(hextext)))
	hex.Decode(ciphertext, hextext)
    check(err)
	
	plainbytes, err := json_DecryptJson(passphrase, ciphertext)
	check(err)

    config := Config{}
    json.Unmarshal(plainbytes, &config)

	return config
}

