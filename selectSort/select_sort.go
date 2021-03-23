package main

import "fmt"

func main(){
	data := []int{2, 1, 4, 3}
	len_data := len(data)
    for i := 0; i < len_data; i++{
		tmp := i
		for j := 0; j < len_data-1 ; j++{
			if data[j] > data[tmp]{
				tmp = j
			}
		data[i], data[tmp] = data[tmp], data[i]
		}
	} 
	fmt.Println(data)
}
