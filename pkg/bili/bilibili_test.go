package bili

import (
	"strconv"
	"testing"
)

func TestAvToBv(t *testing.T) {
	tests := map[int64]string{
		92392262: "BV19E411n7fi",
		88496436: "BV1L7411V7x7",
		93640132: "BV1qE411s7qT",
	}
	for av, ev := range tests {
		bv, err := avToBv(av)
		if err != nil {
			t.Errorf("Transform error: %v", err)
			break
		}
		if ev != bv {
			t.Errorf("Transform failed, expect: %s, got: %s", ev, bv)
		}
	}
}

func TestBvToAv(t *testing.T) {
	tests := map[string]int64{
		"BV19E411n7fi": 92392262,
		"BV1L7411V7x7": 88496436,
		"BV1qE411s7qT": 93640132,
	}
	for bv, ev := range tests {
		av, err := bvToAv(bv)
		if err != nil {
			t.Errorf("Transform error: %v", err)
			break
		}
		if av != ev {
			t.Errorf("Transform failed, expect: %d, got: %d", ev, av)
		}
	}
}

func TestVideoStat(t *testing.T) {
	tests := []string{
		"BV1qE411s7qT",
		"92392262",
	}
	for _, vid := range tests {
		st, err := VideoStat(vid)
		if err != nil {
			t.Errorf("Get video stat error: %v", err)
			break
		}
		if st.BV != vid && strconv.Itoa(int(st.Aid)) != vid {
			t.Errorf("Get video stat failed for %s, %d - %s", vid, st.Aid, st.BV)
		}
	}
}
