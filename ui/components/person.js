export function Create() {
    var person = {};
    person.id = document.getElementById("txtId").value;
    person.name = document.getElementById("txtName").value;
    var req = new XMLHttpRequest();
    req.open('POST', '/person/create', true);
    req.onreadystatechange = function () {
        if (req.readyState === XMLHttpRequest.DONE) {
            debugger;
            switch (req.status) {
                case 200:
                    debugger;
                  break;
              }
        }
    };
    req.send(JSON.stringify(person));
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
                    _personas.forEach(element => {
                        var _newRow = tableBody.insertRow(-1);
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

}

export function Delete() {

}

export default { Create, Read, Update, Delete };