
<h1 align="center">
Golang + Gin + Testify + SQLC + Postgres + Docker
</h1>

<p> Gustavo Noronha Tutorial </p>
<p> Playlist: https://www.youtube.com/watch?v=EYgnlMWhrnM&list=PLcE-9cucnhqW7g8Uw6j1-QAgSbPpeZ6p8</p>

## Ferramentas e configurações
- A solução foi desenvolvida através do Sistema Operacional Ubuntu 20.04.5 LTS; 
-  A IDE utilizada foi o VS Code, versão 1.74.2, sendo adicionadas extensões para o Docker e o Go; 
-  Foi utilizado o Docker versão 20.10.22 e o Go versão 1.19.4.
- Para utilizar a imagem do Postgres no Docker (https://hub.docker.com/_/postgres/):
- ```docker pull postgres:15.1-alpine```
- Para ativar a imagem do Postgres no Docker:
- ```docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres```
- Para executar a imagem do Postgres no Docker:
- ```docker exec -it postgres psql -U postgres```
- Para atualizar as dependências do projeto:
```go mod tidy```
- Para inicializar o server:
- ```go run main.go```

## pgAdmin
- Para visualizar as informações gravadas na imagem do Postgres, utilizou-se o pgAdmin versão 6.18:
- https://www.pgadmin.org/download/pgadmin-4-apt/

- Configurações da criação de um server no pgAdmin:
- nome: go-products
- conexão / host: localhost
- porta: 5432
- database principal (Maintenance database): postgres
- username: postgres
- password: postgres (a mesma definida ao ativar a imagem do Postgres no Docker)
- SSL mode: disable
- Criar o Database 'go_products' no pgAdmin, dentro do server 'go-products'

## Migrate
- Foram utilizadas migrations para criação de tabelas no banco de dados
- Para instalação, seguir os passos disponíveis no link:
- https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Criação da tabela create_table_product:
- ```migrate create -ext sql -dir db/migration -seq create_table_product```
- Após a definição do SQL nos arquivos UP e DOWN e a criação do Database no Postgres, rodar as migrations através do seguinte comando:
- ```migrate -path db/migration -database “postgresql://postgres:postgres@localhost:5432/go_products?sslmode=disable” -verbose up ```

## SQLC
- Para a definição das regras de negócio do projeto, utilizou-se o SQLC (https://github.com/kyleconroy/sqlc)
- Para a instalação (https://docs.sqlc.dev/en/latest/overview/install.html):
- ```docker pull kjconroy/sqlc```
- Execução das configurações definidas no arquivo 'products.sql':
- ```docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate```

## Testify
- Criação de testes unitários
- Instalação do Postgres no projeto (https://github.com/lib/pq):
- ```go get github.com/lib/pq```
- Instalação do Testify (https://github.com/stretchr/testify):
- ```go get github.com/stretchr/testify```

## Gin
- Para a criação de rotas HTTP, utilizou-se o Gin (https://github.com/gin-gonic/gin)
- Instalação:
- ```go get -u github.com/gin-gonic/gin```
