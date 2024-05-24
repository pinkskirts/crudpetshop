package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	//"fyne.io/fyne/v2/layout"

	"crudpetshop/DB"
	"crudpetshop/crud"
)

func main() {
	// Connects to DB so its reference can be used in CRUD menu
	_, err := DB.Init()
	if err != nil {
		panic(err)
	}

	// Cria uma nova aplicação Fyne
	petshopApp := app.New()

	window := petshopApp.NewWindow("LOGIN")
	window.Resize(fyne.NewSize(400, 300))

	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Login")
	senhaEntry := widget.NewPasswordEntry()
	senhaEntry.SetPlaceHolder("Senha")

	errorMessage := widget.NewLabel("")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Login", Widget: loginEntry},
			{Text: "Senha", Widget: senhaEntry},
		},
		OnSubmit: func() {
			// checar login e senha no bd
			if !queryUser(loginEntry.Text, senhaEntry.Text) {
				errorMessage.SetText("Usuário não encontrado.")
			} else {
				errorMessage.SetText("")
				newTablesWindow(petshopApp)
			}
		},
		OnCancel: func() {
			petshopApp.Quit()
		},
		SubmitText: "Enter",
		CancelText: "Sair",
	}

	content := container.NewVBox(
		widget.NewLabel("Entre com seu login e senha:"),
		form,
		errorMessage,
	)

	window.SetContent(content)
	window.ShowAndRun()
}

func queryUser(login string, senha string) bool {
	Usuarios, err := crud.READUsuario()
	if err != nil {
		panic(err)
	}

	for _, usuario := range Usuarios {
		if usuario.Login == login && usuario.Senha == senha {
			return true // Usuário encontrado
		}
	}

	return false
}

func newTablesWindow(app fyne.App) {
	window := app.NewWindow("TABELAS")
	window.Resize(fyne.NewSize(400, 300))

	tables, err := crud.ShowTables()
	if err != nil {
		panic(err)
	}

	content := container.NewVBox()

	content.Add(widget.NewLabel("Tabelas disponiveis:"))

	for _, table := range tables {
		if table != "Raca" && table != "Produto" && table != "TipoPagamento" && table != "TipoServico" && table != "Porte" && table != "Usuario" {
			content.Add(widget.NewButton(table, func() {
				switch table {
				case "Animal":
					newCRUDwindow(app, 1)
				case "Cliente":
					newCRUDwindow(app, 2)
				case "Funcionario":
					newCRUDwindow(app, 3)
				case "Pagamento":
					newCRUDwindow(app, 4)
				case "Servico":
					newCRUDwindow(app, 5)
				}
			}))
		}

	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)
	window.Show()
}

func newCRUDwindow(app fyne.App, tabela int) {
	// Cria uma nova janela
	window := app.NewWindow("CRUD Petshop")
	window.Resize(fyne.NewSize(400, 200))

	// Cria os botões
	buttonCREATE := widget.NewButton("CREATE", func() {
		switch tabela {
		case 1:
			newAnimalWindow(app)
		case 2:
			newClienteWindow(app)
		case 3:
			newFuncionarioWindow(app)
		case 4:
			newPagamentoWindow(app)
		case 5:
			newServicoWindow(app)
		}
	})

	buttonREAD := widget.NewButton("READ", func() {
		switch tabela {
		case 1:
			newShowAnimaisWindow(app)
		case 2:
			newShowClientesWindow(app)
		case 3:
			newShowFuncionariosWindow(app)
		case 4:
			newShowPagamentosWindow(app)
		case 5:
			newShowServicosWindow(app)
		}
	})

	buttonUPDATE := widget.NewButton("UPDATE", func() {
		switch tabela {
		case 1:
			newUpdateAnimalWindow(app)
		case 2:
			newUpdateClienteWindow(app)
		case 3:
			newUpdateFuncionarioWindow(app)
		case 4:
			newUpdatePagamentoWindow(app)
		case 5:
			newUpdateServicoWindow(app)
		}
	})

	buttonDELETE := widget.NewButton("DELETE", func() {
		switch tabela {
		case 1:
			newRemoveAnimalWindow(app)
		case 2:
			newRemoveClienteWindow(app)
		case 3:
			newRemoveFuncionarioWindow(app)
		case 4:
			newRemovePagamentoWindow(app)
		case 5:
			newRemoveServicoWindow(app)
		}
	})

	buttonEXIT := widget.NewButton("Voltar", func() {
		window.Close()
	})

	// Cria um contêiner para organizar os botões verticalmente
	buttonsContainer := container.NewVBox(
		buttonCREATE,
		buttonREAD,
		buttonUPDATE,
		buttonDELETE,
		buttonEXIT,
	)

	// Adiciona o contêiner de botões à janela
	window.SetContent(buttonsContainer)
	window.Show()
}

func newClienteWindow(app fyne.App) {
	window := app.NewWindow("NOVO CLIENTE")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	nomeEntry := widget.NewEntry()
	cpfEntry := widget.NewEntry()

	content.Add(widget.NewLabel("Insira os dados do cliente:"))

	content.Add(widget.NewLabel("Nome:"))
	content.Add(nomeEntry)

	content.Add(widget.NewLabel("CPF:"))
	content.Add(cpfEntry)

	content.Add(widget.NewButton("Enter", func() {
		crud.CREATECliente(nomeEntry.Text, cpfEntry.Text)
		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newAnimalWindow(app fyne.App) {
	window := app.NewWindow("SELECIONAR DONO DO ANIMAL")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	clientes, err := crud.READCliente()
	if err != nil {
		panic(err)
	}

	for id, cliente := range clientes {
		content.Add(widget.NewButton(cliente.Nome, func() {
			newAnimalInputWindow(app, (id + 1))
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newAnimalInputWindow(app fyne.App, cliente int) {
	window := app.NewWindow("DADOS DO ANIMAL")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	nomeEntry := widget.NewEntry()
	porteEntry := widget.NewEntry()
	racaEntry := widget.NewEntry()

	content.Add(widget.NewLabel("Nome:"))
	content.Add(nomeEntry)

	content.Add(widget.NewLabel("Porte:"))
	content.Add(porteEntry)

	content.Add(widget.NewLabel("Raca:"))
	content.Add(racaEntry)

	content.Add(widget.NewButton("Enter", func() {
		p, _ := strconv.Atoi(porteEntry.Text)
		r, _ := strconv.Atoi(racaEntry.Text)
		crud.CREATEAnimal(nomeEntry.Text, p, r, cliente)

		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newShowAnimaisWindow(app fyne.App) { // todo: mudar output de read Animal para aparecer os campos de texto ao inves de ID, fazer JOIN
	window := app.NewWindow("ANIMAIS DISPONIVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	Animais, _ := crud.READAnimal()

	for _, animal := range Animais {
		content.Add(widget.NewLabel("ID: " + strconv.Itoa(animal.ID)))
		content.Add(widget.NewLabel("Nome: " + animal.Nome))
		content.Add(widget.NewLabel("Porte: " + strconv.Itoa(animal.Porte)))
		content.Add(widget.NewLabel("Raca: " + strconv.Itoa(animal.Raca)))
		content.Add(widget.NewLabel("Cliente: " + strconv.Itoa(animal.Cliente)))
		content.Add(widget.NewLabel("___________________________________"))

	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newRemoveAnimalWindow(app fyne.App) {
	window := app.NewWindow("ESCOLHA UM ANIMAL PARA REMOVER")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	Animais, _ := crud.READAnimal()

	for _, animal := range Animais {
		content.Add(widget.NewButton(animal.Nome, func() {
			crud.DELETERowByID("Animal", "idAnimal", animal.ID)
			window.Close()
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateAnimalWindow(app fyne.App) {
	window := app.NewWindow("ANIMAIS DISPONÍVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	animais, err := crud.READAnimal()
	if err != nil {
		panic(err)
	}

	for _, animal := range animais {
		content.Add(widget.NewButton(animal.Nome, func() {
			newUpdateAnimalInputWindow(app, animal.ID)
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateAnimalInputWindow(app fyne.App, id int) {
	window := app.NewWindow("ESCOLHER CAMPO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	content.Add(widget.NewButton("Nome", func() {
		newNomeInputWindow(app, id)
	}))

	content.Add(widget.NewButton("Porte", func() {
		newPorteInputWindow(app, id)
	}))

	content.Add(widget.NewButton("Raca", func() {
		newRacaInputWindow(app, id)
	}))

	content.Add(widget.NewButton("Cliente/Dono", func() {
		newClienteDonoInputWindow(app, id)
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

// Tabela --> Animal
func newNomeInputWindow(app fyne.App, id int) { // todo: fechar abas de input e listagem de animais (refresh names)
	window := app.NewWindow("NOVO NOME")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	nomeEntry := widget.NewEntry()

	content.Add(widget.NewLabel("Nome:"))
	content.Add(nomeEntry)

	content.Add(widget.NewButton("Enter", func() {
		crud.UPDATEAnimal(id, "nome", nomeEntry.Text)
		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

// Tabela --> Animal
func newPorteInputWindow(app fyne.App, id int) {
	window := app.NewWindow("NOVO PORTE")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	content.Add(widget.NewButton("PEQUENO", func() {
		crud.UPDATEAnimal(id, "porte", "1")
		window.Close()
	}))
	content.Add(widget.NewButton("MEDIO", func() {
		crud.UPDATEAnimal(id, "porte", "2")
		window.Close()
	}))
	content.Add(widget.NewButton("GRANDE", func() {
		crud.UPDATEAnimal(id, "porte", "3")
		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

// Tabela --> Animal
func newRacaInputWindow(app fyne.App, id int) {
	window := app.NewWindow("ATUALIZAR RACA")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	racas, err := crud.READRaca()
	if err != nil {
		panic(err)
	}

	for _, raca := range racas {
		content.Add(widget.NewButton(raca.Nome, func() {
			crud.UPDATEAnimal(id, "raca", strconv.Itoa(raca.ID))
			window.Close()
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

// Tabela --> Animal
func newClienteDonoInputWindow(app fyne.App, id int) {
	window := app.NewWindow("ATUALIZAR DONO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	clientes, err := crud.READCliente()
	if err != nil {
		panic(err)
	}

	for _, cliente := range clientes {
		content.Add(widget.NewButton(cliente.Nome, func() {
			crud.UPDATEAnimal(id, "cliente", strconv.Itoa(cliente.ID))
			window.Close()
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newShowClientesWindow(app fyne.App) {
	window := app.NewWindow("CLIENTES DISPONIVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	Clientes, _ := crud.READCliente()

	for _, cliente := range Clientes {
		content.Add(widget.NewLabel("ID: " + strconv.Itoa(cliente.ID)))
		content.Add(widget.NewLabel("Nome: " + cliente.Nome))
		content.Add(widget.NewLabel("CPF: " + cliente.CPF))
		content.Add(widget.NewLabel("___________________________________"))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newRemoveClienteWindow(app fyne.App) {
	window := app.NewWindow("REMOVER CLIENTE")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	clientes, err := crud.READCliente()
	if err != nil {
		panic(err)
	}

	for _, cliente := range clientes {
		content.Add(widget.NewButton(cliente.Nome, func() {
			crud.DELETERowByID("Cliente", "idcliente", cliente.ID) // todo: mudar atributo idcliente para idCliente
			window.Close()
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newRemoveFuncionarioWindow(app fyne.App) {
	window := app.NewWindow("REMOVER FUNCIONARIO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	funcionarios, err := crud.READFuncionario()
	if err != nil {
		panic(err)
	}

	if len(funcionarios) != 0 { // todo: mostrar mensagem de tabela vazia
		for _, funcionario := range funcionarios {
			content.Add(widget.NewButton(funcionario.Nome, func() {
				crud.DELETERowByID("Funcionario", "idFuncionario", funcionario.ID)
				window.Close()
			}))
		}
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newShowFuncionariosWindow(app fyne.App) {
	window := app.NewWindow("FUNCIONARIOS DISPONIVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	funcionarios, err := crud.READFuncionario()
	if err != nil {
		panic(err)
	}

	for _, funcionario := range funcionarios {
		content.Add(widget.NewLabel("ID: " + strconv.Itoa(funcionario.ID)))
		content.Add(widget.NewLabel("Nome: " + funcionario.Nome))
		content.Add(widget.NewLabel("CPF: " + funcionario.CPF))
		content.Add(widget.NewLabel("___________________________________"))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newFuncionarioWindow(app fyne.App) {
	window := app.NewWindow("NOVO FUNCIONARIO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	nomeEntry := widget.NewEntry()
	cpfEntry := widget.NewEntry()

	content.Add(widget.NewLabel("Insira os dados do funcionario:"))

	content.Add(widget.NewLabel("Nome:"))
	content.Add(nomeEntry)

	content.Add(widget.NewLabel("CPF:"))
	content.Add(cpfEntry)

	content.Add(widget.NewButton("Enter", func() {
		crud.CREATEFuncionario(nomeEntry.Text, cpfEntry.Text)
		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newRemovePagamentoWindow(app fyne.App) {
	window := app.NewWindow("REMOVER PAGAMENTO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	pagamentos, err := crud.READPagamento()
	if err != nil {
		panic(err)
	}

	if len(pagamentos) != 0 {
		for _, pagamento := range pagamentos {
			content.Add(widget.NewButton(strconv.Itoa(pagamento.ID), func() {
				crud.DELETERowByID("Pagamento", "TipoPagamento_idTipoPagamento", pagamento.ID)
				window.Close()
			}))
		}
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newShowPagamentosWindow(app fyne.App) {
	window := app.NewWindow("PAGAMENTOS DISPONIVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	Pagamentos, _ := crud.READPagamento()

	for _, pagamento := range Pagamentos { // refatorar nas outras funcoes de show
		content.Add(widget.NewLabel("ID: " + strconv.Itoa(pagamento.ID)))
		content.Add(widget.NewLabel("Tipo: " + strconv.Itoa(pagamento.Tipo))) // todo: pegar tipo via nome
		content.Add(widget.NewLabel("Data de emissão: " + pagamento.Data))
		content.Add(widget.NewLabel("___________________________________"))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newPagamentoWindow(app fyne.App) {
	window := app.NewWindow("NOVO PAGAMENTO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	content.Add(widget.NewLabel("Escolha o tipo de forma de pagamento:"))

	tipos_pagamento, err := crud.READTipoPagamento()
	if err != nil {
		panic(err)
	}

	// crud.CREATEPagamento() // (1,'Dinheiro'),(2,'Cartão'),(3,'Boleto');
	if len(tipos_pagamento) != 0 {
		for _, tipo_pagamento := range tipos_pagamento {
			content.Add(widget.NewButton(tipo_pagamento.Nome, func() {
				crud.CREATEPagamento(tipo_pagamento.ID)
				window.Close()
			}))
		}
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newRemoveServicoWindow(app fyne.App) {
	window := app.NewWindow("REMOVER SERVICO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	servicos, err := crud.READServico()
	if err != nil {
		panic(err)
	}

	if len(servicos) != 0 {
		for _, servico := range servicos {
			content.Add(widget.NewButton(strconv.Itoa(servico.ID), func() { // todo: mostrar multiplos campos no botao
				crud.DELETERowByID("Servico", "idServico", servico.ID)
				window.Close()
			}))
		}
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newServicoWindow(app fyne.App) { // todo: init.sql --> inserir campo preco no servico
	window := app.NewWindow("NOVO SERVICO")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	// Funcionario_idFuncionario
	content.Add(widget.NewLabel("Funcionario a ser designado:"))

	var funcionarioSelectedID int
	var funcionariosNomes []string

	Funcionarios, err := crud.READFuncionario()
	if err != nil {
		panic(err)
	}

	for _, funcionario := range Funcionarios {
		funcionariosNomes = append(funcionariosNomes, funcionario.Nome)
	}

	selectEntryFuncionario := widget.NewSelect(funcionariosNomes, func(value string) {
		for _, funcionario := range Funcionarios {
			if funcionario.Nome == value {
				funcionarioSelectedID = funcionario.ID
			}
		}
	})

	content.Add(selectEntryFuncionario)

	// TipoServico_idTipoServico
	content.Add(widget.NewLabel("Tipo do serviço:"))

	var tipoServicoSelectedID int
	var tiposServicoNomes []string

	TiposServico, err := crud.READTipoServico()
	if err != nil {
		panic(err)
	}

	for _, tipo_servico := range TiposServico {
		tiposServicoNomes = append(tiposServicoNomes, tipo_servico.Nome)
	}

	selectEntryTipoServico := widget.NewSelect(tiposServicoNomes, func(value string) {
		for _, tipo_servico := range TiposServico {
			if tipo_servico.Nome == value {
				tipoServicoSelectedID = tipo_servico.ID
			}
		}
	})

	content.Add(selectEntryTipoServico)

	// Anima_idAnima + Anima_Cliente_idcliente
	content.Add(widget.NewLabel("Animal correspondente:"))

	var animalSelectedID int
	var clienteSelectedID int
	var animaisNomes []string

	Animais, err := crud.READAnimal()
	if err != nil {
		panic(err)
	}

	for _, animal := range Animais {
		animaisNomes = append(animaisNomes, animal.Nome)
	}

	selectEntryAnimal := widget.NewSelect(animaisNomes, func(value string) {
		for _, animal := range Animais {
			if animal.Nome == value {
				animalSelectedID = animal.ID
				clienteSelectedID = animal.Cliente
			}
		}
	})

	content.Add(selectEntryAnimal)

	// Produto_idProduto
	content.Add(widget.NewLabel("Produto:"))

	var produtoSelectedID int
	var produtoNomes []string

	Produtos, err := crud.READProduto()
	if err != nil {
		panic(err)
	}

	for _, produto := range Produtos {
		produtoNomes = append(produtoNomes, produto.Nome)
	}

	selectEntryProduto := widget.NewSelect(produtoNomes, func(value string) {
		for _, produto := range Produtos {
			if produto.Nome == value {
				produtoSelectedID = produto.ID
			}
		}
	})

	content.Add(selectEntryProduto)

	// Pagamento_idFormaPagamento
	content.Add(widget.NewLabel("Forma de pagamento:"))

	var pagamentoSelectedID int
	var tiposPagamentoNomes []string

	TiposPagamento, err := crud.READTipoPagamento()
	if err != nil {
		panic(err)
	}

	for _, tipo_pagamento := range TiposPagamento {
		tiposPagamentoNomes = append(tiposPagamentoNomes, tipo_pagamento.Nome)
	}

	selectEntryTipoPagamento := widget.NewSelect(tiposPagamentoNomes, func(value string) {
		for _, tipo_pagamento := range TiposPagamento {
			if tipo_pagamento.Nome == value {
				pagamentoSelectedID = tipo_pagamento.ID
			}
		}
	})

	content.Add(selectEntryTipoPagamento)

	content.Add(widget.NewButton("Enter", func() {
		crud.CREATEServico(funcionarioSelectedID, tipoServicoSelectedID, animalSelectedID, clienteSelectedID, produtoSelectedID, pagamentoSelectedID)
		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newShowServicosWindow(app fyne.App) {
	window := app.NewWindow("SERVICOS REALIZADOS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	Servicos, _ := crud.READServico()

	for _, servico := range Servicos {
		content.Add(widget.NewLabel("ID: " + strconv.Itoa(servico.ID)))
		content.Add(widget.NewLabel("Funcionario: " + strconv.Itoa(servico.IDFuncionario)))
		content.Add(widget.NewLabel("Data: " + servico.Data))
		content.Add(widget.NewLabel("Tipo de serviço: " + strconv.Itoa(servico.Tipo)))
		content.Add(widget.NewLabel("Animal: " + strconv.Itoa(servico.IDAnimal)))
		content.Add(widget.NewLabel("ID do Cliente: " + strconv.Itoa(servico.IDClienteAnimal)))
		content.Add(widget.NewLabel("ID do Produto: " + strconv.Itoa(servico.IDProduto)))
		content.Add(widget.NewLabel("ID Forma de pagamento: " + strconv.Itoa(servico.IDFormaPagamento)))
		content.Add(widget.NewLabel("___________________________________"))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateClienteWindow(app fyne.App) {
	window := app.NewWindow("CLIENTES DISPONÍVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	clientes, err := crud.READCliente()
	if err != nil {
		panic(err)
	}

	for _, cliente := range clientes {
		content.Add(widget.NewButton(cliente.Nome, func() {
			newUpdateClienteInputWindow(app, cliente.ID)
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateClienteInputWindow(app fyne.App, id int) {
	window := app.NewWindow("CAMPOS A SEREM MODIFICADOS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	nomeEntry := widget.NewEntry()
	cpfEntry := widget.NewEntry()

	content.Add(widget.NewLabel("Insira os novos dados do cliente:"))

	content.Add(widget.NewLabel("Nome:"))
	content.Add(nomeEntry)

	content.Add(widget.NewLabel("CPF:"))
	content.Add(cpfEntry)

	content.Add(widget.NewButton("Enter", func() {
		if cpfEntry.Text == "" {
			crud.UPDATECliente(id, "nome", nomeEntry.Text)
		} else if nomeEntry.Text == "" {
			crud.UPDATECliente(id, "cpf", cpfEntry.Text)
		} else {
			crud.UPDATECliente(id, "nome", nomeEntry.Text)
			crud.UPDATECliente(id, "cpf", cpfEntry.Text)
		}

		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateFuncionarioWindow(app fyne.App) {
	window := app.NewWindow("FUNCIONARIOS DISPONÍVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	funcionarios, err := crud.READFuncionario()
	if err != nil {
		panic(err)
	}

	for _, funcionario := range funcionarios {
		content.Add(widget.NewButton(funcionario.Nome, func() {
			newUpdateFuncionarioInputWindow(app, funcionario.ID)
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateFuncionarioInputWindow(app fyne.App, id int) {
	window := app.NewWindow("CAMPOS A SEREM MODIFICADOS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	nomeEntry := widget.NewEntry()
	cpfEntry := widget.NewEntry()

	content.Add(widget.NewLabel("Insira os novos dados do funcionario:"))

	content.Add(widget.NewLabel("Nome:"))
	content.Add(nomeEntry)

	content.Add(widget.NewLabel("CPF:"))
	content.Add(cpfEntry)

	content.Add(widget.NewButton("Enter", func() {
		if cpfEntry.Text == "" {
			crud.UPDATEFuncionario(id, "nome", nomeEntry.Text)
		} else if nomeEntry.Text == "" {
			crud.UPDATEFuncionario(id, "cpf", cpfEntry.Text)
		} else {
			crud.UPDATEFuncionario(id, "nome", nomeEntry.Text)
			crud.UPDATEFuncionario(id, "cpf", cpfEntry.Text)
		}

		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdatePagamentoWindow(app fyne.App) {
	window := app.NewWindow("PAGAMENTOS DISPONÍVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	pagamentos, err := crud.READPagamento()
	if err != nil {
		panic(err)
	}

	for _, pagamento := range pagamentos {
		content.Add(widget.NewButton(pagamento.Data, func() {
			newUpdatePagamentoInputWindow(app, pagamento.ID)
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdatePagamentoInputWindow(app fyne.App, id int) {
	window := app.NewWindow("CAMPOS A SEREM MODIFICADOS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	// Pagamento_idFormaPagamento
	content.Add(widget.NewLabel("Forma de pagamento:"))

	var tipo_pagamentoSelectedID int
	var tiposPagamentoNomes []string

	TiposPagamento, err := crud.READTipoPagamento()
	if err != nil {
		panic(err)
	}

	for _, tipo_pagamento := range TiposPagamento {
		tiposPagamentoNomes = append(tiposPagamentoNomes, tipo_pagamento.Nome)
	}

	selectEntryTipoPagamento := widget.NewSelect(tiposPagamentoNomes, func(value string) {
		for _, tipo_pagamento := range TiposPagamento {
			if tipo_pagamento.Nome == value {
				tipo_pagamentoSelectedID = tipo_pagamento.ID
			}
		}
	})

	content.Add(selectEntryTipoPagamento)

	content.Add(widget.NewButton("Enter", func() {
		crud.UPDATEPagamento(id, "TipoPagamento_idTipoPagamento", strconv.Itoa(tipo_pagamentoSelectedID))
		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateServicoWindow(app fyne.App) {
	window := app.NewWindow("SERVICOS DISPONÍVEIS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	servicos, err := crud.READServico()
	if err != nil {
		panic(err)
	}

	for _, servico := range servicos {
		content.Add(widget.NewButton(strconv.Itoa(servico.ID), func() {
			newUpdateServicoInputWindow(app, servico.ID)
		}))
	}

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)

	window.Show()
}

func newUpdateServicoInputWindow(app fyne.App, id int) {
	window := app.NewWindow("CAMPOS A SEREM MODIFICADOS")
	window.Resize(fyne.NewSize(400, 300))

	content := container.NewVBox()

	// Funcionario_idFuncionario
	content.Add(widget.NewLabel("Funcionario a ser designado:"))

	var funcionarioSelectedID int
	var funcionariosNomes []string
	funcionarioRadio := widget.NewRadioGroup([]string{"Modificar Funcionario"}, nil)

	Funcionarios, err := crud.READFuncionario()
	if err != nil {
		panic(err)
	}

	for _, funcionario := range Funcionarios {
		funcionariosNomes = append(funcionariosNomes, funcionario.Nome)
	}

	selectEntryFuncionario := widget.NewSelect(funcionariosNomes, func(value string) {
		for _, funcionario := range Funcionarios {
			if funcionario.Nome == value {
				funcionarioSelectedID = funcionario.ID
			}
		}
	})

	content.Add(funcionarioRadio)
	content.Add(selectEntryFuncionario)

	// TipoServico_idTipoServico
	content.Add(widget.NewLabel("Tipo do serviço:"))

	var tipoServicoSelectedID int
	var tiposServicoNomes []string
	tipoServicoRadio := widget.NewRadioGroup([]string{"Modificar Tipo do Serviço"}, nil)

	TiposServico, err := crud.READTipoServico()
	if err != nil {
		panic(err)
	}

	for _, tipo_servico := range TiposServico {
		tiposServicoNomes = append(tiposServicoNomes, tipo_servico.Nome)
	}

	selectEntryTipoServico := widget.NewSelect(tiposServicoNomes, func(value string) {
		for _, tipo_servico := range TiposServico {
			if tipo_servico.Nome == value {
				tipoServicoSelectedID = tipo_servico.ID
			}
		}
	})

	content.Add(tipoServicoRadio)
	content.Add(selectEntryTipoServico)

	// Anima_idAnima + Anima_Cliente_idcliente
	content.Add(widget.NewLabel("Animal correspondente:"))

	var animalSelectedID int
	var clienteSelectedID int
	var animaisNomes []string
	animalRadio := widget.NewRadioGroup([]string{"Modificar Animal"}, nil)

	Animais, err := crud.READAnimal()
	if err != nil {
		panic(err)
	}

	for _, animal := range Animais {
		animaisNomes = append(animaisNomes, animal.Nome)
	}

	selectEntryAnimal := widget.NewSelect(animaisNomes, func(value string) {
		for _, animal := range Animais {
			if animal.Nome == value {
				animalSelectedID = animal.ID
				clienteSelectedID = animal.Cliente
			}
		}
	})

	content.Add(animalRadio)
	content.Add(selectEntryAnimal)

	// Produto_idProduto
	content.Add(widget.NewLabel("Produto:"))

	var produtoSelectedID int
	var produtoNomes []string
	produtoRadio := widget.NewRadioGroup([]string{"Modificar Produto"}, nil)

	Produtos, err := crud.READProduto()
	if err != nil {
		panic(err)
	}

	for _, produto := range Produtos {
		produtoNomes = append(produtoNomes, produto.Nome)
	}

	selectEntryProduto := widget.NewSelect(produtoNomes, func(value string) {
		for _, produto := range Produtos {
			if produto.Nome == value {
				produtoSelectedID = produto.ID
			}
		}
	})

	content.Add(produtoRadio)
	content.Add(selectEntryProduto)

	// Pagamento_idFormaPagamento
	content.Add(widget.NewLabel("Forma de pagamento:"))

	var pagamentoSelectedID int
	var tiposPagamentoNomes []string
	pagamentoRadio := widget.NewRadioGroup([]string{"Modificar Forma de Pagamento"}, nil)

	TiposPagamento, err := crud.READTipoPagamento()
	if err != nil {
		panic(err)
	}

	for _, tipo_pagamento := range TiposPagamento {
		tiposPagamentoNomes = append(tiposPagamentoNomes, tipo_pagamento.Nome)
	}

	selectEntryTipoPagamento := widget.NewSelect(tiposPagamentoNomes, func(value string) {
		for _, tipo_pagamento := range TiposPagamento {
			if tipo_pagamento.Nome == value {
				pagamentoSelectedID = tipo_pagamento.ID
			}
		}
	})

	content.Add(pagamentoRadio)
	content.Add(selectEntryTipoPagamento)

	content.Add(widget.NewButton("Enter", func() {
		if funcionarioRadio.Selected == "Modificar Funcionario" {
			crud.UPDATEServico(id, "Funcionario_idFuncionario", funcionarioSelectedID)
		}
		if tipoServicoRadio.Selected == "Modificar Tipo do Serviço" {
			crud.UPDATEServico(id, "TipoServico_idTipoServico", tipoServicoSelectedID)
		}
		if animalRadio.Selected == "Modificar Animal" {
			crud.UPDATEServico(id, "Anima_idAnima", animalSelectedID)
			crud.UPDATEServico(id, "Anima_cliente_idcliente", clienteSelectedID)
		}
		if produtoRadio.Selected == "Modificar Produto" {
			crud.UPDATEServico(id, "Produto_idProduto", produtoSelectedID)
		}
		if pagamentoRadio.Selected == "Modificar Forma de Pagamento" {
			crud.UPDATEServico(id, "Pagamento_idFormaPagamento", pagamentoSelectedID)
		}

		window.Close()
	}))

	content.Add(widget.NewButton("Voltar", func() {
		window.Close()
	}))

	window.SetContent(content)
	window.Show()
}
