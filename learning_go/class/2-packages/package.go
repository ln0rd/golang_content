package main

// Quando temos mais de um pacote precisamos usar os modulos
// Eles são equivalentes ao package.json em nodejs, onde irá centralizar
// As dependencias do projeto, para criar o comando ẽ:
// go mod init nome_do_pacote

// Para fazer o build do projeto go build

// para realizar o import e o nome do pacote/funcao do pacote
import (
	"fmt"
	"packages/helper"
)

func main() {
	helper.Write()
	fmt.Println("Working in golang packages")
	helper.Validating()
}
