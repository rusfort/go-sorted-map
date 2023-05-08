package functions

import (
	"fmt"
	"sort"
	"sync"
)

type pair struct {
	key interface{}
	value interface{}
}

type BaseMap struct {
	mtx sync.Mutex //TODO: use!
	size int
	orderedKeys []interface{}
	container []pair
	m map[interface{}]interface{}
}

func NewMap() *BaseMap {
	return &BaseMap{
		m: make(map[interface{}]interface{}),
	}
}

func (b *BaseMap) Map() map[interface{}]interface{} {
	return b.m
}

func (b *BaseMap) reorder() {
	if len(b.m) != b.size {
		//TODO: refresh orderedKeys and container!
	}
	if len(b.orderedKeys) == 0 {
		return
	}
	sort.Slice(b.orderedKeys, func(i, j int) bool {
		switch b.orderedKeys[0].(type) {
		case int:
			return b.orderedKeys[i].(int) < b.orderedKeys[j].(int)
		case int64:
			return b.orderedKeys[i].(int64) < b.orderedKeys[j].(int64)
		case float64:
			return b.orderedKeys[i].(float64) < b.orderedKeys[j].(float64)
		case string:
			return b.orderedKeys[i].(string) < b.orderedKeys[j].(string)
		default:
			panic(fmt.Sprintf("type %T is not allowed as a map key", b.orderedKeys[0]))
		}
	})
}

func (b *BaseMap) Keys() []interface{} {
	b.reorder()
	return b.orderedKeys
}

