
//PARTE 1 DA AULA
package main

type animal interface {
	sound() string
	move() string
}

type cao struct {
	nome string
}

func (d cao) sound() string {
	return d.nome + " Say, roofing\n"
}

func (d cao) move() string {
	return d.nome + " Is running\n"
}

func makeMove(a animal) {
	print(a.move())
	print(a.sound())
}


func main() {
  cachorro := cao{
    nome: " Pituco",
  }

  makeMove(cachorro)
}