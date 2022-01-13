package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strconv"
)



type rsaInstance struct {
	hashOptions
	prk *rsa.PrivateKey
	puk *rsa.PublicKey
}

type rsaOption func(o *rsaInstance) error

func NewRsa(h Hash, opts ...rsaOption) (*rsaInstance, error) {
	o := &rsaInstance{}

	sh := crypto.Hash(h)
	if !sh.Available() {
		return nil, errors.New("unknown hash value " + strconv.Itoa(int(h)))
	}

	o.h = sh

	for _, f := range opts {
		if err := f(o); err != nil {
			return nil, err
		}
	}

	// 未设置私钥时随机生成密钥
	if o.prk == nil {
		prk, err := rsa.GenerateKey(rand.Reader, defaultRsaBits)
		if err != nil {
			return nil, err
		}

		o.prk = prk
	}

	// 公钥未设置时从私钥导出
	if o.puk == nil {
		o.puk = &o.prk.PublicKey
	}

	return o, nil
}

func ParsePKCS1PrivateKey(p []byte) rsaOption {
	return func(o *rsaInstance) error {
		block, _ := pem.Decode(p)
		if block == nil {
			return errors.New("failed to decode private key")
		}

		prk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return err
		}

		o.prk = prk

		return nil
	}
}

func ParsePKIXPublicKey(p []byte) rsaOption {
	return func(o *rsaInstance) error {
		block, _ := pem.Decode(p)
		if block == nil {
			return errors.New("failed to decode public key")
		}

		k, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return err
		}

		puk := k.(*rsa.PublicKey)

		o.puk = puk

		return nil
	}
}

func (o *rsaInstance) Encrypt(msg []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, o.puk, msg)
}

func (o *rsaInstance) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, o.prk, ciphertext)
}

func (o *rsaInstance) Sign(msg []byte) ([]byte, error) {
	hashed := o.getHash(msg)
	sign, err := rsa.SignPKCS1v15(rand.Reader, o.prk, o.h, hashed)
	return sign, err
}

func (o *rsaInstance) Verify(msg, sign []byte) bool {
	hashed := o.getHash(msg)
	if err := rsa.VerifyPKCS1v15(o.puk, o.h, hashed, sign); err != nil {
		return false
	}

	return true
}
