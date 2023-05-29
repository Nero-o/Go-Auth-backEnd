
# Controladores de Autenticação
Este pacote contém os controladores de autenticação para a aplicação CRUD Go.





## Estrutura do projeto
```
- crud_go
  - controllers
    - authControllers.go
  - database
    - connection.go
  - models
    - user.go
  - routes
    - router.go
  - main.go

```
**Explicação da estruturação de arquivos**
- authControllers.go: Este arquivo contém as funções de controle para autenticação, como registro de usuário, login, logout e recuperação de informações do usuário.

- connection.go: Neste arquivo, está a função responsável por estabelecer a conexão com o banco de dados e realizar as migrações, garantindo que as tabelas do banco de dados estejam atualizadas com as definições dos modelos.

- user.go: Este arquivo define a estrutura do modelo User, que representa um usuário no sistema. Ele especifica os campos do usuário, como ID, nome de usuário, email e senha.

- router.go: Aqui, estão definidas as rotas da aplicação utilizando o pacote github.com/gofiber/fiber/v2. As rotas são configuradas para associar cada rota a uma função de controle correspondente, que será executada quando a rota for acessada.

- main.go: O arquivo main.go é o ponto de entrada da aplicação. Ele é responsável por estabelecer a conexão com o banco de dados, configurar as rotas e iniciar o servidor Fiber.


Essa estrutura de pastas e arquivos organiza o projeto em módulos relacionados, separando as responsabilidades de cada componente. Os arquivos nas subpastas são agrupados de acordo com sua função, facilitando a manutenção e a compreensão do código.
## Funcionalidades
**Registro de usuário:**
- O endpoint /api/register permite que um usuário se registre no sistema.
- Os dados enviados são validados, como a verificação de senhas correspondentes, validação de formato de e-mail e requisitos de senha.
- Se o registro for bem-sucedido, o usuário é criado e armazenado no banco de dados.

**Autenticação de usuário:**

- O endpoint /api/login permite que um usuário faça login no sistema.
- Os dados enviados são verificados em relação aos registros de usuário no banco de dados.
- Se as credenciais estiverem corretas, um token JWT é gerado e retornado como resposta.
- O token é armazenado em um cookie com o nome "jwt" e é usado para autenticar as solicitações subsequentes do usuário.

**Recuperação de informações do usuário:**

- O endpoint /api/user é usado para recuperar as informações do usuário autenticado.
- O token JWT armazenado no cookie é verificado para garantir a autenticidade do usuário.
- As informações do usuário são recuperadas do banco de dados e retornadas como resposta.

**Logout do usuário:**

- O endpoint /api/logout é usado para fazer logout do usuário.
- O cookie "jwt" é excluído, invalidando o token JWT anteriormente armazenado.

Essas funcionalidades permitem que os usuários se registrem, façam login, acessem suas informações e façam logout do sistema. O projeto utiliza o Fiber como framework web e o GORM como ORM para interagir com o banco de dados MySQL.
## Rodando localmente

- Certifique-se de que o Go esteja instalado em sua máquina. Você pode baixar e instalar o Go em https://golang.org/dl/. 

- Verifique se o Go está configurado corretamente definindo a variável de ambiente GOPATH e adicionando o diretório bin do Go ao seu PATH.

- Instale o MySQL em sua máquina e crie um banco de dados chamado "users". Certifique-se de ter um usuário e senha válidos para acessar o banco de dados.

- Clone o projeto

```bash
git clone https://github.com/Nero-o/Go-Auth-backEnd
```

- Entre no diretório do projeto

```bash
cd crud-go
```

- Instale as dependências

```bash
  go mod download
```

- Abra o arquivo conecction.go no diretório database e altere a string de conexão para o MySQL de acordo com suas configurações. Substitua "root:1361@/users" pela sua string de conexão válida.

- Inicie o servidor

```bash
  go run main.go
```
O servidor será iniciado na porta 8000.

**Caso deseje utilizar o front que eu desenvolvi para esse projeto acesse:
https://github.com/Nero-o/React-auth-Go-project-**

Agora você pode acessar as rotas do projeto localmente usando um cliente HTTP, como o cURL ou o Postman, para testar as funcionalidades do projeto. Por exemplo, você pode registrar um novo usuário, fazer login, obter informações do usuário e fazer logout. Certifique-se de ajustar as URLs de acordo com a porta e o host local em que o servidor está sendo executado.
## Contribuição
Para contribuir com o projeto, siga as seguintes etapas:

- Faça um fork do repositório.
- Clone o seu fork localmente.
- Crie uma nova branch para a sua feature:
```bash
git checkout -b minha-feature
```
- Implemente as alterações necessárias e faça commit das mudanças:
```bash
git commit -m "Adiciona funcionalidade X"
```
- Envie as mudanças para o seu repositório no GitHub:
```bash
git push origin minha-feature
```

## Autores

- [@Nero-o](https://github.com/Nero-o)

