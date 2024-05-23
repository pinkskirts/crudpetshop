// todo: refactor db pointer

package crud

import (
	"database/sql"
	"fmt"
	"log"

	"strconv"
	"time"

	"crudpetshop/DB"
)

type Animal struct {
	ID      int    // idAnimal
	Nome    string // nome
	Porte   int    // Porte_idPorte
	Raca    int    // Raca_idRaca
	Cliente int    // cliente_idcliente -- change
}

type Cliente struct {
	ID   int    // idcliente -- change
	Nome string // nome
	CPF  string // cpf
}

type Funcionario struct {
	ID   int    // idFuncionario
	Nome string // nome
	CPF  string // cpf
}

type Pagamento struct {
	ID   int    // idFormaPagamento
	Tipo int    // TipoPagamento_idTipoPagamento
	Data string // data
}

type Servico struct {
	ID               int    // idServiço
	IDFuncionario    int    // Funcionario_idFuncionario
	Data             string // data
	Tipo             int    // TipoServico_idTipoServico
	IDAnimal         int    // Animal_idAnimal
	IDClienteAnimal  int    // Anima_cliente_idcliente -- change
	IDProduto        int    // Produto_idProduto
	IDFormaPagamento int    // Pagamento_idFormaPagamento
}

type Raca struct {
	ID   int    // idRaca
	Nome string // nome
}

type TipoPagamento struct {
	ID   int    // idTipoPagamento
	Nome string // nome
}

type TipoServico struct {
	ID   int    // idTipoServico
	Nome string // nome
}

type Produto struct {
	ID    int     // idProduto
	Nome  string  // nome
	Preco float64 // preço
}

type Usuario struct {
	ID    int    // idUsuario
	Login string // login
	Senha string // senha
}

const timeLayout string = "2006-01-02"

// OTHER METHODS

func checkNullDb() {
	var db *sql.DB = DB.DbRef
	if db == nil {
		log.Fatal("DB doesn't exist!")
	}
}

func ShowTables() ([]string, error) {
	var db *sql.DB = DB.DbRef

	// Consulta SQL para listar todas as tabelas no banco de dados atual
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, fmt.Errorf("error while displaying tables: %w", err)
	}
	defer rows.Close()

	// Variável para armazenar os nomes das tabelas
	var tables []string

	// Iterar sobre os resultados da consulta
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %w", err)
		}
		tables = append(tables, tableName)
	}

	// Verificar se houve algum erro durante a iteração
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return tables, nil
}

// CREATE METHODS

// Tabela - Animal
func CREATEAnimal(nome string, Porte_idPorte int, Raca_idRaca int, cliente_idcliente int) (int64, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	result, err := db.Exec("INSERT INTO Animal (nome, Porte_idPorte, Raca_idRaca, cliente_idcliente) VALUES (?, ?, ?, ?)", nome, Porte_idPorte, Raca_idRaca, int64(cliente_idcliente))
	if err != nil {
		return 0, fmt.Errorf("createAnimal: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addAnimal: %v", err)
	}

	fmt.Println("\nAnimal created!")

	return id, nil
}

// Tabela - Cliente
func CREATECliente(nome string, cpf string) (int64, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	result, err := db.Exec("INSERT INTO Cliente (nome, cpf) VALUES (?, ?)", nome, cpf)
	if err != nil {
		return 0, fmt.Errorf("createCliente: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("createCliente: %v", err)
	}

	fmt.Println("\nCliente created!")

	return id, nil
}

func CREATEFuncionario(nome string, cpf string) (int64, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	result, err := db.Exec("INSERT INTO Funcionario (nome, cpf) VALUES (?, ?)", nome, cpf)
	if err != nil {
		return 0, fmt.Errorf("createFuncionario: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("createFuncionario: %v", err)
	}

	fmt.Println("\nFuncionario created!")

	return id, nil
}

func CREATEPagamento(tipo int) (int64, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	dateNow := time.Now()

	parsedDate := dateNow.Format(timeLayout)

	result, err := db.Exec("INSERT INTO Pagamento (TipoPagamento_idTipoPagamento, data) VALUES (?, ?)", tipo, parsedDate)
	if err != nil {
		return 0, fmt.Errorf("createPagamento: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("createPagamento: %v", err)
	}

	fmt.Println("\nPagamento created!")

	return id, nil
}

func CREATEServico(idFuncionario int, idTipoServico int, idAnimal int, idAnimalCliente int, idProduto int, idFormaPagamento int) (int64, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	dateNow := time.Now()

	parsedDate := dateNow.Format(timeLayout)

	result, err := db.Exec("INSERT INTO Servico (Funcionario_idFuncionario, data, TipoServico_idTipoServico, Anima_idAnima, Anima_cliente_idcliente, Produto_idProduto, Pagamento_idFormaPagamento) VALUES (?, ?, ?, ?, ?, ?, ?)", idFuncionario, parsedDate, idTipoServico, idAnimal, idAnimalCliente, idProduto, idFormaPagamento)
	if err != nil {
		return 0, fmt.Errorf("createServico: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("createServico: %v", err)
	}

	fmt.Println("\nServico created!")

	return id, nil
}

// READ METHODS

//func READtable() { todo: fazer metodo read generico para todas as structs

//}

// Tabela - Animal
func READAnimal() ([]Animal, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Animais []Animal

	rows, err := db.Query("SELECT * FROM Animal")
	if err != nil {
		return nil, fmt.Errorf("readAnimal: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var animal Animal

		if err := rows.Scan(&animal.ID, &animal.Nome, &animal.Porte, &animal.Raca, &animal.Cliente); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		Animais = append(Animais, animal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readAnimal: %w", err)
	}

	return Animais, nil
}

// Tabela - Cliente
func READCliente() ([]Cliente, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Clientes []Cliente

	rows, err := db.Query("SELECT * FROM Cliente")
	if err != nil {
		return nil, fmt.Errorf("readCliente: %w", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var cliente Cliente

		if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.CPF); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		Clientes = append(Clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readAnimal: %w", err)
	}

	return Clientes, nil
}

func READRaca() ([]Raca, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Racas []Raca

	rows, err := db.Query("SELECT * FROM Raca")
	if err != nil {
		return nil, fmt.Errorf("readRaca: %w", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var raca Raca

		if err := rows.Scan(&raca.ID, &raca.Nome); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		Racas = append(Racas, raca)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readRaca: %w", err)
	}

	return Racas, nil
}

func READFuncionario() ([]Funcionario, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Funcionarios []Funcionario

	rows, err := db.Query("SELECT * FROM Funcionario")
	if err != nil {
		return nil, fmt.Errorf("readFuncionario: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var funcionario Funcionario

		if err := rows.Scan(&funcionario.ID, &funcionario.Nome, &funcionario.CPF); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		Funcionarios = append(Funcionarios, funcionario)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readFuncionario: %w", err)
	}

	return Funcionarios, nil
}

func READPagamento() ([]Pagamento, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Pagamentos []Pagamento

	rows, err := db.Query("SELECT * FROM Pagamento")
	if err != nil {
		return nil, fmt.Errorf("readPagamento: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var pagamento Pagamento

		if err := rows.Scan(&pagamento.ID, &pagamento.Tipo, &pagamento.Data); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		t, err := time.Parse(time.RFC3339, pagamento.Data)
		if err != nil {
			return nil, fmt.Errorf("error while parsing the time string: %v", err)
		}

		// Remove timestamp, maintain data
		dateWithoutTime := t.Format(timeLayout)

		// parsing - time.Time to string
		pagamento.Data = dateWithoutTime

		Pagamentos = append(Pagamentos, pagamento)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readFuncionario: %w", err)
	}

	return Pagamentos, nil
}

func READTipoPagamento() ([]TipoPagamento, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var TipoPagamentos []TipoPagamento

	rows, err := db.Query("SELECT * FROM TipoPagamento")
	if err != nil {
		return nil, fmt.Errorf("readTipoPagamento: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var tipo_pagamento TipoPagamento

		if err := rows.Scan(&tipo_pagamento.ID, &tipo_pagamento.Nome); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		TipoPagamentos = append(TipoPagamentos, tipo_pagamento)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readFuncionario: %w", err)
	}

	return TipoPagamentos, nil
}

func READServico() ([]Servico, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Servicos []Servico

	rows, err := db.Query("SELECT * FROM Servico")
	if err != nil {
		return nil, fmt.Errorf("readServico: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var servico Servico

		if err := rows.Scan(&servico.ID, &servico.IDFuncionario, &servico.Data, &servico.Tipo, &servico.IDAnimal, &servico.IDClienteAnimal, &servico.IDProduto, &servico.IDFormaPagamento); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		t, err := time.Parse(time.RFC3339, servico.Data)
		if err != nil {
			return nil, fmt.Errorf("error while parsing the time string: %v", err)
		}

		// Remove timestamp, maintain data
		dateWithoutTime := t.Format(timeLayout)

		// parsing - time.Time to string
		servico.Data = dateWithoutTime

		Servicos = append(Servicos, servico)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readServico: %w", err)
	}

	return Servicos, nil
}

func READTipoServico() ([]TipoServico, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var TipoServicos []TipoServico

	rows, err := db.Query("SELECT * FROM TipoServico")
	if err != nil {
		return nil, fmt.Errorf("readTipoServico: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var tipo_servico TipoServico

		if err := rows.Scan(&tipo_servico.ID, &tipo_servico.Nome); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		TipoServicos = append(TipoServicos, tipo_servico)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readTipoServico: %w", err)
	}

	return TipoServicos, nil
}

// (1,'Banho e tosa',100),(2,'Banhho',60),(3,'Bolinha',5),(4,'Cama  50cm x 100cm',70),(5,'Cama  50cm x 50cm',50),(6,'Cama  70cm x 70cm',60),(7,'Shampoo 500ml',18.32),(8,'Shampoo 1000ml',25.5),(9,'Consultas ',80.5);
func READProduto() ([]Produto, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Produtos []Produto

	rows, err := db.Query("SELECT * FROM Produto")
	if err != nil {
		return nil, fmt.Errorf("readProduto: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var produto Produto

		if err := rows.Scan(&produto.ID, &produto.Nome, &produto.Preco); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		Produtos = append(Produtos, produto)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readProduto: %w", err)
	}

	return Produtos, nil
}

func READUsuario() ([]Usuario, error) {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	var Usuarios []Usuario

	rows, err := db.Query("SELECT * FROM Usuario")
	if err != nil {
		return nil, fmt.Errorf("readUsuario: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var usuario Usuario

		if err := rows.Scan(&usuario.ID, &usuario.Login, &usuario.Senha); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		Usuarios = append(Usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("readUsuario: %w", err)
	}

	return Usuarios, nil
}

// UPDATE METHODS

// Tabela - Animal
// todo: mudar campo cliente_idcliente para Cliente_idCliente
func UPDATEAnimal(id int, attribute string, value string) error {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	switch attribute {
	case "nome":
		fmt.Println(id)
		_, err := db.Exec("UPDATE Animal SET nome = ? WHERE idAnimal = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateAnimal: %v", err)
		}
	case "porte": //todo: usar sprintf para remover switch, tratar apenas a entrada de string com if else
		value_aux, _ := strconv.Atoi(value)

		_, err := db.Exec("UPDATE Animal SET Porte_idPorte = ? WHERE idAnimal = ?", value_aux, int64(id))
		if err != nil {
			return fmt.Errorf("updateAnimal: %v", err)
		}
	case "raca":
		_, err := db.Exec("UPDATE Animal SET Raca_idRaca = ? WHERE idAnimal = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateAnimal: %v", err)
		}
	case "cliente":
		_, err := db.Exec("UPDATE Animal SET cliente_idcliente = ? WHERE idAnimal = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateAnimal: %v", err)
		}
	}

	return nil
}

func UPDATECliente(id int, attribute string, value string) error {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	switch attribute {
	case "nome":
		_, err := db.Exec("UPDATE Cliente SET nome = ? WHERE idcliente = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateCliente: %v", err)
		}
	case "cpf":
		_, err := db.Exec("UPDATE Cliente SET cpf = ? WHERE idcliente = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateCliente: %v", err)
		}
	}

	return nil
}

func UPDATEFuncionario(id int, attribute string, value string) error {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	switch attribute {
	case "nome":
		_, err := db.Exec("UPDATE Funcionario SET nome = ? WHERE idFuncionario = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateFuncionario: %v", err)
		}
	case "cpf":
		_, err := db.Exec("UPDATE Funcionario SET cpf = ? WHERE idFuncionario = ?", value, int64(id))
		if err != nil {
			return fmt.Errorf("updateFuncionario: %v", err)
		}
	}

	fmt.Println("Funcionario updated!")
	return nil
}

func UPDATEPagamento(id int, attribute string, value string) error {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	_, err := db.Exec("UPDATE Pagamento SET TipoPagamento_idTipoPagamento = ? WHERE idFormaPagamento = ?", value, int64(id))
	if err != nil {
		return fmt.Errorf("updatePagamento: %v", err)
	}

	fmt.Println("Pagamento updated!")
	return nil
}

func UPDATEServico(id int, id_attribute string, value int) error {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	fmt.Printf("Updating Servico: SET %s = %d WHERE idServico = %d\n", id_attribute, value, id)

	query := fmt.Sprintf("UPDATE Servico SET %s = ? WHERE idServico = ?", id_attribute)
	_, err := db.Exec(query, value, id)
	if err != nil {
		return fmt.Errorf("updateServico: %v", err)
	}

	fmt.Println("Servico updated!")
	return nil
}

// DELETE METHOD

func DELETERowByID(table string, id_attribute string, id int) error {
	var db *sql.DB = DB.DbRef
	checkNullDb()

	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = ?", table, id_attribute), int64(id))
	if err != nil {
		return fmt.Errorf("updateAnimal: %v", err)
	}

	fmt.Println("\nRow deleted!")
	return nil
}
