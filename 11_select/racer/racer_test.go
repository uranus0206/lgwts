package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// func TestRacer(t *testing.T) {
// 	url1 := "http://www.facebook.com"
// 	url2 := "http://www.quii.co.uk"

// 	want := url2
// 	got := WebSiteRacer(url1, url2)

// 	if got != want {
// 		t.Errorf("got '%s', want '%s'", got, want)
// 	}
// }

func TestRacer(t *testing.T) {
	t.Run("return fastest", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := WebSiteRacer(slowUrl, fastUrl)
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {
			slowServer := makeDelayServer(11 * time.Second)
			fastServer := makeDelayServer(12 * time.Second)
			defer slowServer.Close()
			defer fastServer.Close()

			slowUrl := slowServer.URL
			fastUrl := fastServer.URL

			_, err := WebSiteRacer(slowUrl, fastUrl)
			if err == nil {
				t.Error("expected an error but didn't get one.")
			}
		})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))

	return server
}
