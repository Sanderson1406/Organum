## Documentação

Para desenvolvedores, consulte nossa API Reference para obter detalhes sobre endpoints, autenticação e parâmetros.

https://documenter.getpostman.com/view/35096123/2sA3QqgYiR

# Documentação da API Organum

Bem-vindo à documentação da API Organum, uma plataforma multifuncional para organizar leituras de maneira eficiente. A API Organum é composta por um módulo principal:

## Lectio - Leitura Interativa

O módulo Lectio é dedicado aos amantes de livros, proporcionando uma maneira interativa de gerenciar suas leituras. Com Lectio, os usuários podem:

- Adicionar livros às suas listas, fornecendo detalhes como título, autor, edição e número de páginas.
- Categorizar livros em ‘Lidos’, ‘Quero Ler’ e ‘Estou Lendo’ para melhor organização.
- Atualizar facilmente o status de leitura de cada livro.
- Avaliar e escrever resenhas, compartilhando opiniões com a comunidade de leitores.
- Registrar o progresso de leitura, incluindo o número de páginas lidas no total e no dia.
- Salvar citações favoritas e acessá-las a qualquer momento.

## URL Base

A URL base para acessar a API é: `https://api.organum.com`

## Documentação Completa

A URL para acessar a documentação completa: https://documenter.getpostman.com/view/35096123/2sA3QqgYiR

## Endpoints

A API Organum oferece os seguintes endpoints:

### User
- `POST /createuser`: Cadastra um novo usuário.
- `POST /login`: Realiza login de um usuário.
- `POST /logout`: Realiza logout de um usuário.

### Lectio - Endpoints de Leitura
- `POST /lectio/book`: Adiciona um novo livro à lista do usuário.
- `GET /lectio/book/l`: Retorna uma lista de livros baseada no status especificado ('Lidos', 'Quero Ler', 'Estou Lendo').
- `POST /lectio/book/updattag`: Atualiza o status de um livro específico.
- `POST /lectio/review`: Permite que os usuários avaliem e escrevam resenhas sobre um livro.
- `GET /lectio/review`: Retorna resenhas de livros.
- `PUT /lectio/book/progress`: Atualiza o progresso de leitura de um livro.
- `GET /lectio/book/progress`: Retorna o progresso de leitura de um livro.
- `POST /lectio/quote`: Permite que os usuários salvem suas citações favoritas de livros.
- `GET /lectio/quote`: Retorna citações de livros.
- `DELETE /lectio/book`: Deleta um livro da lista do usuário.


## Códigos de Status de Erro

A API Organum retorna os seguintes códigos de status de erro:

- `400 Bad Request`: A requisição foi malformada ou contém dados inválidos.
- `401 Unauthorized`: O usuário não tem permissão para acessar o recurso.
- `404 Not Found`: O recurso solicitado não foi encontrado.
- `500 Internal Server Error`: Ocorreu um erro no servidor.

## Instalação

Instruções para instalar o projeto backend localmente.

### Pré-requisitos

Certifique-se de ter o Go instalado na sua máquina. Você pode baixar e instalar o Go [aqui](https://golang.org/dl/).

```bash
# Verifique a instalação do Go
go version
```

### Passos de instalação local do backend

```bash
# Clone o repositório
git clone https://gitlab.com/ltp2-b-crepusculo/ps-backend-sanderson-machado.git

# Navegue para o diretório do projeto
cd ps-backend-sanderson-machado

# Instale as dependências
go mod tidy

# Faça a build dos arquivos.go
go build *.go

# Rode a aplicação .exe
./organum.exe
