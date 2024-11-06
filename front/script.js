// Função para enviar um novo cliente
document.getElementById('cliente-form').addEventListener('submit', async function(event) {
    event.preventDefault(); // Evita o envio padrão do formulário

    const urlAddress = 'http://localhost:8080/cliente/create';
    const name = document.getElementById('cliente-name').value;
    const sexo = document.getElementById('cliente-sexo').value;
    const idade = parseInt(document.getElementById('cliente-idade').value);
    const nascimento = document.getElementById('cliente-nascimento').value;

    try {
        // Enviar dados para o endpoint da API
        const response = await fetch(urlAddress, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name, sexo, idade, nascimento })
        });

        if (response.ok) {
            alert('Cliente adicionado com sucesso!');
            fetchClientes(); // Atualiza a lista de clientes
        } else {
            alert('Erro ao adicionar cliente.');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Falha ao enviar os dados.');
    }
});

async function fetchClientes() { // Função para buscar os clientes da API
    try {
        const response = await fetch('http://localhost:8080/cliente');
        const data = await response.json();
        displayData(data, 'Clientes');
    } catch (error) {
        console.error('Error fetching clientes:', error);
    }
}

async function fetchFornecedores() { // Função para buscar os fornecedores da API
    try {
        const response = await fetch('http://localhost:8080/fornecedor');
        const data = await response.json();
        displayData(data, 'Fornecedores');
    } catch (error) {
        console.error('Error fetching fornecedores:', error);
    }
}

async function fetchIngredientes() { // Função para buscar os ingredientes da API
    try {
        const response = await fetch('http://localhost:8080/ingrediente');
        const data = await response.json();
        displayData(data, 'Ingredientes');
    } catch (error) {
        console.error('Error fetching ingredientes:', error);
    }
}

async function fetchPratos() { // Função para buscar os pratos da API
    try {
        const response = await fetch('http://localhost:8080/prato');
        const data = await response.json();
        displayData(data, 'Pratos');
    } catch (error) {
        console.error('Error fetching pratos:', error);
    }
}

async function fetchUsos() { // Função para buscar os usos/funções da API
    try {
        const response = await fetch('http://localhost:8080/usos');
        const data = await response.json();
        displayData(data, 'Usos');
    } catch (error) {
        console.error('Error fetching usos:', error);
    }
}

async function fetchVendas() { // Função para buscar as vendas da API
    try {
        const response = await fetch('http://localhost:8080/venda');
        const data = await response.json();
        displayData(data, 'Vendas');
    } catch (error) {
        console.error('Error fetching vendas:', error);
    }
}

// Função para exibir os dados no container
function displayData(data, title) {
    const dataOutput = document.getElementById('data-output');
    dataOutput.innerHTML = `<h3>${title}</h3>`;

    data.forEach(item => {
        const div = document.createElement('div');
        div.textContent = JSON.stringify(item);
        dataOutput.appendChild(div);
    });
}
