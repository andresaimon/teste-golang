/*Passo 1: definir o SQL da criação das tabelas */
CREATE TABLE "products" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" varchar NOT NULL,
    "price" integer NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

/*
Passo 3: executar a migration através do comando:
migrate -path db/migration -database “postgresql://postgres:postgres@localhost:5432/go_products?sslmode=disable” -verbose up
                                                    usuário   senha        
 
Passo 4: instalar o SQLC e criar o arquivo sqlc.yaml na raiz no projeto, para configurá-lo

Passo 5: criação da pasta query e do arquivo product.sql, dentro da pasta db, para codificar as querys
 
Passo 6: executar o SQLC a partir do comando: docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
Então, cria-se automaticamente a pasta sqlc, com os arquivos:
db.go
models.go
product.sql.go
querier.go

Passo 7: rodar o comando: go mod init teste
Então, cria-se automaticamente o arquivo go.mod
 */