package racer

import (
	"fmt"
	"net/http"
	"time"
)

func WebSiteRacer(url1, url2 string) (winner string, err error) {
	// aDuration := measureResponseTime(url1)
	// bDuration := measureResponseTime(url2)

	// if aDuration < bDuration {
	// 	return url1
	// }
	// return url2

	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("time out waiting for %s and %s", url1, url2)
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) <-chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
