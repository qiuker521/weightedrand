package weightedrand

//我写这个代码的时候我是知道这个代码是啥意思的
//但是现在你再问，我已经不知道了
import (
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Choice struct {
	Item   string
	Weight float64
}

type Chooser struct {
	mutex  sync.RWMutex
	data   []Choice
	totals []float64
	max    float64
}

func (chs *Chooser) NewChooser(cs ...Choice) {
	n := len(cs)
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Weight < cs[j].Weight
	})
	totals := make([]float64, n, n)
	var runningTotal float64 = 0
	for i, c := range cs {
		runningTotal += c.Weight
		totals[i] = runningTotal
	}
	chs.mutex.Lock()
	defer chs.mutex.Unlock()

	chs.data = cs
	chs.totals = totals
	chs.max = runningTotal
}

/*
Pick pick by totally random.
*/
func (chs *Chooser) Pick() string {
	return chs.pickFloat64(rand.Float64() * chs.max)
}

/*
PickByHash pick by totally consistency hash.
*/
func (chs *Chooser) PickByHash(hash float64) string {
	return chs.pickFloat64(hash)
}

/*
pick one
*/
func (chs *Chooser) pickFloat64(r float64) string {
	chs.mutex.RLock()
	defer chs.mutex.RUnlock()
	//预防可能存在的误用
	if r >= chs.max {
		r = math.Mod(r, chs.max)
	}
	i := sort.SearchFloat64s(chs.totals, r)
	return chs.data[i].Item
}
