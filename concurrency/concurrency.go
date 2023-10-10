package main

import "fmt"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for i, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
			fmt.Printf("Result Channel %d\n", i)
			fmt.Println("resultChannel", resultChannel)
			// results[u] = wc(u)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		fmt.Println("r %d", i)
		fmt.Println("rrrr", r)
		fmt.Println("rstring", r.string)
		fmt.Println("rbool", r.bool)
		results[r.string] = r.bool
	}

	fmt.Println("results")
	fmt.Println(results)

	return results
}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func main() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	CheckWebsites(mockWebsiteChecker, websites)
}
