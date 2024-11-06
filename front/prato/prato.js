async function listarpratos() {
    try {
        const response = await fetch('http://localhost:8080/prato');
        if (!response.ok) throw new Error('Erro ao listar pratos');

        const pratos = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = '';

        pratos.forEach(prato => {
            const div = document.createElement('div');
            div.textContent = `ID: ${prato.id}, Nome: ${prato.nome}, Quantidade: ${prato.quantidade}`;
            output.appendChild(div);
        });
    } catch (error) {
        console.error(error);
        alert('Erro ao listar pratos');
    }
}

async function adicionarprato() {
    const nome = prompt('Digite o nome do prato:');
    const quantidade = prompt('Digite a quantidade do prato:');

    try {
        const response = await fetch('http://localhost:8080/prato/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ nome, quantidade })
        });
        
        if (!response.ok) throw new Error('Erro ao adicionar prato');
        
        alert('prato adicionado com sucesso');
        listarpratos();
    } catch (error) {
        console.error(error);
        alert('Erro ao adicionar prato');
    }
}

async function removerprato() {
    const id = prompt('Digite o ID do prato a ser removido:');
    
    try {
        const response = await fetch(`http://localhost:8080/prato/delete/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) throw new Error('Erro ao remover prato');
        
        alert('prato removido com sucesso');
        listarpratos();
    } catch (error) {
        console.error(error);
        alert('Erro ao remover prato');
    }
}

async function atualizarprato() {
    const id = prompt('Digite o ID do prato a ser atualizado:');
    const nome = prompt('Digite o novo nome do Prato:');
    const quantidade = prompt('Digite a nova quantidade do Prato:');
    
    try {
        const response = await fetch(`http://localhost:8080/prato/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ nome, quantidade })
        });
        
        if (!response.ok) throw new Error('Erro ao atualizar Prato');
        
        alert('Prato atualizado com sucesso');
        listarpratos();
    } catch (error) {
        console.error(error);
        alert('Erro ao atualizar Prato');
    }
}

async function pesquisarprato() {
    const id = prompt('Digite o ID do Prato a ser pesquisado:');
    
    try {
        const response = await fetch(`http://localhost:8080/prato/${id}`);
        
        if (!response.ok) throw new Error('Prato não encontrado');
        
        const prato = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = `ID: ${prato.id}, Nome: ${prato.nome}, Quantidade: ${prato.quantidade}`;
    } catch (error) {
        console.error(error);
        alert('Erro ao pesquisar prato');
    }
}
