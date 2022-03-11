let form = document.getElementById("hash-form");

const numPattern = /^[0-9]+$/g;

form.addEventListener('submit', (e) => {
    e.preventDefault();

    const data = document.getElementsByClassName("select-alg");
    const formData = new FormData(e.target);
    const value = Object.fromEntries(formData.entries());

    Object.entries(value).forEach(([_, val]) => {
        if (val === "") {
            swal({
                title : "Неверные значения полей",
                text : "Не должно быть пустых полей",
                icon :"warning",
            });
            throw new Error('Bad request');
        }
        if (!val.match(numPattern)) {
            swal({
                title : "Неверные значения полей",
                text : 'Значения полей "От", "До", "Шаг" и "Число замеров" должны быть числами',
                icon :"warning",
            });
            throw new Error('Bad request');
        }
    })

    let query = {
        "from" : +value.from,
        "to" : +value.to,
        "step" : +value.step,
        "num_measurements" : +value.num,
        "algorithms": []
    }

    for (const datum of data) {
        if (!hashAlgorithms.includes(datum.value)) {
            swal({
                title : "Неверные значения полей",
                text : 'Необходимо выбрать корректный алгоритм',
                icon :"warning",
            });
            return;
        }
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
