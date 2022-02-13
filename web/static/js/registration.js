const defaultPattern = /^[а-яА-Яa-zA-Z0-9]+$/g;
const namePatternRU = /^[а-яА-Я]+$/g
const namePatternEN = /^[a-zA-Z]+$/g

let passwordMessage = document.getElementById('password-message')
let confirmPasswordMessage = document.getElementById('confirm-password-message')
let nameMessage = document.getElementById("name-message")
let userNameMessage = document.getElementById("username-message")

const form = document.getElementsByName('login-info')[0];
form.addEventListener('submit', e => {
    e.preventDefault()

    const data = new FormData(e.target);
    const value = Object.fromEntries(data.entries());
    console.log({value})
    let allDone = 0

    if (value.password.length < 6) {
        passwordMessage.textContent = "Длина пароля должна превышать 6 символов";
        passwordMessage.setAttribute('style', 'color: red;');
    } else if (!value.password.match(defaultPattern)) {
        passwordMessage.textContent = "Пароль может содержать только символы латинского" +
            " алфавита, кириллицу и числа от нуля до 9";
        passwordMessage.setAttribute('style', 'color: red;');
    } else {
        passwordMessage.textContent = "Надежная длина пароля";
        passwordMessage.setAttribute('style', 'color: green;');
        allDone++
    }

    if (value.password !== value.confirm) {
        confirmPasswordMessage.textContent = "Пароли должны совпадать!";
        confirmPasswordMessage.setAttribute('style', 'color: red;');
    } else {
        confirmPasswordMessage.textContent = "Пароли совпадают";
        confirmPasswordMessage.setAttribute('style', 'color: green;');
        allDone++
    }

    if (value.name === "") {
        nameMessage.textContent = "Это поле не может быть пустым!";
        nameMessage.setAttribute('style', 'color: red;');
    } else if (!value.name.match(namePatternRU) && !value.name.match(namePatternEN)) {
        nameMessage.textContent = "Допускаются только символы русского или латинского алфавитов";
        nameMessage.setAttribute('style', 'color: red;');
    } else {
        nameMessage.setAttribute('style', 'color: white;');
        allDone++
    }

    if (value.username === "") {
        userNameMessage.textContent = "Это поле не может быть пустым!";
        userNameMessage.setAttribute('style', 'color: red;');
    } else {
        userNameMessage.setAttribute('style', 'color: white;');
        allDone++
    }

    if (allDone === 4) {
        jsonBodyQuery('POST', '/auth/sign-up','/auth/login', value)
    }
});
