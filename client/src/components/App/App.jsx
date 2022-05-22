import React, {useEffect, useState} from 'react'
import {BrowserRouter, Link, Route, Routes} from "react-router-dom"

import './App.css'

import NavBar from "../NavBar/NavBar"
import Intro from "../Intro/Intro"
import SignIn from "../SignIn/SignIn"
import SignUp from "../SignUp/SignUp"
import ExperimentChoice from "../ExperimentChoice/ExperimentChoice"
import Experiment from "../Experiment/Experiment"
import Profile from "../Profile/Profile"
import History from "../History/History"
import Results from "../Results/Results"

const App = () => {
    const [token, setToken] = useState("")
    const [label, setLabel] = useState("Войдите")

    const updateToken = t => setToken("Bearer " + t)
    const deleteToken = () => {
        fetch("/auth/logout")
            .then(_ => setToken(""))
            .then(_ => setLabel("Войдите"))
            .then(_ => window.location.href="/c/auth/login")
            .catch()
    }

    useEffect(() => {
        if (token === "") {
            const t = localStorage.getItem("toketoken")
            if (t)
                updateToken(t)
        }
    })

    const cipherAlgorithms =
        [
            "Кузнечик-ECB", "Кузнечик-CFB", "Кузнечик-GCM", "Кузнечик-OFB", "Кузнечик-CTR",
            "AES128-ECB", "AES128-GCM", "AES128-CFB", "AES128-OFB", "AES128-CTR",
            "Магма-ECB", "Магма-CFB", "Магма-OFB", "Магма-CTR",
            "DES-ECB", "DES-CFB", "DES-OFB", "DES-CTR",
            "Blowfish-ECB", "Blowfish-CFB", "Blowfish-OFB", "Blowfish-CTR",

            "RSA",
        ]

    const hashAlgorithms =
        [
            "Streebog-256", "Streebog-512", "bcrypt", "MD5",
            "SHA-224", "SHA-256", "SHA-384", "SHA-512",
            "SHA3-224", "SHA3-256", "SHA3-384", "SHA3-512",
            "SHA3-SHAKE128", "SHA3-SHAKE256",
            "RIPEMD-128", "RIPEMD-160", "RIPEMD-256", "RIPEMD-320",
        ]


    return (
        <div className={"main"}>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<NavBar label={label} />}>
                        <Route index element={<Intro />}/>
                        <Route path={"/c/profile"} element={<Profile token={token} logout={deleteToken} renewToken={updateToken}/>}/>
                        <Route path={"/c/history"} element={<History token={token} renewToken={updateToken}/>}/>

                        <Route path={"/c/fetch-result/"} element={<Results token={token} renewToken={updateToken}/>}/>

                        <Route path={"/c/experiment"} element={<ExperimentChoice/>}/>
                        <Route path={"/c/experiment/hashes"} element={<Experiment list={hashAlgorithms}  token={token} renewToken={updateToken}/>}/>
                        <Route path={"/c/experiment/ciphers"} element={<Experiment list={cipherAlgorithms}  token={token} renewToken={updateToken}/>}/>

                        <Route path={"/c/auth/register"} element={<SignUp/>}/>
                        <Route path={"/c/auth/login"} element={<SignIn updateToken={updateToken} updateLabel={setLabel}/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </div>

    )
}

export default App