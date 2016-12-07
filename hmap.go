// Package hmap implements a wrapper around the map[interface{}]interface{} type.
// Contrary to the wrappee, keys are unconstrained in terms of types that may be
// used.
// slice, map, and func based types, if used for keys, will trigger deep-value
// comparison of map keys.
// However, we lose the O(1) behaviour of standard maps for slice, maps and func
// keys.
package hmap

import (
	"errors"
	"reflect"

	"github.com/atdiar/sac"
)

// Container is the type defining the implementation of the hash map.
// Not safe for concurrent use by multiple goroutines.
type Container struct {
	hashmap map[interface{}]interface{}
	ull     *sac.Container
}

// New returns a new empty map
func New() Container {
	return Container{
		hashmap: make(map[interface{}]interface{}),
		ull:     sac.New(),
	}
}

// Get retrieves an element of the map if it exists. Otherwise, a non-nil
// sentinel error is returned.
func (c Container) Get(key interface{}) (val interface{}, err error) {
	if !reflect.TypeOf(key).Comparable() {
		return c.ull.Get(key)
	}
	val, ok := c.hashmap[key]
	if !ok {
		err = errors.New("NOTFOUND")
	}
	return val, err
}

// Put inserts a value into a map for the provided key.
func (c Container) Put(key, val interface{}) {
	if !reflect.TypeOf(key).Comparable() {
		c.ull.Put(key, val)
		return
	}
	c.hashmap[key] = val
}

// Delete erases a value from a map for a given key.
func (c Container) Delete(key interface{}) {
	if !reflect.TypeOf(key).Comparable() {
		c.ull.Delete(key)
		return
	}
	delete(c.hashmap, key)
}

// Length returns the number of elements stored into the map.
func (c Container) Length() int {
	l := 0
	for range c.hashmap {
		l++
	}
	return l + c.ull.Length()
}

// Clear will delete all elements of a map
func (c Container) Clear() {
	for k := range c.hashmap {
		delete(c.hashmap, k)
	}
	c.ull.Clear()
}

// Clone will create and return a copy of a map.
func (c Container) Clone() Container {
	j := New()
	j.ull = c.ull.Clone()
	for k, v := range c.hashmap {
		j.hashmap[k] = v
	}
	return j
}
