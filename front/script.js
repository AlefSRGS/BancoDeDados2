// Manipulação do envio do formulário
document.getElementById('data-form').addEventListener('submit', async function(event) {
    event.preventDefault(); // Evita o envio padrão do formulário

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;

    try {
        // Enviar dados para o endpoint da API
        const response = await fetch('http://localhost:8080/submit', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name: name, email: email })
        });

        const result = await response.json();
        alert('Data submitted successfully: ' + result.message);
    } catch (error) {
        console.error('Error:', error);
        alert('Failed to submit data.');
    }
});

// Função para buscar dados da API e mostrá-los na página
async function fetchData() {
    try {
        const response = await fetch('http://localhost:8080/data');
        const data = await response.json();

        // Atualizar o container com os dados recebidos
        const dataOutput = document.getElementById('data-output');
        dataOutput.innerHTML = ''; // Limpa o conteúdo atual

        data.forEach(item => {
            const div = document.createElement('div');
            div.textContent = `Name: ${item.name}, Email: ${item.email}`;
            dataOutput.appendChild(div);
        });
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}

// Buscar dados ao carregar a página
window.onload = fetchData;
