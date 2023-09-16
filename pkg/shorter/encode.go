package shorter

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/url"
)

var link *url.URL


func Encode(s string) (string, error) {

	url, err := AddrValidation(s)
	if err != nil {
		log.Printf("%s is not a URL", s)
		return "", err
	} else {
		log.Printf("Encoding %s", s)
		encode := sha256.New()
		encode.Write([]byte(url.Path))
		encode.Sum(nil)
		return fmt.Sprintf("%X", string(encode.Sum(nil)[:5])), nil
	}
}

func Print(u *url.URL) {
	fmt.Println(u)
}