document.addEventListener("DOMContentLoaded", function () {
    /**
     * Gerenciamento de contatos - permite adicionar até 4 contatos.
     */
    const addContactBtn = document.getElementById("add-contact");
    const contactsContainer = document.getElementById("contacts");
    const MAX_CONTACTS = 4;

    function updateAddButtonState() {
        // Desativa o botão se atingir o limite máximo de contatos
        addContactBtn.disabled = contactsContainer.getElementsByClassName("contact").length >= MAX_CONTACTS;
    }

    function createContactElement() {
        if (contactsContainer.getElementsByClassName("contact").length >= MAX_CONTACTS) return;

        const contactDiv = document.createElement("div");
        contactDiv.classList.add("contact");

        // Criação do select para tipo de contato
        const select = document.createElement("select");
        select.innerHTML = `
            <option value="Celular">Celular</option>
            <option value="Email">Email</option>
            <option value="Residencial">Residencial</option>
        `;

        // Input para número ou e-mail
        const input = document.createElement("input");
        input.type = "text";
        input.placeholder = "Número"; // Placeholder inicial

        // Input para observações
        const obsInput = document.createElement("input");
        obsInput.type = "text";
        obsInput.placeholder = "Observações";

        // Botão para remover contato
        const removeBtn = document.createElement("button");
        removeBtn.textContent = "-";
        removeBtn.classList.add("remove-contact");

        // Atualiza placeholder do input baseado na seleção
        select.addEventListener("change", function () {
            input.placeholder = select.value === "Email" ? "Email" : "Número";
        });

        // Evento para remover contato
        removeBtn.addEventListener("click", function () {
            contactsContainer.removeChild(contactDiv);
            updateAddButtonState();
        });

        // Adiciona os elementos à div do contato
        contactDiv.append(select, input, obsInput, removeBtn);
        contactsContainer.appendChild(contactDiv);
        updateAddButtonState();
    }

    addContactBtn.addEventListener("click", createContactElement);
});

/**
 * Busca de endereço via CEP utilizando a API ViaCEP
 */
document.getElementById('buscar-cep').addEventListener('click', function () {
    const cep = document.getElementById('cep').value.replace(/\D/g, ''); // Remove caracteres não numéricos

    if (cep.length !== 8) {
        alert('CEP inválido. O CEP deve conter 8 dígitos.');
        return;
    }

    fetch(`https://viacep.com.br/ws/${cep}/json/`)
        .then(response => response.json())
        .then(data => {
            if (!data.erro) {
                // Preenche os campos do formulário
                document.getElementById('logradouro').value = data.logradouro || '';
                document.getElementById('bairro').value = data.bairro || '';
                document.getElementById('cidade').value = data.localidade || '';
                document.getElementById('estado').value = data.uf || '';
                document.getElementById('pais').value = 'Brasil';
                
                // Busca latitude e longitude no OpenStreetMap
                buscarLatitudeLongitude(`${data.logradouro}, ${data.bairro}, ${data.localidade}, ${data.uf}, Brasil`);
            } else {
                alert('CEP não encontrado. Verifique o CEP e tente novamente.');
            }
        })
        .catch(error => console.error('Erro ao buscar CEP:', error));
});

/**
 * Busca latitude e longitude no OpenStreetMap
 */
function buscarLatitudeLongitude(endereco) {
    if (!endereco.trim()) {
        alert('Endereço não encontrado. Verifique o CEP e tente novamente.');
        return;
    }

    fetch(`https://nominatim.openstreetmap.org/search?format=json&q=${encodeURIComponent(endereco)}`)
        .then(response => response.json())
        .then(data => {
            if (data.length > 0) {
                document.getElementById('latitude').value = data[0].lat;
                document.getElementById('longitude').value = data[0].lon;
                atualizarMapa(data[0].lat, data[0].lon);
            } else {
                alert('Endereço não encontrado no OpenStreetMap.');
            }
        })
        .catch(error => console.error('Erro ao buscar endereço:', error));
}

/**
 * Atualiza o mapa com a localização encontrada
 */
function atualizarMapa(lat, lon) {
    const map = L.map('map').setView([lat, lon], 15);
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
    }).addTo(map);

    L.marker([lat, lon]).addTo(map)
        .bindPopup('Localização do endereço encontrado.')
        .openPopup();
}

/**
 * Envio de pedido para o backend
 */
document.getElementById('order-form').addEventListener('submit', function (event) {
    event.preventDefault();

    const pedido = {
        client_name: document.getElementById('nome-cliente').value,
        cpf: document.getElementById('cpf-cliente').value,
        test_name: document.getElementById('teste-tecnico').value,
        weight: parseFloat(document.getElementById('peso').value),  // Converte para número
        address: `${document.getElementById('logradouro').value}, ${document.getElementById('numero').value}, ${document.getElementById('bairro').value}, ${document.getElementById('complemento').value}, ${document.getElementById('cidade').value} - ${document.getElementById('estado').value}, ${document.getElementById('pais').value}`,
        latitude: parseFloat(document.getElementById('latitude').value),  // Converte para número
        longitude: parseFloat(document.getElementById('longitude').value),  // Converte para número
        order_status: "Pending"
    };
    console.log(pedido)
    fetch('http://localhost:8080/api/v1/deliveries', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(pedido),
    })
    .then(response => response.json())
    .then(() => alert('Pedido salvo com sucesso!'))
    .catch(error => console.error('Erro ao salvar pedido:', error));
});

/**
 * Controle de exibição do mapa
 */
document.getElementById('ver-mapa').addEventListener('click', function () {
    document.getElementById('map-container').style.display = 'block';
});

document.getElementById('fechar-mapa').addEventListener('click', function () {
    document.getElementById('map-container').style.display = 'none';
});