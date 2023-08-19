//PARTE 4 DA AULA
package main

func PrintVal(a, b interface{}){

  if str, ok := b.(string); ok{
    println(len(str))
    println(str)
  } else {
    print("B não é uma string\n")
  }

  if value, ok := a.(int); ok{
    println(value+1)
    println(value)
  } else {
    print("A não é um inteiro")
  }
  
}

func main (){
  PrintVal(15, 4)
}

