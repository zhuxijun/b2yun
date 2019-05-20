package root

// Cryptor 加密接口
type Cryptor interface {
	// Salt 加盐
	Salt(s string) (string, error)
	// Compare 是否相等
	Compare(hash string, s string) (bool, error)
}
