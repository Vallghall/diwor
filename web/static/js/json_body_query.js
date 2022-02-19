function jsonBodyQuery(method, queryUrl, replaceUrl, value) {
    let xhr = new XMLHttpRequest();
    xhr.onload = () => {

        if (xhr.status >= 200 && xhr.status < 300) {
            const response = JSON.parse(xhr.responseText);
            console.log(response);
            if (replaceUrl === 'nil') {return}
            window.location.replace(replaceUrl)
        } else {
            swal(
                xhr.statusText,
                JSON.parse(xhr.responseText).message,
                "warning"
            );
        }
    };

    xhr.open(method, queryUrl);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify(value));
}