// Pop a item from stack.

package stack

func (s *Stack) Pop() (Item, bool)  {
	var popped Item
	var verdict bool
	if(s.IsEmpty()){
		popped, verdict =  nil, false
	} else {
		popped, verdict = s.datastore[s.height-1], true
		s.datastore = s.datastore[:s.height-1]
		s.height -= 1
	}
	return popped, verdict
}