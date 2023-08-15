package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		coutner := Counter{}
		coutner.Inc()
		coutner.Inc()
		coutner.Inc()

		if coutner.Value() != 3 {
			t.Errorf("got %d, want %d", coutner.Value(), 3)
		}
	})
}
