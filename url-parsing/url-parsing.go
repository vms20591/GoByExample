package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:password@localhost.com:5332/path?k=v#f"

	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	// Get credential information from UserInfo
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	fmt.Println(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)

	fmt.Println(u.Fragment)

	// get raw query in k=v format
	fmt.Println(u.RawQuery)

	// parse raw query into a map
	// map[string][]string
	qp, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(qp)
	fmt.Println(qp["k"][0])
}
