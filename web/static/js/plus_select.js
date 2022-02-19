const hashAlgorithms = ["Streebog-256","Streebog-512","SHA-224","SHA-256","SHA-384","SHA-512","RIPEMD-128","RIPEMD-160","RIPEMD-256","RIPEMD-320","MD5"];
let numOfSelects = 1;

function AddSelect() {
    if (numOfSelects === 6) {
        alert("Maximum is 6 at once");
        return;
    }
    numOfSelects++

    if (numOfSelects > 1) {
        let minus = document.getElementById("-");
        minus.disabled = false;
    }

    let select = document.createElement("select");
    select.id = `alg-${numOfSelects}`;
    select.className = "select-alg";

    let defaultOption = document.createElement("option");
    defaultOption.disabled = true;
    defaultOption.selected = true;
    defaultOption.textContent = "Выберите алгоритм хеширования";
    select.appendChild(defaultOption);

    for (const alg of hashAlgorithms) {
        let option = document.createElement("option");
        option.value = alg;
        option.textContent = alg;

        select.appendChild(option);
    }

    let div = document.createElement("div");
    div.id = `alg-${numOfSelects}-wrapper`;
    div.className = "select-wrapper";
    div.appendChild(select);

    document.getElementById("hash-form").insertAdjacentElement('beforeend',div);
}