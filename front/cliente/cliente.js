async function adicionarCliente() { // Add Client
    const nome = document.getElementById('cliente-nome').value;
    const sexo = document.getElementById('cliente-sexo').value;
    const idade = parseInt(document.getElementById('cliente-idade').value);
    const nascimento = document.getElementById('cliente-nascimento').value;

    const response = await fetch('http://localhost:8080/cliente/create', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ nome, sexo, idade, nascimento })
    });

    if (response.ok) {
        alert('Cliente adicionado com sucesso!');
    } else {
        alert('Erro ao adicionar cliente.');
    }
}

async function listarClientes() { // Client List
    const response = await fetch('http://localhost:8080/cliente');
    const clientes = await response.json();
    const listaClientesDiv = document.getElementById('lista-clientes');
    listaClientesDiv.innerHTML = '';

    clientes.forEach(cliente => {
        const div = document.createElement('div');
        div.textContent = JSON.stringify(cliente);
        listaClientesDiv.appendChild(div);
    });
}

async function pesquisarCliente() { // Search Client by ID
    const id = document.getElementById('cliente-id-pesquisar').value;
    const response = await fetch(`http://localhost:8080/cliente/${id}`);
    const cliente = await response.json();
    document.getElementById('cliente-detalhes').textContent = JSON.stringify(cliente);
}

async function atualizarCliente() { // Update Client
    const id = document.getElementById('cliente-id-atualizar').value;
    const nome = document.getElementById('cliente-nome-atualizar').value;

    const response = await fetch(`http://localhost:8080/cliente/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ nome })
    });

    if (response.ok) {
        alert('Cliente atualizado com sucesso!');
    } else {
        alert('Erro ao atualizar cliente.');
    }
}

async function removerCliente() { // Remove Client
    const id = document.getElementById('cliente-id-remover').value;
    const response = await fetch(`http://localhost:8080/cliente/${id}`, { method: 'DELETE' });

    if (response.ok) {
        alert('Cliente removido com sucesso!');
    } else {
        alert('Erro ao remover cliente.');
    }
}
