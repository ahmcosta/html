package html

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// Title expects a URLs array. For each URL, it gets the page, cuts its body, search its title using regex
func Title(urls ...string) <-chan string { // <-chan - canal somente-leitura
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { // anonymous function
			beginTime := time.Now()
			resp, _ := http.Get(url)             // get the page by url
			html, _ := ioutil.ReadAll(resp.Body) // get the body of the page gotten

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // compile the regex
			title := r.FindStringSubmatch(string(html))[1]
			finishTime := time.Now()
			totalTime := finishTime.Sub(beginTime)
			c <- fmt.Sprintf("Elapsed Time: %d\tTitle: %s", totalTime, title) // return the string from compiled regex
		}(url) // It's needed because it's a anonymous functions. So it must be called
	}
	return c
}
