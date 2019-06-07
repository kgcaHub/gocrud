function ReadPerson() {
    debugger;
    var req = new XMLHttpRequest();
    req.open('GET', '/person/read', true);
    req.onreadystatechange = function () {
        if (req.readyState === XMLHttpRequest.DONE) {
            debugger;
            var _jsonResponse = {};

            if (req.response !== "") {
                _jsonResponse = JSON.parse(req.response);
            }
        }
    };
    req.send(null);
}