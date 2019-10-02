// Initialise Stack variable.

package stack

func (s * Stack) InitialiseStack(inputItems ...item) {
	s.height = 0
	s.datastore = make([]item, 0)
	s.datastore = append(s.datastore, inputItems...)
	s.height = int64(len(s.datastore))
}