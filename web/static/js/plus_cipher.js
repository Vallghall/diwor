const cipherAlgorithms = [
    "Кузнечик",
    "AES128-GCM",
    "AES128-CFB",
    "DES-GCM", "DES-CFB",
    "RSA",
    "BF-GCM",
    "BF-CFB",
]

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
    defaultOption.textContent = "Выберите алгоритм шифрования";
    select.appendChild(defaultOption);

    for (const alg of cipherAlgorithms) {
        let option = document.createElement("option");
        option.value = alg;
        option.textContent = alg;

        select.appendChild(option);
    }

    let div = document.createElement("div");
    div.id = `alg-${numOfSelects}-wrapper`;
    div.className = "select-wrapper";
    div.appendChild(select);

    document.getElementById("cipher-form").insertAdjacentElement('beforeend',div);
}