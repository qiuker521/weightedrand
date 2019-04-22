package weightedrand

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

var minRatio = 0.98
var maxRatio = 1.02

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func v15() (float64, float64, float64, float64, float64) {
	return float64(rand.Intn(1000000)) * rand.Float64(),
		float64(rand.Intn(1000000)) * rand.Float64(),
		float64(rand.Intn(1000000)) * rand.Float64(),
		float64(rand.Intn(1000000)) * rand.Float64(),
		float64(rand.Intn(1000000)) * rand.Float64()
}

func newChooser(v1, v2, v3, v4, v5 float64) *Chooser {
	log.Println(v1, v2, v3, v4, v5)
	var c = Chooser{}
	c.NewChooser(
		Choice{Item: "1", Weight: v1},
		Choice{Item: "2", Weight: v2},
		Choice{Item: "3", Weight: v3},
		Choice{Item: "4", Weight: v4},
		Choice{Item: "5", Weight: v5},
	)
	return &c
}

func i15() (float64, float64, float64, float64, float64) {
	return 0.0, 0.0, 0.0, 0.0, 0.0
}

func c15(count, total, i1, i2, i3, i4, i5, v1, v2, v3, v4, v5 float64) (float64, float64, float64, float64, float64) {
	c1 := (i1 / count) / (v1 / total)
	c2 := (i2 / count) / (v2 / total)
	c3 := (i3 / count) / (v3 / total)
	c4 := (i4 / count) / (v4 / total)
	c5 := (i5 / count) / (v5 / total)
	return c1, c2, c3, c4, c5
}

func calcBs(i1, i2, i3, i4, i5, i6 *float64, bs []string) (float64, float64, float64, float64, float64, float64) {
	for _, v := range bs {
		switch v {
		case "1":
			*i1++
		case "2":
			*i2++
		case "3":
			*i3++
		case "4":
			*i4++
		case "5":
			*i5++
		default:
			*i6++
		}
	}
	return *i1, *i2, *i3, *i4, *i5, *i6
}

func TestRandChoice(t *testing.T) {
	var v1, v2, v3, v4, v5 = v15()
	var i1, i2, i3, i4, i5 = i15()
	var i6 = 0.0
	var c = newChooser(v1, v2, v3, v4, v5)
	total := v1 + v2 + v3 + v4 + v5
	var bs = []string{}

	var count float64 = 10000000

	for i := 1.0; i <= count; i++ {
		x := c.Pick()
		bs = append(bs, x)
	}
	i1, i2, i3, i4, i5, i6 = calcBs(&i1, &i2, &i3, &i4, &i5, &i6, bs)

	var c1, c2, c3, c4, c5 = c15(count, total, i1, i2, i3, i4, i5, v1, v2, v3, v4, v5)

	if minRatio > c1 || maxRatio < c1 {
		t.Error(c1, "expected error")
	}
	if minRatio > c2 || maxRatio < c2 {
		t.Error(c2, "expected error")
	}
	if minRatio > c3 || maxRatio < c3 {
		t.Error(c3, "expected error")
	}
	if minRatio > c4 || maxRatio < c4 {
		t.Error(c4, "expected error")
	}
	if minRatio > c5 || maxRatio < c5 {
		t.Error(c5, "expected error")
	}
	if i6 != 0 {
		t.Error("out of case expected 0, get ", i6)
	}
	t.Log(c1, c2, c3, c4, c5)
}

func TestHashChoice(t *testing.T) {
	var v1, v2, v3, v4, v5 = v15()
	var i1, i2, i3, i4, i5 = i15()
	var i6 = 0.0
	var c = newChooser(v1, v2, v3, v4, v5)
	total := v1 + v2 + v3 + v4 + v5
	var bs = []string{}

	var count float64 = 10000000

	for i := 1.0; i <= count; i++ {
		x := c.PickByHash(rand.Float64() * count * count)
		bs = append(bs, x)
	}
	i1, i2, i3, i4, i5, i6 = calcBs(&i1, &i2, &i3, &i4, &i5, &i6, bs)

	var c1, c2, c3, c4, c5 = c15(count, total, i1, i2, i3, i4, i5, v1, v2, v3, v4, v5)

	if minRatio > c1 || maxRatio < c1 {
		t.Error(c1, "expected error")
	}
	if minRatio > c2 || maxRatio < c2 {
		t.Error(c2, "expected error")
	}
	if minRatio > c3 || maxRatio < c3 {
		t.Error(c3, "expected error")
	}
	if minRatio > c4 || maxRatio < c4 {
		t.Error(c4, "expected error")
	}
	if minRatio > c5 || maxRatio < c5 {
		t.Error(c5, "expected error")
	}
	if i6 != 0 {
		t.Error("out of case expected 0, get ", i6)
	}
	t.Log(c1, c2, c3, c4, c5)
}

func BenchmarkRandChooser(b *testing.B) {
	var v1 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v2 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v3 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v4 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v5 float64 = float64(rand.Intn(1000000)) * rand.Float64()

	c := Chooser{}
	c.NewChooser(
		Choice{Item: "1", Weight: v1},
		Choice{Item: "2", Weight: v2},
		Choice{Item: "3", Weight: v3},
		Choice{Item: "4", Weight: v4},
		Choice{Item: "5", Weight: v5},
	)

	for i := 0; i < b.N; i++ {
		c.Pick()
	}
}

func BenchmarkHashChooser(b *testing.B) {
	var v1 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v2 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v3 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v4 float64 = float64(rand.Intn(1000000)) * rand.Float64()
	var v5 float64 = float64(rand.Intn(1000000)) * rand.Float64()

	c := Chooser{}
	c.NewChooser(
		Choice{Item: "1", Weight: v1},
		Choice{Item: "2", Weight: v2},
		Choice{Item: "3", Weight: v3},
		Choice{Item: "4", Weight: v4},
		Choice{Item: "5", Weight: v5},
	)

	for i := 0; i < b.N; i++ {
		c.Pick()
	}
}
