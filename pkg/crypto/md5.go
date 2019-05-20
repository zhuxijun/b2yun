package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Crypto MD5加密
type MD5Crypto struct{}

// newMD5Crypto 返回新的MD5加密对象
func newMD5Crypto() *MD5Crypto {
	return &MD5Crypto{}
}

// Salt 加密
func (m *MD5Crypto) Salt(s string) (string, error) {

	h := md5.New()

	h.Write([]byte(s))

	cipher := h.Sum(nil)

	return hex.EncodeToString(cipher), nil

}

// Compare 对比hash字符串是否正确
func (m *MD5Crypto) Compare(hash string, s string) (bool, error) {

	var b bool

	cipherStr, _ := m.Salt(s)

	if cipherStr == hash {
		b = true
		return b, nil
	}
	return false, nil

}
