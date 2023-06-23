# API

---

## Como Executar a Aplicação

---

Para executar a aplicação, siga as etapas abaixo:

1. Certifique-se de ter o Go instalado em sua máquina. Para mais informações sobre como instalar o Go, consulte a documentação oficial em https://golang.org/doc/install.

2. Faça o download ou clone este repositório para o seu ambiente local.

3. Navegue até a pasta raiz da aplicação.

4. Instale as dependências executando o seguinte comando:

```bash
    go mod download

```

5. Configure as informações necessárias no arquivo config/config.go, como detalhes de conexão com o banco de dados, configurações do servidor, chaves de autenticação, etc.

6. Execute o seguinte comando para iniciar a aplicação:


```bash
    go run main.go
```

or:

```bash
    go run api
```

7. A API agora está em execução localmente. Você pode acessá-la em http://localhost:5000 ou no endereço configurado no arquivo config/config.go



## Pastas do projeto

---

A seguir, uma breve explicação sobre as pastas e arquivos importantes da sua API:

- banco: Esta pasta contém os arquivos relacionados à configuração e interação com o banco de dados da aplicação.

- autenticacao: Aqui estão os arquivos relacionados à autenticação e autorização da API. Isso inclui middleware de autenticação, gerenciamento de tokens, etc.

- config: A pasta de configuração contém os arquivos de configuração da aplicação, como conexões com bancos de dados, configurações de servidor, chaves de autenticação, etc.

- controllers: Nesta pasta, você encontrará os controladores da API. Os controladores são responsáveis por receber as requisições HTTP, processá-las e retornar as respostas adequadas.

- modelos: Aqui estão os modelos de dados da aplicação. Os modelos definem a estrutura dos dados que serão utilizados pela API.

- repositorios: Os repositórios são responsáveis pela interação direta com o banco de dados. Esta pasta contém os arquivos que implementam as operações de acesso ao banco de dados para os modelos de dados.

- router: O arquivo router/router.go contém as definições de rotas da API. Aqui você configura as rotas disponíveis, associa cada rota a um controlador e define os métodos HTTP permitidos para cada rota.
