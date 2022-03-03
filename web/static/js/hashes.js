let form = document.getElementById("hash-form");

form.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = document.getElementsByClassName("select-alg");
    const formData = new FormData(e.target);
    const value = Object.fromEntries(formData.entries());

    let query = {
        "from" : +value.from,
        "to" : +value.to,
        "step" : +value.step,
        "num_measurements" : +value.num,
        "algorithms": []
    }

    for (const datum of data) {
        query.algorithms.push(datum.value)
    }
    console.table(query)

    jsonBodyQuery('POST', '/api/experiment/start-hash-experiment','nil', query)
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
