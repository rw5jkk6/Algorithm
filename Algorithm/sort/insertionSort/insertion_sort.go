package main

import "fmt"

func main(){
	data:=[]int{3, 4, 2, 1}
	len_data:=len(data)
	for i:=1; i < len_data; i++{
		for j :=0; j<i; j++{
			if data[i - j -1]> data[i-j]{
				data[i - j -1], data[i -j] = data[i-j], data[i-j -1]
			}else{
				break
			}
		} 
	}
	fmt.Println(data)
}