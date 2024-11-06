async function listarVendas() {
    try {
        const response = await fetch('http://localhost:8080/venda');
        if (!response.ok) throw new Error('Erro ao listar vendas');

        const vendas = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = '';

        vendas.forEach(venda => {
            const div = document.createElement('div');
            div.textContent = `ID: ${venda.id}, Cliente: ${venda.cliente}, Valor: ${venda.valor}`;
            output.appendChild(div);
        });
    } catch (error) {
        console.error(error);
        alert('Erro ao listar vendas');
    }
}

async function adicionarVenda() {
    const cliente = prompt('Digite o nome do cliente:');
    const valor = prompt('Digite o valor da venda:');

    try {
        const response = await fetch('http://localhost:8080/venda/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ cliente, valor })
        });
        
        if (!response.ok) throw new Error('Erro ao adicionar venda');
        
        alert('Venda adicionada com sucesso');
        listarVendas();
    } catch (error) {
        console.error(error);
        alert('Erro ao adicionar venda');
    }
}

async function removerVenda() {
    const id = prompt('Digite o ID da venda a ser removida:');
    
    try {
        const response = await fetch(`http://localhost:8080/venda/delete/${id}`, {
            method: 'DELETE'
        });
        
        if (!response.ok) throw new Error('Erro ao remover venda');
        
        alert('Venda removida com sucesso');
        listarVendas();
    } catch (error) {
        console.error(error);
        alert('Erro ao remover venda');
    }
}

async function atualizarVenda() {
    const id = prompt('Digite o ID da venda a ser atualizada:');
    const cliente = prompt('Digite o novo nome do cliente:');
    const valor = prompt('Digite o novo valor da venda:');
    
    try {
        const response = await fetch(`http://localhost:8080/venda/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ cliente, valor })
        });
        
        if (!response.ok) throw new Error('Erro ao atualizar venda');
        
        alert('Venda atualizada com sucesso');
        listarVendas();
    } catch (error) {
        console.error(error);
        alert('Erro ao atualizar venda');
    }
}

async function pesquisarVenda() {
    const id = prompt('Digite o ID da venda a ser pesquisada:');
    
    try {
        const response = await fetch(`http://localhost:8080/venda/${id}`);
        
        if (!response.ok) throw new Error('Venda não encontrada');
        
        const venda = await response.json();
        const output = document.getElementById('data-output');
        output.innerHTML = `ID: ${venda.id}, Cliente: ${venda.cliente}, Valor: ${venda.valor}`;
    } catch (error) {
        console.error(error);
        alert('Erro ao pesquisar venda');
    }
}
