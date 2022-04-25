import React, {useEffect, useState} from 'react'
import classes from "./SignUp.module.css"
import "./SignUp.module.css"
import {Link, useNavigate} from "react-router-dom"

const SignUp = () => {
    const [name, setName] = useState("")
    const [userName, setUserName] = useState("")
    const [pw, setPW] = useState("")
    const [confirmation, setConfirmation] = useState("")
    const navigate = useNavigate()

    const [errVisibility, setErrVisibility] = useState({
        username: "hidden",
        name: "hidden",
        password: "hidden",
        confirmation: "hidden"
    })

    const namePatternRU = /[а-яА-Я]+/g
    const namePatternEN = /[a-zA-Z]+/g

    useEffect(() => {
        if ((name.match(namePatternEN) && !name.match(namePatternRU)) || (!name.match(namePatternEN) && name.match(namePatternRU)))
            setErrVisibility({...errVisibility, name: "hidden"})
        else
            setErrVisibility({...errVisibility, name: "visible"})
    }, [name, userName, pw, confirmation])

    useEffect(() => {
        if (userName)
            setErrVisibility({...errVisibility, username: "hidden"})
        else
            setErrVisibility({...errVisibility, username: "visible"})
    }, [userName])

    useEffect(() => {
        if (pw.length < 6)
            setErrVisibility({...errVisibility, password: "visible"})
        else
            setErrVisibility({...errVisibility, password: "hidden"})
    }, [pw])

    useEffect(() => {
        if (pw === confirmation)
            setErrVisibility({...errVisibility, confirmation: "hidden"})
        else
            setErrVisibility({...errVisibility, confirmation: "visible"})
    }, [confirmation])

    const handleSubmit = e => {
        e.preventDefault()

        const query = JSON.stringify(
            {
                username: userName,
                name: name,
                password: pw,
            })

        fetch("/auth/sign-up",
            {
                method: "POST",
                body: query,
            })
            .then(resp => {
                if (resp.ok) navigate("/c/auth/login")
            })
            .catch()
    }

    return (
        <form className={classes.form} onSubmit={handleSubmit}>
            <h1>Регистрация</h1>
            <label htmlFor="name">Имя пользователя:</label><br/>
            <input type="text"
                   name="name"
                   value={name}
                   onChange={e => setName(e.target.value)}
            /> <span
            style={{color: "red", visibility: errVisibility.name}}>Допустимы символы [а-яА-Я] или [a-zA-Z]!</span> <br/><br/>

            <label htmlFor="username">Логин:</label><br/>
            <input type="text"
                   name="username"
                   value={userName}
                   onChange={e => setUserName(e.target.value)}
            /> <span
            style={{color: "red", visibility: errVisibility.username}}>Логин не должен быть пустым!</span> <br/><br/>

            <label htmlFor="password">Пароль:</label><br/>
            <input type="password"
                   name="password"
                   value={pw}
                   onChange={e => setPW(e.target.value)}
            /> <span
            style={{color: "red", visibility: errVisibility.password}}>Длина пароля должна составлять 6 и более символов!</span> <br/><br/>

            <label htmlFor="confirm-password">Подтверждение пароля:</label><br/>
            <input type="password"
                   name="confirm"
                   value={confirmation}
                   onChange={e => setConfirmation(e.target.value)}
            /> <span
            style={{color: "red", visibility: errVisibility.confirmation}}>Пароли должны совпадать!</span>  <br/><br/>

            <button type="submit">Подтвердить</button>
            <Link to="/c/auth/login">Уже зарегистрированы? Войти</Link>
        </form>
    )
}

export default SignUp