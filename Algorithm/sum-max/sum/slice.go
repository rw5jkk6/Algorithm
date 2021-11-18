package main
 
import (
   "fmt"
)

const n1 = 10000000

func sum(nl []int) int{
   sum := 0
   for _, v := range nl{
      sum += v
   }
   return sum
}

func main(){
   nl := make([]int, n1)
   for i:= 0; i < n1; i++{
      nl = append(nl,i)
   }

   num := sum(nl)
   fmt.Println(num)
}
