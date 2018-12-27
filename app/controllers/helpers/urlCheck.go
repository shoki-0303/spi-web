package helpers

import (
	"fmt"
	"log"
	"net/url"
	"path"
)

func CheckURL(endpoint, sub, query string) string {
	fmt.Printf("\n")
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Printf("action=CheckURL err=%s", err)
	}
	u.Path = path.Join(u.Path, sub)
	q := u.Query()
	q.Set("token", query)
	u.RawQuery = q.Encode()
	url := u.String()
	return url
}
