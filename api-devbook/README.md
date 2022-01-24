### Devbook

Essa aplicação simula uma rede social, e tem ações como:

- CRUD
- Seguir outro usuário
- Para de seguir outro usuário
- Buscar todos os usuários que segue
- Buscar os usuáios que são seguidos
- Atualizar senha

Contendo duas tabelas

- Usuarios
- Seguidores

Esse projeto nasce na realização de um curso sobre Golang, aprendendo toda a sintaxe e a forma de construir um projeto.

#### Estruta da aplicação

-- Main `Inicio do programa` <br>
-- Router `Rotas` <br>
-- Controllers `Controllers` <br>
-- Modelos `Próximo ao que conhecemos com entidades` <br>
-- Repositórios `Manipulação do banco de dados` <br>
-- Config `Variaveis de ambiente, credenciais`<br>
-- Banco `Somente para abrir banco de dados` <br>
-- Autenticação `Login, Tokens` <br>
-- Middleware `Entre a requisção e reposta, para ver se o usuário tá autenticado` <br>
-- Segurança `Cuidar das senhas` <br>
-- Respostas `Padronização de repsostas da API` <br>

#### Commands to create this project

initialize `go.mod`

```
go mod init api-devbook
```

install dependency

```
go get github.com/gorilla/mux
```
