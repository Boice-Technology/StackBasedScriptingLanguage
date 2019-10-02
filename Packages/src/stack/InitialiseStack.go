// Initialise Stack variable.

package stack

func (s * Stack) InitialiseStack(inputItems ...Item) {
	s.height = 0
	s.datastore = make([]Item, 0)
	s.datastore = append(s.datastore, inputItems...)
	s.height = int64(len(s.datastore))
}