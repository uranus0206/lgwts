package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultCh := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultCh <- result{url, wc(url)}
		}(url)
	}

	// for v := range resultCh {
	// 	results[v.string] = v.bool
	// }
	for i := 0; i < len(urls); i++ {
		res := <-resultCh
		results[res.string] = res.bool
	}

	return results
}
