package hashset

//Hashset is a set implementation using builtin map
type Hashset struct {
	data map[interface{}]struct{}
}

//New returns a reference to a new Hashset
func New() *Hashset {
	var hs Hashset
	hs.data = make(map[interface{}]struct{})
	return &hs
}

//Empty test if hs is empty
func (hs *Hashset) Empty() bool {
	return len(hs.data) == 0
}

//Size returns the size of hs
func (hs *Hashset) Size() int {
	return len(hs.data)
}

//Insert inserts v in hs
func (hs *Hashset) Insert(v interface{}) {
	hs.data[v] = struct{}{}
}

//Delete deletes v from hs if hs contains v
func (hs *Hashset) Delete(v interface{}) {
	delete(hs.data, v)
}

//Content returns a slice containing all the elements in hs
func (hs *Hashset) Content() []interface{} {
	c := make([]interface{}, len(hs.data))
	i := 0
	for k := range hs.data {
		c[i] = k
		i++
	}
	return c
}

//Clear removes all the content in hs
func (hs *Hashset) Clear() {
	hs.data = make(map[interface{}]struct{})
}

//Has test if hs contains v
func (hs *Hashset) Has(v interface{}) bool {
	_, c := hs.data[v]
	return c
}
