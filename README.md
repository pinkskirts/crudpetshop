# Primeiros Passos

## Requisitos
- Sistema operacional baseado em Unix
- Instância do MySQL Client em execução
- Go (Golang) versão 1.22.3 ou superior
- fyne.io/fyne/v2 v2.4.5
- go-sql-driver/mysql v1.8.1
- joho/godotenv v1.5.1

## Instalação

1. Clone o repositório para sua máquina local:
   ```
   git clone https://github.com/pinkskirts/crudpetshop.git
   ```

2. Navegue até o diretório do projeto:
   ```
   cd ./crudpetshop
   ```

3. Instale as dependências:
   ```
   go mod tidy
   ```

4. Conectar-se à instância do MySQL Client
   ```
   mysql -u username -p -h hostname -P port
   ```

5. Execute o script de inicialização do Banco de Dados padrão do projeto - database petshop
   ```
   source ./script.sql
   ```
   Caso o script NÃO estiver no mesmo diretório de onde iniciou o MySQL Client, basta inserir o caminho relativo até ao arquivo.

6. Crie um arquivo .env no root do projeto e insira as informações para conexão do banco de dados <br >

   Por exemplo:
   ```
   DB_USER="root"
   DB_PASSWORD=""
   DB_ADDR="127.0.0.1:3306"
   DB_NAME="petshop" 
   ```
   <b>Observação: Para que a aplicação estabeleça conexão a um database nomeado "petshop", a variável de ambiente DB_NAME não deve ser diferente da deste exemplo.<b/>

## Uso

1. Inicie a aplicação:
   ```
   go run main.go
   ```

2. Após a aparição da tela da aplicação, realize o login como administrador
   ```
   login: admin
   senha: petshop
   ```

3. Selecione a tabela que preferir modificar.
   * Tabelas disponíveis:
       - Animal
       - Cliente
       - Funcionario
       - Pagamento
       - Servico
        
5. Selecione a operação CRUD desejada
   * CREATE
   * READ
   * UPDATE
   * DELETE

6. Siga as instruções das janelas subsequentes

7. Caso deseje encerrar a aplicação, clique em "Sair" na janela de login

## Estrutura do Projeto
   * main.go - Ponto central da aplicação, responsável pela chamada dos métodos da leitura do BD;
   * db.go - Estabelece a conexão com o banco de dados, utilizando variáveis de ambiente para configuração;
   * init.sql - Script inicial do Banco de Dados MySql;
   * crud.go - Realiza as operações CRUD com base em um ponteiro da instância da conexão do BD.
