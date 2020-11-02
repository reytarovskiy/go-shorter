package shorter

import "math/rand"

type Shorter interface {
	Short(url string) string
}

type RandomShorter struct {
	len int
}

func NewRandomShorter(len int) *RandomShorter {
	return &RandomShorter{len: len}
}

func (r *RandomShorter) Short(url string) string {
	short := ""
	for i := 0; i < r.len; i++ {
		short += getRandomSymbol()
	}

	return short
}

var alphabet = []string{
	"q", "w", "r", "t", "y", "u", "i", "o",
	"p", "a", "s", "d", "f", "g", "h", "j",
	"k", "l", "z", "x", "c", "v", "b", "n",
	"m", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func getRandomSymbol() string {
	min := 0
	max := len(alphabet)
	n := rand.Intn(max - min) + min
	return alphabet[n]
}
