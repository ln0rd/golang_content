### Conteudo sobre Golang

Nesse Readme haverá exemplos e explicações sobre a linguagem e como cada coisa é feita.

1. **Packages**: <br>
   Cada arquivo é separado por um package logo, a primeira linha de cada arquivo deve conter `package name-package`, para cada projeto nos temos um go.mod onde está centralizado os packages externos do seu projeto. Nessa pasta `GO_STUDY_2` havera apenas 1 go.mod para todo os exemplos. <br>

   Idealmente o nome do modulo do seu projeto deve seguir casando com o repositorio por exemplo nesse projeto deveria ser: `github.com/ln0rd/golang_content/go_study_2` <br>

   1.1 **Build**: <br>
   Para compilar um projeto nós executamos `go build` que irá gerar um arquivo binário com o nome do projeto que está no `go.mod`, esse no caso é o que levamos para ser executado em algum ambiente real de produção. Para que isso ocorra o **main.go** deverá estar no mesmo nivel que o `go.mod`. <br>

   Para executar o binário basta executar `./nome-do-binário`. <br>

   1.2 **Imports**: <br>
   `Toda função com inicio letra maiuscula por exemplo "Write()" é a forma de dizer que a função é public, se estiver minuscula é private` <br><br>
   Quando importar da main, você coloca o nome do modulo/package, exemplo: `import "go_study_2/packages/helper"`; <br>
   Quando importar um arquivo no mesmo nivel, basta chamar a função, ou Struct, não precisa realizar import; <br>
   Quando importar um arquivo em outro package um nivel acima você sempre importa o caminho completo começando pelo nome do seu modulo <br>

   1.3 **External packages** <br>
   Para importar uma dependencia externa o comando é: `go get github.com/badoux/checkmail` <br>
   Para remover dependencias que não estão sendo usadas: `go mod tidy` <br>

   1.4 **Vendor**
   Vendor normalmente é quem disponibiliza alguma dependencia. Considere que hoje o repositório do Golang é o github, logo você corre o risco de alguem excluir algum repositorio e você perder sua dependencia, para isso você pode adotar uma estratégia como deixar as dependencias local no seu projeto executando: `go mod vendor`, isso irá criar uma pasta `vendor` com os códigos de suas dependencias.
