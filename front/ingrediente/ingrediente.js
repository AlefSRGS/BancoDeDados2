async function listarIngredientes() {
    try {
        const response = await fetch('http://localhost:8080/ingrediente');
        if (!response.ok) throw new Error('Erro ao listar ingredientes');

        const ingredientes = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = '';

        ingredientes.forEach(ingrediente => {
            const div = document.createElement('div');
            div.textContent = `ID: ${ingrediente.id}, Nome: ${ingrediente.nome}, Quantidade: ${ingrediente.quantidade}`;
            output.appendChild(div);
        });
    } catch (error) {
        console.error(error);
        alert('Erro ao listar ingredientes');
    }
}

async function adicionarIngrediente() {
    const nome = prompt('Digite o nome do ingrediente:');
    const quantidade = prompt('Digite a quantidade do ingrediente:');

    try {
        const response = await fetch('http://localhost:8080/ingrediente/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ nome, quantidade })
        });
        
        if (!response.ok) throw new Error('Erro ao adicionar ingrediente');
        
        alert('Ingrediente adicionado com sucesso');
        listarIngredientes();
    } catch (error) {
        console.error(error);
        alert('Erro ao adicionar ingrediente');
    }
}

async function removerIngrediente() {
    const id = prompt('Digite o ID do ingrediente a ser removido:');
    
    try {
        const response = await fetch(`http://localhost:8080/ingrediente/delete/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) throw new Error('Erro ao remover ingrediente');
        
        alert('Ingrediente removido com sucesso');
        listarIngredientes();
    } catch (error) {
        console.error(error);
        alert('Erro ao remover ingrediente');
    }
}

async function atualizarIngrediente() {
    const id = prompt('Digite o ID do ingrediente a ser atualizado:');
    const nome = prompt('Digite o novo nome do ingrediente:');
    const quantidade = prompt('Digite a nova quantidade do ingrediente:');
    
    try {
        const response = await fetch(`http://localhost:8080/ingrediente/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ nome, quantidade })
        });
        
        if (!response.ok) throw new Error('Erro ao atualizar ingrediente');
        
        alert('Ingrediente atualizado com sucesso');
        listarIngredientes();
    } catch (error) {
        console.error(error);
        alert('Erro ao atualizar ingrediente');
    }
}

async function pesquisarIngrediente() {
    const id = prompt('Digite o ID do ingrediente a ser pesquisado:');
    
    try {
        const response = await fetch(`http://localhost:8080/ingrediente/${id}`);
        
        if (!response.ok) throw new Error('Ingrediente não encontrado');
        
        const ingrediente = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = `ID: ${ingrediente.id}, Nome: ${ingrediente.nome}, Quantidade: ${ingrediente.quantidade}`;
    } catch (error) {
        console.error(error);
        alert('Erro ao pesquisar ingrediente');
    }
}
