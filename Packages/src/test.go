package main

import ("fmt";
		)
		
func main(){
	var a interface{} = 1
	fmt.Printf("%T %v\n",a,a)
	fmt.Println(a.(int)+2)
}