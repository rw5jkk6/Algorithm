package main
 
import (
   "fmt"
)

const n1 = 100000000

func max(s []int) int{
   num := 0
   for _, v := range s{
      if v > num{
         num = v
      }
   }
   return num
}

func main(){
   nl := make([]int, n1)
   for i:= 0; i < n1; i++{
      nl = append(nl,i)
   }

   num := max(nl)
   fmt.Println(num)
}
