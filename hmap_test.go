package hmap

import (
	"reflect"
	"testing"
)

func TestPutGetDelete(t *testing.T) {
	tKey := []int{1, 2, 3}
	tVal := "whatever"

	c := New()
	c.Put(tKey, tVal)

	val, err := c.Get(tKey)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(val, tVal) {
		t.Fatalf("Expected %v but got %v. \n", tVal, val)
	}

	c.Delete(tKey)
	val, err = c.Get(tKey)
	if err == nil {
		t.Fatalf("Was expecting nothing but got %v", val)
	}
}

func TestCloneClear(t *testing.T) {
	c := New()
	c.Put(12, 123)
	c.Put([]int{1, 2, 6776}, "test")
	nc := c.Clone()

	if l := c.Length(); l != nc.Length() && l != 2 {
		t.Fatal("Cloning a map failed")
	}

	val1, err1 := nc.Get(12)
	val2, err2 := nc.Get([]int{1, 2, 6776})

	if err1 != nil || err2 != nil {
		t.Fatal("Cloning failed.")
	}

	if !reflect.DeepEqual(val1, 123) || !reflect.DeepEqual(val2, "test") {
		t.Fatalf("Cloning failed. Expected %v and %v but got %v and %v \n", 123, "test", val1, val2)
	}

	c.Clear()
	if l := c.Length(); l != 0 {
		t.Fatalf("Expected empty map but length is %v", l)
	}

	val1, err1 = c.Get(12)
	val2, err2 = c.Get([]int{1, 2, 6776})
	if err1 == nil {
		t.Fatalf("Expected error NOTFOUND but got %v", val1)
	}

	if err2 == nil {
		t.Fatalf("Expected error NOTFOUND but got %v", val2)
	}

}
