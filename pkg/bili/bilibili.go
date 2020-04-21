package bili

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

const (
	xor = 177451812
	add = 8728348608
)

var once sync.Once

var (
	table = []byte("fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF")
	s     = []int{11, 10, 3, 8, 4, 6}

	tr map[byte]int64
)

// transform BV to AV
func bvToAv(bv string) (av int64, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Transform BV: %s panic: %v", bv, x)
		}
	}()
	once.Do(func() {
		tr = make(map[byte]int64)
		n := len(table)
		for i := 0; i < n; i++ {
			tr[table[i]] = int64(i)
		}
	})
	var r int64
	for i := 0; i < 6; i++ {
		r += tr[bv[s[i]]] * int64(math.Pow(float64(58), float64(i)))
	}
	av = (r - add) ^ xor
	return
}

func avToBv(av int64) (bv string, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Transform AV: %s panic: %v", bv, x)
		}
	}()
	av = av ^ xor + add
	r := []byte("BV1  4 1 7  ")
	for i := 0; i < 6; i++ {
		p := math.Pow(float64(58), float64(i))
		v := math.Floor(float64(av) / p)
		idx := int(v) % 58
		r[s[i]] = table[idx]
	}
	bv = string(r)
	return
}

// BvToAv transform BV to AV
func BvToAv(bv string) (string, error) {
	av, err := bvToAv(bv)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(av, 10), nil
}
