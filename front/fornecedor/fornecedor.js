async function listarFornecedores() { // List Suppliers
    try {
        const response = await fetch('http://localhost:8080/fornecedor');
        if (!response.ok) throw new Error('Erro ao listar fornecedores');

        const fornecedores = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = '';

        fornecedores.forEach(fornecedor => {
            const div = document.createElement('div');
            div.textContent = `ID: ${fornecedor.id}, Nome: ${fornecedor.nome}, Contato: ${fornecedor.contato}`;
            output.appendChild(div);
        });
    } catch (error) {
        console.error(error);
        alert('Erro ao listar fornecedores');
    }
}

async function adicionarFornecedor() { // Add Suppliers
    const nome = prompt('Digite o nome do fornecedor:');
    const contato = prompt('Digite o contato do fornecedor:');

    try {
        const response = await fetch('http://localhost:8080/fornecedor/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ nome, contato })
        });
        
        if (!response.ok) throw new Error('Erro ao adicionar fornecedor');
        
        alert('Fornecedor adicionado com sucesso');
        listarFornecedores();
    } catch (error) {
        console.error(error);
        alert('Erro ao adicionar fornecedor');
    }
}

async function removerFornecedor() { // Remove Suppliers
    const id = prompt('Digite o ID do fornecedor a ser removido:');
    
    try {
        const response = await fetch(`http://localhost:8080/fornecedor/delete/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) throw new Error('Erro ao remover fornecedor');
        
        alert('Fornecedor removido com sucesso');
        listarFornecedores();
    } catch (error) {
        console.error(error);
        alert('Erro ao remover fornecedor');
    }
}

async function atualizarFornecedor() { // Update Suppliers
    const id = prompt('Digite o ID do fornecedor a ser atualizado:');
    const nome = prompt('Digite o novo nome do fornecedor:');
    const contato = prompt('Digite o novo contato do fornecedor:');
    
    try {
        const response = await fetch(`http://localhost:8080/fornecedor/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ nome, contato })
        });
        
        if (!response.ok) throw new Error('Erro ao atualizar fornecedor');
        
        alert('Fornecedor atualizado com sucesso');
        listarFornecedores();
    } catch (error) {
        console.error(error);
        alert('Erro ao atualizar fornecedor');
    }
}

async function pesquisarFornecedor() { // Search Suppliers
    const id = prompt('Digite o ID do fornecedor a ser pesquisado:');
    
    try {
        const response = await fetch(`http://localhost:8080/fornecedor/${id}`);
        
        if (!response.ok) throw new Error('Fornecedor não encontrado');
        
        const fornecedor = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = `ID: ${fornecedor.id}, Nome: ${fornecedor.nome}, Contato: ${fornecedor.contato}`;
    } catch (error) {
        console.error(error);
        alert('Erro ao pesquisar fornecedor');
    }
}
