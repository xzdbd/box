package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	salt1 := "x!z@d#b$d"
	salt2 := "b^o&x"

	io.WriteString(h, salt1)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	return last
}
