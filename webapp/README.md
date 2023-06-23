# WebApp

---

## Como Executar a Aplicação

---

Para executar a aplicação, siga as etapas abaixo:
Este é o README da parte web (webapp) da sua aplicação em Go, que contém informações importantes sobre como executar e utilizar o webapp.

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
    go run webapp
```

## Entendendo a árvore de pastas do Webapp


O webapp da aplicação foi desenvolvido em Go e possui as seguintes partes principais:

- main.go: Este é o ponto de entrada do webapp. Ele configura as rotas e os manipuladores HTTP para cada rota.

- src/config: O diretório "config" contém arquivos relacionados à configuração do webapp, como arquivos de configuração ou variáveis de ambiente.

- src/controllers: O diretório "controllers" contém os controladores do webapp, que são responsáveis por processar as requisições HTTP e retornar as respostas apropriadas. Cada controlador pode lidar com um conjunto específico de rotas relacionadas.

- src/cookies: O diretório "cookies" contém arquivos relacionados ao gerenciamento de cookies no webapp.

- src/utils: O diretório "utils" contém funções auxiliares que podem ser utilizadas em diferentes partes do webapp.

- view: O diretório "view" contém os templates utilizados para renderizar as páginas HTML do webapp. Você pode adicionar ou modificar os templates de acordo com as necessidades do seu webapp.

- assets: O diretório "assets" contém subpastas para arquivos JavaScript, CSS e outros arquivos estáticos necessários para o funcionamento correto do webapp.