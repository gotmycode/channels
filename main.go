package main

import (
	"fmt"
	"net/http"
)

//status checker - whether or not certain sites respond to traffic
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://thalamed.com",
	}

	//index not interesting for this exercise
	for _, link := range links {
		checkLink(link)
	}
}

//take a link and see if it responds to traffic
func checkLink(link string) {
	//consider if error comes back, not care about the response (which is here the _). If err value is not nil, that means site is up
	_, err := http.Get(link)
	if err != nil {
		// , or + in this case here no need for string concatenration, just separate the two into separate args and then print out correct string
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
