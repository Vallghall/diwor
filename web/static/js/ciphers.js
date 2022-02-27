let form = document.getElementById("cipher-form");

form.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = document.getElementsByClassName("select-alg");
    let query = {"algorithms" : []};

    for (const datum of data) {
        query.algorithms.push(datum.value)
    }

    jsonBodyQuery('POST', '/api/experiment/start-cipher-experiment','nil', query)
    swal({
        title : "Данные взяты на обработку",
        text : "Данные были приняты на обработку сервером. Вы можете ознакомиться с результатами в профиле пользователя, когда они будут готовы",
        icon :"success",
        buttons : {
            cancel: "OK",
            profile: "Профиль"
        }
    })
        .then((value) => {
            if (value === "profile") {
                window.location.href="/api/profile"
            }
        })
})