package container

import (
	"fmt"
	"reflect"
	"testing"
)


func TestPut(t *testing.T) {
	typeOf := 0
	c := NewContainer(reflect.TypeOf(typeOf), 63)
	c.Put(12)

	if err := c.Get(12); err != nil {
		fmt.Println("get err: ", err)
	}
}

