package shorter

import (
	"crypto/sha256"
	"fmt"
)

type link struct {
	addr string
}

func NewLink(addr string) link {
	url := link{
		addr: addr,
	}
	return url
}

func (l *link) Encode() string {
	encode := sha256.New()
	encode.Write([]byte(l.addr))
	encode.Sum(nil)
	return fmt.Sprintf("%X", string(encode.Sum(nil)[:5]))
}

func (l *link) print() {
	fmt.Println(l)
}