package functions

import (
	"fmt"
	"sort"
)

type pair struct {
	key   interface{}
	value interface{}
}

type keys []interface{}

func (p *pair) String() string {
	return fmt.Sprintf("%v: %v", p.key, p.value)
}

func (k keys) String() string {
	res := "["
	notFirst := false
	for _, key := range k {
		if notFirst {
			res += ", "
		}
		res += fmt.Sprintf("%v", key)
		notFirst = true
	}
	res += "]"
	return res
}

type BaseMap struct {
	size        int
	orderedKeys []interface{}
	container   []pair
	m           map[interface{}]interface{}
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
	keys := make([]interface{}, 0, len(b.m))
	t := ""
	for k := range b.m {
		if t == "" {
			t = fmt.Sprintf("%T", k)
		}
		if t != fmt.Sprintf("%T", k) {
			panic(fmt.Sprintf("different key types in the map: %T and %s", k, t))
		}
		keys = append(keys, k)
	}
	b.orderedKeys = keys
	b.size = len(b.m)

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
	b.container = nil
	for _, k := range b.orderedKeys {
		b.container = append(b.container, pair{k, b.m[k]})
	}
}

func (b *BaseMap) Keys() keys {
	b.reorder()
	return b.orderedKeys
}

func (b *BaseMap) String() string {
	b.reorder()
	res := "{"
	notFirst := false
	for _, p := range b.container {
		if notFirst {
			res += "; "
		}
		res += fmt.Sprintf("%v: %v", p.key, p.value)
		notFirst = true
	}
	res += "}"
	return res
}

func (b *BaseMap) Size() int {
	b.reorder()
	return b.size
}
