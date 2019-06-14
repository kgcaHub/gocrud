export function Create() {
    var person = {};
    person.id = document.getElementById("txtId").value;
    person.name = document.getElementById("txtName").value;
    var req = new XMLHttpRequest();
    req.open('POST', '/person/create', true);
    req.onreadystatechange = function () {
        if (req.readyState === XMLHttpRequest.DONE) {
            switch (req.status) {
                case 200:
                    Read();
                    break;
            }
        }
    };
    req.send(JSON.stringify(person));
}

var ReadPerson = function (person) {
    return function () {
        document.getElementById("txtId").value = person.id;
        document.getElementById("txtName").value = person.name;
    };
}

export function Read() {
    debugger;
    var req = new XMLHttpRequest();
    req.open('GET', '/person/read', true);
    req.onreadystatechange = function () {
        if (req.readyState === XMLHttpRequest.DONE) {
            debugger;
            switch (req.status) {
                case 200:
                    var _personas = {};
                    if (req.response !== "") {
                        _personas = JSON.parse(req.response);
                    }
                    var tableBody = document.getElementById("tbPersons");
                    while (tableBody.hasChildNodes()) {
                        tableBody.removeChild(tableBody.firstChild);
                    }
                    _personas.forEach(element => {
                        var _newRow = tableBody.insertRow(-1);
                        _newRow.onclick = ReadPerson(element);
                        var _cellId = _newRow.insertCell(0);
                        _cellId.append(element.id);
                        var _cellName = _newRow.insertCell(1);
                        _cellName.append(element.name);
                    });
                    break;
            }
        }
    };
    req.send(null);
}

export function Update() {
    var person = {};
    person.id = document.getElementById("txtId").value;
    person.name = document.getElementById("txtName").value;
    var req = new XMLHttpRequest();
    req.open('POST', '/person/update', true);
    req.onreadystatechange = function () {
        if (req.readyState === XMLHttpRequest.DONE) {
            switch (req.status) {
                case 200:
                    Read();
                    break;
            }
        }
    };
    req.send(JSON.stringify(person));
}

export function Delete() {
    var person = {};
    person.id = document.getElementById("txtId").value;
    person.name = document.getElementById("txtName").value;
    var req = new XMLHttpRequest();
    req.open('POST', '/person/delete', true);
    req.onreadystatechange = function () {
        if (req.readyState === XMLHttpRequest.DONE) {
            switch (req.status) {
                case 200:
                    Read();
                    break;
            }
        }
    };
    req.send(JSON.stringify(person));
}

export default { Create, Read, Update, Delete };