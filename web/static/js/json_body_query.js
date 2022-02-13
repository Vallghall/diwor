function jsonBodyQuery(method, queryUrl, replaceUrl, value) {
    let xhr = new XMLHttpRequest();
    xhr.onload = () => {

        if (xhr.status >= 200 && xhr.status < 300) {
            const response = JSON.parse(xhr.responseText);
            console.log(response);

            window.location.replace(replaceUrl)
        } else {
            const response = {
                statusCode: xhr.statusText,
                responseInfo: JSON.parse(xhr.responseText)
            }
            alert(response);
        }
    };

    xhr.open(method, queryUrl);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify(value));
}