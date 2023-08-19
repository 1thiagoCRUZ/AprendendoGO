//PARTE 3 DA AULA
package main 

type Animal interface{
  sound() string
	move() string
}

type cao struct {
	nome string
}


//o * indica que a variavel sera um ponteiro

func (d *cao) sound() string {
  d.nome = "Rex"
	return d.nome + " Say, roofing\n"
}

func (d *cao) move() string {
	return d.nome + " Is running\n"
}

func makeMove(a Animal) {
	print(a.move())
	print(a.sound())
}

func main (){
  cachorro := cao{
    nome: " Pituco",
  }

  makeMove(&cachorro)
  println(cachorro.nome)
}
