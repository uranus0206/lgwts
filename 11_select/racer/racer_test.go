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
	slowServer := makeDelayServer(20 * time.Millisecond)
	fastServer := makeDelayServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	t.Run("return fastest", func(t *testing.T) {
		want := fastUrl
		got, _ := WebSiteRacer(slowUrl, fastUrl)
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s",
		func(t *testing.T) {
			// _, err := WebSiteRacer(slowUrl, fastUrl)
			_, err := ConfigurableRacer(slowUrl, fastUrl, 1)
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
