package main

import (
	"net/url"
	"fmt"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//s := "https://golang.org/pkg/net/url/#URL.Parse"
	u, err := url.Parse(s)
	//exception: defer, panic, recover
	//defer: defer some function to a stack
	//panic: different from throw exception, just exit current Routine after process all defered functions
	//recover: handle panic, used in defer function
	if err != nil {
		panic(err)
	}
	fmt.Printf("Scheme: %s\n", u.Scheme);
	fmt.Printf("Opaque: %s\n", u.Opaque);
	fmt.Printf("User: %s\n", u.User);
	p, _ := u.User.Password();
	fmt.Printf("User: %s\n", p);
	fmt.Printf("Host: %s\n", u.Host);
	fmt.Printf("Path: %s\n", u.Path);
	fmt.Printf("RawPath: %s\n", u.RawPath);
	fmt.Printf("ForceQuery: %s\n", u.ForceQuery);
	fmt.Printf("RawQuery: %s\n", u.RawQuery);
	fmt.Printf("Fragment: %s\n", u.Fragment);

}