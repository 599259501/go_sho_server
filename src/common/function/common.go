package function

import (
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func GetRandStr(length int)(string,error){
	if length == 0 {
		return "",nil
	}
	clen := len(StdChars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return "",err
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = StdChars[c%clen]
			i++
			if i == length {
				return string(b),nil
			}
		}
	}

}

func AesEncryStr(str string, encryption string)(string,error){
	c, err := aes.NewCipher([]byte(encryption))
	if err!=nil{
		return "",err
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(str))
	cfb.XORKeyStream(ciphertext, []byte(str))

	return base64.StdEncoding.EncodeToString(ciphertext),nil
}

func AesDecryStr(str string, encryption string)(string,error){
	c, err := aes.NewCipher([]byte(encryption))
	if err!=nil{
		return "",err
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(str))

	cfbdec.XORKeyStream(plaintextCopy, []byte(str))
	return string(plaintextCopy),nil
}
