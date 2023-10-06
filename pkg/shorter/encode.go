package shorter

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/url"
)

func Encode(u *url.URL) (string, error) {

	log.Printf("Encoding %s", u)
	encode := sha256.New()
	encode.Write([]byte(fmt.Sprint(u.String())))
	encode.Sum(nil)
	return fmt.Sprintf("%X", string(encode.Sum(nil)[:5])), nil

}


