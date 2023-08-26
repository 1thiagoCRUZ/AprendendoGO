
//PARTE 2 DA AULA
package main 

func main (){
  x:= 10
  b := &x


  //printando os endereços de memória das variaveis & | e os valores também 
  println(&x)
  println("X: ", x)
  
  println(b)
  println("B: ", *b)


  *b = 55
  
  println(&x)
  println("X: ", x)
  
  println(b)
  println("B: ", *b)
  
}