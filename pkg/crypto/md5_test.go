package crypto_test

import (
	"testing"
	"b2yun/pkg/crypto"
)

func TestMD5(t *testing.T) {

	crypto := crypto.NewCrypto()

	md5 := crypto.MD5Crypto()

	testString := "00010001"

	hash, err := md5.Salt(testString)

	if err != nil {
		t.Error(err)
	}

	t.Error(hash)

	ok, err := md5.Compare(hash, testString)

	if !ok {
		t.Error("MD5解密结果不符预期")
	}

	testString2 := testString + "testString"

	ok, err = md5.Compare(hash, testString2)
	if ok {
		t.Error("MD5解密结果不符预期")
	}

}
