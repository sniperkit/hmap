#hmap

[![GoDoc](https://godoc.org/github.com/atdiar/hmap?status.svg)](https://godoc.org/github.com/atdiar/hmap)

hmap - An alternative map implementation
-------------------------------------------------------------

This package defines a datastructure that may be used in lieu of
a map[interface{}]interface{}.

As a key/value store, it deals with **6** operations:

* `Get()` tries to retrieve a Value. Returns an error if key is absent.  

* `Put()` inserts a Value for a given Key.  

* `Delete() `  

* `Clear()` deletes all elements.  

* `Length()`returns the number of elements stored.

* `Clone()`  will return a deep-copy of a hmap.

It is not safe for concurrent use. Reason being that we do not want to tie the
content of the map to the scheduling of goroutines. It would render the content
undeterministic between runs of a same program.

For performance comparisons, have a look at the test file which includes synthetic benchmarks.

#Why not simply use a standard map?
A standard Go map can be defined as accepting an interface key.
However, in Go1, types that are based on slices, funcs, or maps are not
comparable. Henceforth, they may not be used as map keys (it panics).

This is an implicit restriction that we hope to see lifted in Go2.
In the meanwhile, the hmap package lifts he  implicit restriction on map keys.


For completeness, please refer to the package [documentation].

[documentation]:https://godoc.org/github.com/atdiar/hmap
