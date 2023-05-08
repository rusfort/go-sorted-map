package functions

import (
	//"context"
	//"fmt"
)

type pair struct {
	key interface{}
	value interface{}
}

type BaseMap struct {
	orderedKeys []interface{}
	container []pair
}

func NewMap() *BaseMap {
	return &BaseMap{}
}

