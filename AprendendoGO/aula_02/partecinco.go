//PARTE 5 DA AULA
package main

import (
	"fmt"
)

type number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}


//Função criada para achar o maior número 

func findMax[T number](Arr []T) T {
	max := Arr[0] 
	for _, value := range Arr {
	  if value > max {
		 max = value
		
		}
	}
	 return max
}

func main (){
  maxInt := findMax([] int8{1,2,9,56,0,3,7})
  maxFloat := findMax([] float32{0.8, 5.2, 5.3, 1.0, 7.9})

  fmt.Println(maxInt)
  fmt.Printf("%.2f", maxFloat)

  //"%.2f" serve para formatar as casas após a vírgula
}