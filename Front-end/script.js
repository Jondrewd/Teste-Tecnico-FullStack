document.addEventListener("DOMContentLoaded", function () {
    const addContactBtn = document.getElementById("add-contact");
    const contactsContainer = document.getElementById("contacts");
    const MAX_CONTACTS = 4;

    function updateAddButtonState() {
        // Desativa o botão se atingir o limite de 4 contatos
        if (contactsContainer.getElementsByClassName("contact").length >= MAX_CONTACTS) {
            addContactBtn.disabled = true;
        } else {
            addContactBtn.disabled = false;
        }
    }

    function createContactElement() {
        if (contactsContainer.getElementsByClassName("contact").length >= MAX_CONTACTS) return;

        const contactDiv = document.createElement("div");
        contactDiv.classList.add("contact");

        // Select (tipo de contato)
        const select = document.createElement("select");
        select.innerHTML = `
            <option value="Celular">Celular</option>
            <option value="Email">Email</option>
            <option value="Residencial">Residencial</option>
        `;

        // Input para número ou email
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

        // Atualiza placeholder conforme seleção
        select.addEventListener("change", function () {
            input.placeholder = select.value === "Email" ? "Email" : "Número";
        });

        // Evento para remover contato
        removeBtn.addEventListener("click", function () {
            contactsContainer.removeChild(contactDiv);
            updateAddButtonState();
        });

        // Adiciona os elementos na div do contato
        contactDiv.appendChild(select);
        contactDiv.appendChild(input);
        contactDiv.appendChild(obsInput);
        contactDiv.appendChild(removeBtn);

        // Adiciona a div ao container de contatos
        contactsContainer.appendChild(contactDiv);

        updateAddButtonState();
    }

    // Evento para adicionar um novo contato
    addContactBtn.addEventListener("click", createContactElement);
});
