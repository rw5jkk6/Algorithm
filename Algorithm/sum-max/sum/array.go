package main
 
import (
   "fmt"
)

const n1 = 10000000

func sum(nl [n1]int) int{
   sum := 0
   for _, v := range nl{
      sum += v
   }
   return sum
}

func main(){
   nl := [n1]int{}
   for i:= 0; i < n1; i++{
      nl[i]=i
   }

   num := sum(nl)
   fmt.Println(num)
}
