### Conteudo sobre Golang

Nesse Readme haverá exemplos e explicações sobre a linguagem e como cada coisa é feita.

1. **Packages**: <br>
   Cada arquivo é separado por um package logo, a primeira linha de cada arquivo deve conter `package name-package`, para cada projeto nos temos um go.mod onde está centralizado os packages externos do seu projeto. Nessa pasta `GO_STUDY_2` havera apenas 1 go.mod para todo os exemplos. <br>

   1.1 **Build**: <br>
   Para compilar um projeto nós executamos `go build` que irá gerar um arquivo binário com o nome do projeto que está no `go.mod`, esse no caso é o que levamos para ser executado em algum ambiente real de produção. Para que isso ocorra o **main.go** deverá estar no mesmo nivel que o `go.mod`. <br>

   Para executar o binário basta executar `./nome-do-binário`. <br>

   1.2 **Imports**: <br>
   `Toda função com inicio letra maiuscula por exemplo "Write()" é a forma de dizer que a função é public, se estiver minuscula é private` <br><br>
   Quando importar da main, você coloca o nome do modulo/package, exemplo: `import "go_study_2/packages/helper"`; <br>
   Quando importar um arquivo no mesmo nivel, basta chamar a função, ou Struct, não precisa realizar import; <br>
   Quando importar um arquivo em outro package um nivel acima ...?

   1.3 **External packages**
   Para importar uma dependencia externa o comando é: `go get github.com/badoux/checkmail`
   Para remover dependencias que não estão sendo usadas: `go mod tidy`
