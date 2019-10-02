// Define Stack data structure

package stack 

type Item interface{}

type Stack struct{
	datastore []Item
	height int64
}