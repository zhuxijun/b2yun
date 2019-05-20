package crypto

// Crypto 加密
type Crypto struct {
	md5 *MD5Crypto
}

// NewCrypto 返回新的加密对象
func NewCrypto() *Crypto {

	var crtpto Crypto

	crtpto.md5 = newMD5Crypto()

	return &crtpto

}

// MD5Crypto 返回md5加密对象
func (c *Crypto) MD5Crypto() *MD5Crypto {
	return c.md5
}
