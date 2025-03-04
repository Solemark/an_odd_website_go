/**
 * @typedef Client
 * @type {object}
 * @property {number} Id - client id
 * @property {string} FirstName - client first name
 * @property {string} LastName - client last name
 * @property {string} EmailAddress - client email address
 * @property {boolean} Visible - is the client visible
 */

/**
 * On load function
 * @returns {void}
 */
const getClients = () => {
    fetch("/data/clients")
        .then(async (res) => {
            /** @type {Client[]} */
            const data = await res.json()
            updateTable(data)
        })
        .catch((err) => {
            console.log(err)
        })
}

/**
 * find client id for edit and remove client
 * @param {number} id - the client id
 * @param {string} option - the dialog that is being set
 */
const displayDialog = (id, option) => {
    document.getElementById(`${option}-id`).value = id
    if (option == "remove") {
        removeClient.showModal()
    }
    if (option == "edit") {
        editClient.showModal()
    }
}

/**
 * Updates the client table
 * @param {Client[]} data - array of client data
 * @returns {void}
 */
const updateTable = (data) => {
    let tbody = document.getElementById("client-table")
    data.forEach(client => {
        let i = client.Id
        let row = document.createElement("tr")
        let id = document.createElement("td")
        let fn = document.createElement("td")
        let ln = document.createElement("td")
        let ea = document.createElement("td")
        let bc = document.createElement("td")

        let eb = document.createElement("button")
        eb.setAttribute("type", "button")
        eb.setAttribute("class", "yellowButton")
        eb.setAttribute("onclick", `displayDialog(${i}, "edit")`)

        let rb = document.createElement("button")
        rb.setAttribute("type", "button")
        rb.setAttribute("class", "redButton")
        rb.setAttribute("onclick", `displayDialog(${i}, "remove")`)

        id.innerHTML = i
        fn.innerHTML = client.FirstName
        ln.innerHTML = client.LastName
        ea.innerHTML = client.EmailAddress
        eb.innerHTML = "Edit"
        rb.innerHTML = "Remove"
        bc.append(eb, rb)
        row.append(id, fn, ln, ea, bc)
        tbody.append(row)
    })
}

/**
 * Clear the modal inputs
 * @returns {void}
 */
const clearModal = () => {
    newClient.close()
    removeClient.close()
    editClient.close()
    document.getElementById("new-fn").value = ""
    document.getElementById("new-ln").value = ""
    document.getElementById("new-ea").value = ""
    document.getElementById("edit-id").value = ""
    document.getElementById("edit-fn").value = ""
    document.getElementById("edit-ln").value = ""
    document.getElementById("edit-ea").value = ""
    document.getElementById("remove-id").value = ""
}
