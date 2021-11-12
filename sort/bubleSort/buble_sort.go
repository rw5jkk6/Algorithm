package main

import "fmt"

func main(){
	data := []int{3, 1, 2,4}
	len_data := len(data)
	for i:=0; i<len_data-1; i++{
		for j:=0; j<len_data-i-1; j++{
			if data[j]>data[j+1]{
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println(data)
}
