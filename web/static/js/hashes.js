let form = document.getElementById("hash-form");

form.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = document.getElementsByClassName("select-alg");
    let query = {"algorithms" : []};

    for (const datum of data) {
        query.algorithms.push(datum.value)
    }

    jsonBodyQuery('POST', '/api/experiment/start-hash-experiment','/api/experiment/hash-results', query)
})