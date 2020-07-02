function setFormMessage(message, className = "error-text") {
    const formMessage = document.getElementsByClassName("form-message")[0];
    while (formMessage.firstChild) formMessage.removeChild(formMessage.firstChild);
    formMessage.insertAdjacentText("afterbegin", message);
    formMessage.className = "form-message " + className;
}

function getQueryVariable(name) {
    let query = location.search.substring(1).split('&');
    for (let pair of query.map(pair => pair.split("="))) {
        if (decodeURIComponent(pair[0]) === name) {
            return decodeURIComponent(pair[1]);
        }
    }
    return null;
}

function POST(path, body) {
    return fetch(path, {
        method: "POST",
        headers: {
            'Accept': 'application/json, text/plain, */*',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    });
}

export { getQueryVariable, POST, setFormMessage };