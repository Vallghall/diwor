import React, {useState} from 'react'
import {Link, useNavigate} from "react-router-dom"
import SweetAlert from "sweetalert";
import classes from "./SignIn.module.css"
import "./SignIn.module.css"

const SignIn = ({updateToken, updateLabel}) => {
    const [input, setInput] = useState({
        username: "",
        name: "",
        password: "",
        confirmation: ""
    })
    const navigate = useNavigate()

    const handleSubmit = (e) => {
        e.preventDefault()

        const query = JSON.stringify(input)
        fetch('/auth/sign-in', {
            method: "POST",
            body: query
        })
            .then(resp => resp.json())
            .then(body => updateToken(body.token))
            .then(_ => updateLabel(input.username))
            .then(() => navigate("/"))
            .catch(e => SweetAlert("Bad request", e, "warning"))
    }
    return (
        <form className={classes.form}  onSubmit={handleSubmit}>
            <h1>Вход в систему</h1>
            <label htmlFor="username">Логин:</label><br/>
            <input type="text"
                   name="username"
                   value={input.username + ""}
                   onChange={(e)=>{
                       setInput({...input, username: e.target.value})
                   }}
            /><br/><br/>

            <label htmlFor="password">Пароль:</label><br/>
            <input type="password"
                   name="password"
                   value={input.password + ""}
                   onChange={(e)=>{
                       setInput({...input, password: e.target.value})
                   }}
            /> <br/><br/>
            <button type="submit">Подтвердить</button>
            <Link to="/c/auth/register">Не зарегистрированы?</Link>
        </form>
    )
}

export default SignIn