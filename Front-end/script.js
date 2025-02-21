document.addEventListener('DOMContentLoaded', function () {
    // Inicializa funções ao carregar a página
    carregarClientes();
    updateClientCount();
    setupAddressListener();
    setupOrderFormListener();
    setupMapControls();
    carregarPedidos();
});

function setupAddressListener() {
    document.getElementById('buscar-endereco').addEventListener('click', function () {
        const endereco = document.getElementById('endereco').value.trim();

        if (!endereco) {
            alert('Por favor, insira um endereço válido.');
            return;
        }

        buscarLatitudeLongitude(endereco);
    });
}

function buscarLatitudeLongitude(endereco) {
    fetch(`https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(endereco)}`)
        .then(response => response.json())
        .then(data => {
            if (data.length > 0) {
                const primeiroResultado = data[0];
                preencherEndereco(primeiroResultado);
                atualizarMapa(primeiroResultado.lat, primeiroResultado.lon);
            } else {
                alert('Endereço não encontrado no OpenStreetMap.');
            }
        })
        .catch(error => console.error('Erro ao buscar endereço:', error));
}

function preencherEndereco(data) {
    // Extrai os componentes do endereço
    const enderecoCompleto = data.display_name.split(', ');
    document.getElementById('logradouro').value = enderecoCompleto[0] || '';
    document.getElementById('bairro').value = enderecoCompleto[1] || '';
    document.getElementById('cidade').value = enderecoCompleto[2] || '';
    document.getElementById('estado').value = enderecoCompleto[3] || '';
    document.getElementById('pais').value = enderecoCompleto[enderecoCompleto.length - 1] || '';
    document.getElementById('latitude').value = data.lat;
    document.getElementById('longitude').value = data.lon;
}

function atualizarMapa(lat, lon) {
    const map = L.map('map').setView([lat, lon], 15);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map);

    L.marker([lat, lon]).addTo(map)
        .bindPopup('Localização do endereço encontrado.')
        .openPopup();
}

function setupOrderFormListener() {
    document.getElementById('order-form').addEventListener('submit', function (event) {
        event.preventDefault();

        const pedido = {
            client_name: document.getElementById('nome-cliente').value,
            client_cpf: document.getElementById('cpf-cliente').value,
            test_name: document.getElementById('teste-tecnico').value,
            weight: parseFloat(document.getElementById('peso').value),  // Converte para número
            logradouro: document.getElementById('logradouro').value,
            numero: document.getElementById('numero').value,
            bairro: document.getElementById('bairro').value,
            complemento: document.getElementById('complemento').value,
            cidade: document.getElementById('cidade').value,
            estado: document.getElementById('estado').value,
            pais: document.getElementById('pais').value,
            latitude: parseFloat(document.getElementById('latitude').value),  
            longitude: parseFloat(document.getElementById('longitude').value),  
            order_status: "Pendente"
        };

        console.log("Enviando pedido:", pedido);

        fetch('http://localhost:8080/api/v1/deliveries', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(pedido),
        })
        .then(response => response.json())
        .then(() => alert('Pedido salvo com sucesso!'))
        .catch(error => console.error('Erro ao salvar pedido:', error));
    });
}

function setupMapControls() {
    document.getElementById('ver-mapa').addEventListener('click', function () {
        document.getElementById('map-container').style.display = 'block';
    });

    document.getElementById('fechar-mapa').addEventListener('click', function () {
        document.getElementById('map-container').style.display = 'none';
    });
}

async function updateClientCount() {
    try {
        const response = await fetch('http://localhost:8080/api/v1/clients/count');
        if (!response.ok) {
            throw new Error('Erro ao buscar a contagem de clientes');
        }
        const data = await response.json(); // Converte a resposta para um objeto JavaScript

        // Acessa a propriedade correta dentro do objeto retornado
        const count = data.total_clients; 

        // Seleciona o elemento <p> que contém o número de clientes
        const countElement = document.querySelector('.span-header-2 .p2');
        if (countElement) {
            countElement.textContent = `${count} cliente(s)`;
        }
    } catch (error) {
        console.error('Erro:', error);
    }
}

document.addEventListener("DOMContentLoaded", function () {
    const cpfInput = document.getElementById("cpf");
    const cnpjInput = document.getElementById("cnpj");
    const phoneInput = document.getElementById("phone");

    // Aplica máscara ao CPF (XXX.XXX.XXX-XX)
    cpfInput.addEventListener("input", function () {
        let value = cpfInput.value.replace(/\D/g, ""); // Remove tudo que não for número
        if (value.length > 11) value = value.slice(0, 11);
        value = value.replace(/(\d{3})(\d)/, "$1.$2");
        value = value.replace(/(\d{3})(\d)/, "$1.$2");
        value = value.replace(/(\d{3})(\d{1,2})$/, "$1-$2");
        cpfInput.value = value;
    });

    // Aplica máscara ao CNPJ (XX.XXX.XXX/XXXX-XX)
    cnpjInput.addEventListener("input", function () {
        let value = cnpjInput.value.replace(/\D/g, "");
        if (value.length > 14) value = value.slice(0, 14);
        value = value.replace(/^(\d{2})(\d)/, "$1.$2");
        value = value.replace(/^(\d{2})\.(\d{3})(\d)/, "$1.$2.$3");
        value = value.replace(/\.(\d{3})(\d)/, ".$1/$2");
        value = value.replace(/(\d{4})(\d{1,2})$/, "$1-$2");
        cnpjInput.value = value;
    });

    // Aplica máscara ao telefone (XX XXXXX-XXXX ou XX XXXX-XXXX)
    phoneInput.addEventListener("input", function () {
        let value = phoneInput.value.replace(/\D/g, "");
        if (value.length > 11) value = value.slice(0, 11);
        if (value.length === 11) {
            value = value.replace(/(\d{2})(\d{5})(\d{4})/, "($1) $2-$3");
        } else if (value.length >= 10) {
            value = value.replace(/(\d{2})(\d{4})(\d{4})/, "($1) $2-$3");
        }
        phoneInput.value = value;
    });
});

// Função para cadastrar cliente
function cadastrarCliente() {
    const name = document.getElementById("name").value.trim();
    const cpf = document.getElementById("cpf").value.trim();
    const cnpj = document.getElementById("cnpj").value.trim();
    const birthdate = document.getElementById("birthdate").value;
    const email = document.getElementById("email").value.trim();
    const phone = document.getElementById("phone").value.trim();

    // Validação básica
    if (!name) {
        alert("O nome é obrigatório!");
        return;
    }
    if (!email && !phone) {
        alert("Informe pelo menos um contato (E-mail ou Telefone).");
        return;
    }

    // Formata os dados para JSON
    const clientData = {
        name,
        cpf: cpf,
        cnpj: cnpj,
        birth_date: birthdate,
        email: email,
        deliveries: null,
        phone: phone,
    };

    // Simula envio para API (substituir com fetch real)
    fetch("http://localhost:8080/api/v1/clients", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(clientData),
    })
        .then(response => response.json())
        .then(data => alert("Cliente cadastrado com sucesso!"))
        .catch(error => console.error("Erro ao cadastrar cliente:", error));
}
function carregarPedidos() {
    fetch('http://localhost:8080/api/v1/deliveries')
        .then(response => response.json())
        .then(data => {
            if (data.length > 0) {
                exibirPedidos(data);
            } else {
                document.getElementById('pedidos').innerHTML = '<p>Nenhum pedido encontrado.</p>';
            }
        })
        .catch(error => {
            console.error('Erro ao buscar pedidos:', error);
            document.getElementById('pedidos').innerHTML = '<p>Erro ao carregar pedidos. Tente novamente mais tarde.</p>';
        });
}function exibirPedidos(pedidos) {
    const sectionPedidos = document.getElementById('pedidos');
    sectionPedidos.innerHTML = ''; // Limpa o conteúdo anterior

    pedidos.forEach(pedido => {
        // Monta o endereço com base nos novos campos
        const address = `${pedido.logradouro}, ${pedido.numero}, ${pedido.bairro}, ${pedido.complemento}, ${pedido.cidade} - ${pedido.estado}, ${pedido.pais}`;

        const pedidoCard = document.createElement('div');
        pedidoCard.className = 'pedido-card';

        pedidoCard.innerHTML = `
            <div class="pedido-info">
                <h3>${pedido.client_name}</h3>
                <p><strong>CPF:</strong> ${pedido.client_cpf}</p>
                <p><strong>Teste Técnico:</strong> ${pedido.test_name}</p>
                <p><strong>Peso:</strong> ${pedido.weight} kg</p>
                <p><strong>Endereço:</strong> ${address}</p>
                <p><strong>Status:</strong> ${pedido.order_status}</p>
            </div>
            <div class="pedido-acoes">
                <button onclick="excluirPedido(${pedido.id})">Excluir</button>
                <button onclick="verMapa(${pedido.id}, ${pedido.latitude}, ${pedido.longitude})">Ver no mapa</button>
            </div>
            <div id="mapa-${pedido.id}" class="mapa-pedido"></div>
        `;

        sectionPedidos.appendChild(pedidoCard);
    });
}
function verMapa(id, latitude, longitude) {
    // Seleciona o contêiner do mapa
    const mapaContainer = document.getElementById(`mapa-${id}`);

    // Alterna a visibilidade do mapa
    if (mapaContainer.style.display === "none" || !mapaContainer.style.display) {
        mapaContainer.style.display = "block"; // Exibe o mapa

        // Inicializa o mapa
        const map = L.map(mapaContainer).setView([latitude, longitude], 15);

        // Adiciona a camada de tiles do OpenStreetMap
        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);

        // Adiciona um marcador no mapa
        L.marker([latitude, longitude]).addTo(map)
            .bindPopup('Localização da entrega.')
            .openPopup();
    } else {
        mapaContainer.style.display = "none"; // Oculta o mapa
    }
}
function excluirPedido(id) {
    if (confirm('Tem certeza que deseja excluir este pedido?')) {
        fetch(`http://localhost:8080/api/v1/deliveries/${id}`, {
            method: 'DELETE',
        })
            .then(response => {
                if (response.ok) {
                    alert('Pedido excluído com sucesso!');
                    carregarPedidos(); // Recarrega a lista de pedidos
                } else {
                    alert('Erro ao excluir pedido.');
                }
            })
            .catch(error => console.error('Erro ao excluir pedido:', error));
    }
}
function pesquisarPedidos() {
    const criterio = document.getElementById('criterio-pesquisa').value; // Pega o critério selecionado (city, name ou cpf)
    const valor = document.getElementById('valor-pesquisa').value.trim(); // Pega o valor digitado

    if (!valor) {
        alert('Por favor, digite um valor para pesquisar.');
        return;
    }

    if (criterio == "city") {
        // Se o critério for "city", usa o endpoint de deliveries
        url = `http://localhost:8080/api/v1/deliveries/city/${encodeURIComponent(valor)}`;
    } else {
        // Para outros critérios (name ou cpf), usa o endpoint de clients
        url = `http://localhost:8080/api/v1/clients/${criterio}/${encodeURIComponent(valor)}`;
    };

    fetch(url)
        .then(response => response.json())
        .then(data => {
            if (data.length > 0) {
                exibirPedidos(data); // Exibe os pedidos encontrados
            } else {
                document.getElementById('pedidos').innerHTML = '<p>Nenhum pedido encontrado.</p>';
            }
        })
        .catch(error => {
            console.error('Erro ao buscar pedidos:', error);
            document.getElementById('pedidos').innerHTML = '<p>Erro ao carregar pedidos. Tente novamente mais tarde.</p>';
        });
}
function abrirFormularioPedido() {
    // Oculta a seção de pedidos
    document.getElementById('pedidos').style.display = 'none';

    // Exibe o formulário de cadastro de pedido
    document.querySelector('.cadastrar-pedido').style.display = 'block';
}

function fecharFormularioPedido() {
    // Oculta o formulário de cadastro de pedido
    document.querySelector('.cadastrar-pedido').style.display = 'none';

    // Exibe a seção de pedidos
    document.getElementById('pedidos').style.display = 'block';
}
// Função para carregar todos os clientes
function carregarClientes() {
    fetch('http://localhost:8080/api/v1/clients')
        .then(response => response.json())
        .then(data => {
            if (data.length > 0) {
                exibirClientes(data);
            } else {
                document.getElementById('clientes').innerHTML = '<p>Nenhum cliente encontrado.</p>';
            }
        })
        .catch(error => {
            console.error('Erro ao buscar clientes:', error);
            document.getElementById('clientes').innerHTML = '<p>Erro ao carregar clientes. Tente novamente mais tarde.</p>';
        });
}



function pesquisarClientes() {
    const criterio = document.getElementById('criterio-pesquisa').value; // Pega o critério selecionado (name ou cpf)
    const valor = document.getElementById('valor-pesquisa').value.trim(); // Pega o valor digitado

    if (!valor) {
        alert('Por favor, digite um valor para pesquisar.');
        return;
    }

    // Monta a URL da API com base no critério e no valor
    const url = `http://localhost:8080/api/v1/clients/${criterio}/${encodeURIComponent(valor)}`;

    console.log('URL da pesquisa:', url); // Log para depuração

    fetch(url)
        .then(response => response.json())
        .then(data => {
            console.log('Dados retornados pela API:', data); // Log para depuração

            // Verifica se os dados são válidos
            if (data && (Array.isArray(data) ? data.length > 0 : Object.keys(data).length > 0)) {
                exibirClientes(Array.isArray(data) ? data : [data]); // Garante que os dados sejam um array
            } else {
                document.getElementById('clientes').innerHTML = '<p>Nenhum cliente encontrado.</p>';
            }
        })
        .catch(error => {
            console.error('Erro ao buscar clientes:', error);
            document.getElementById('clientes').innerHTML = '<p>Erro ao carregar clientes. Tente novamente mais tarde.</p>';
        });
}
// Função para exibir os clientes na lista
function exibirClientes(clientes) {
    const sectionClientes = document.getElementById('clientes');
    sectionClientes.innerHTML = ''; // Limpa o conteúdo anterior

    clientes.forEach(cliente => {
        const clienteCard = document.createElement('div');
        clienteCard.className = 'cliente-card';

        clienteCard.innerHTML = `
            <div class="cliente-info">
                <h3>${cliente.name}</h3>
                <p><strong>CPF:</strong> ${cliente.cpf || 'Não informado'}</p>
                <p><strong>CNPJ:</strong> ${cliente.cnpj || 'Não informado'}</p>
                <p><strong>Email:</strong> ${cliente.email}</p>
                <p><strong>Telefone:</strong> ${cliente.phone}</p>
            </div>
            <div class="cliente-acoes">
                <button onclick="excluirCliente(${cliente.id})">Excluir</button>
            </div>
        `;

        sectionClientes.appendChild(clienteCard);
    });
}
// Função para abrir o formulário de cadastro de cliente
function abrirFormularioCliente() {
    // Oculta a lista de clientes
    document.getElementById('clientes').style.display = 'none';

    // Exibe o formulário de cadastro de cliente
    document.querySelector('.cadastrar-cliente').style.display = 'block';
}

// Função para fechar o formulário de cadastro de cliente
function fecharFormularioCliente() {
    // Oculta o formulário de cadastro de cliente
    document.querySelector('.cadastrar-cliente').style.display = 'none';

    // Exibe a lista de clientes
    document.getElementById('clientes').style.display = 'block';
}
