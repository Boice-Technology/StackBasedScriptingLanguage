// Check whether a stack is empty or not.

package stack

func (s * Stack) IsEmpty() bool {
	var verdict bool
	if(s.height == 0){
		verdict = true
	} else if(s.height > 0) {
		verdict = false
	}
	return verdict
}