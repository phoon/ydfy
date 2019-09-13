package ydfy

import (
	"crypto/md5"
	"fmt"
	"time"
)

func timestamp() string {
	//ms level timestamp
	return fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
}

func sign(e, t string) string {
	data := []byte("fanyideskweb" + e + t + "n%A-rKaT5fb[Gy?;N5@Tj")
	return fmt.Sprintf("%x", md5.Sum(data))
}
