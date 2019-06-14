import Person from './components/person.js';

window.goto = function goto(url) {
  var _li = event.target;

  var req = new XMLHttpRequest();
  req.open('GET', url, true);
  req.onreadystatechange = function () {
    if (req.readyState === XMLHttpRequest.DONE) {
      switch (req.status) {
        case 200:
          var _container = document.getElementById("container");
          _container.innerHTML = req.responseText;

          removeAllActives();

          _li.classList.add("active");

          Person.Read();

          break;
      }
    }
  };
  req.send(null);
}

window.onBtnClick = function onBtnClick(action) {

  switch (action) {
    case "create":
      Person.Create();
      break;
    case "read":
      Person.Read();
      break;
    case "update":
      Person.Update();
      break;
    case "delete":
      Person.Delete();
      break;
  }

}

function removeAllActives() {
  var lis = document.getElementById("nav").getElementsByTagName("li");
  [].forEach.call(lis, function (li) {
    li.classList.remove("active");
  });
}