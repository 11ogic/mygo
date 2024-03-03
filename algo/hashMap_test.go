package algo

import "testing"

var (
	capacity = 100
)

type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
}

func NewArrayHashMap() *arrayHashMap {
	buckets := make([]*pair, capacity)
	return &arrayHashMap{buckets: buckets}
}

func (a *arrayHashMap) hashFunc(key int) int {
	return key % capacity
}

func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "Not found"
	}
	return pair.val
}

func (a *arrayHashMap) set(key int, val string) {
	index := a.hashFunc(key)
	pair := &pair{key, val}
	a.buckets[index] = pair
}

func (a *arrayHashMap) delete(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

func (a *arrayHashMap) each(cb func(val string, key int)) {
	for _, bucket := range a.buckets {
		if bucket != nil {
			cb(bucket.val, bucket.key)
		}
	}
}

func TestHashMap(t *testing.T) {
	hashMap := NewArrayHashMap()
	hashMap.set(888, "龙年大吉")
	hashMap.set(666, "身体健康")
	hashMap.each(func(val string, key int) {
		t.Logf("%v %v \n", val, key)
	})
}
