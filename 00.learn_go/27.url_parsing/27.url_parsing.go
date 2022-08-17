package main

import (
	"fmt"
	"net/url"
)

func main() {

	s := "http://mypage.com/photos?animal=cat&color=black#f"
	u, err := url.Parse(s)
	handleErr(err)
	fmt.Println(u.Scheme) // http

	fmt.Println(u.Host)     // mypage.com
	fmt.Println(u.Path)     //  /photos
	fmt.Println(u.Fragment) //  f

	args, err := url.ParseQuery(u.RawQuery)
	handleErr(err)
	fmt.Println(args) //  map[animal:[cat] color:[black]]
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
