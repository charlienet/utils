package crypto_test

import (
	"encoding/hex"
	"testing"

	"github.com/charlienet/utils/crypto"
	"github.com/tjfoc/gmsm/sm4"
)

func TestGmsmSm4(t *testing.T) {
	key, _ := hex.DecodeString("0123456789ABCDEFFEDCBA9876543210")
	msg, _ := hex.DecodeString("F0A2B07E64DD2C2590F93E4EDD90FBB4")

	c, err := sm4.Sm4Ecb(key, msg, true)
	t.Log(hex.EncodeToString(c), err)

	d, err := sm4.Sm4Ecb(key, c, false)
	t.Log(hex.EncodeToString(d), err)
}

func TestPadding(t *testing.T) {
	msg, _ := hex.DecodeString("F0A2B07E64DD2C2590F93E4EDD90FBB4")

	blockSize := sm4.BlockSize
	padding := blockSize - len(msg)%blockSize
	t.Log(padding)
}

func TestSm4(t *testing.T) {
	key := []byte("1234567890abcdef")
	msg := []byte("123321123321123321123321123321")

	cipherText, err := crypto.Sm4(key).ECB().Encrypt(msg)
	t.Log("ECB加密:", hex.EncodeToString(cipherText), err)

	de, err := crypto.Sm4(key).ECB().Decrypt(cipherText)
	t.Log("ECB解密:", string(de), err)

	cipherText, err = crypto.Sm4(key).CBC().WithIV(key).Encrypt(msg)
	t.Log("CBC加密:", hex.EncodeToString(cipherText), err)

	de, err = crypto.Sm4(key).CBC().WithIV(key).Decrypt(cipherText)
	t.Log("CBC解密:", string(de), err)
}
