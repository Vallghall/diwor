const form = document.getElementsByName('login-info')[0];
form.addEventListener('submit', event => {
    event.preventDefault()

    const data = new FormData(event.target);
    const value = Object.fromEntries(data.entries());
    console.log(value);
    jsonBodyQuery('POST', '/auth/sign-in','/api/experiment', value)
});
