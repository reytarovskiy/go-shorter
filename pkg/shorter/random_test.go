package shorter

import "testing"

func TestFixedLen(t *testing.T) {
	shorter := NewRandomShorter(10)

	res := shorter.Short("https://google.com")

	if len(res) != 10 {
		t.Errorf("invalid len, %s: %d", res, len(res))
	}
}
