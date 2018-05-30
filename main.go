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

	//create a channel -- can only be accessed within func main
	//need to pass on to checkLink
	c := make(chan string)

	//index not interesting for this exercise
	for _, link := range links {
		//checkLink(link) //without adding new go routine

		//with new Go Routine
		//runs the blocking call http.Get(link), do nothing else, then go thru next element in iteration
		//and launch/spawns another new Go Routine
		//think of "go" as the engine that starts to chump thru code inside of a single func
		go checkLink(link, c) //channel passed as argument
	}
	//receive a value from a channel
	//fmt.Println(<-c)
	//this routine is waiting for something to happen, so it is put into a pause
	//once message is received (blocking call), wakes up then it prints it out.
	//No more code to run, then exits

	//adding a 2nd Println - to show 2 log statements
	//repeats same process as first, main wakes up and goes through next blocking call
	//if no of Println > than URLs, it will just hang waiting for code to run
	//fmt.Println(<-c)

	//c inspired for loop
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

//take a link and see if it responds to traffic
//channel passed on as 2nd arg
func checkLink(link string, c chan string) {
	//consider if error comes back, not care about the response (which is here the _). If err value is not nil, that means site is up
	//http.Get(link) - blocking call . while this func is being executed the main go routine can do nothing else
	_, err := http.Get(link)
	if err != nil {
		// , or + in this case here no need for string concatenration, just separate the two into separate args and then print out correct string
		fmt.Println(link, "might be down!")

		//send a message into the channel
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")

	//send a message into the channel
	c <- "Yes, it is up"

}
