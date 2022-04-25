import React, { useState } from 'react'
import { BrowserRouter, Route, Routes } from "react-router-dom"

import './App.css'

import NavBar from "../NavBar/NavBar"
import Intro from "../Intro/Intro"
import SignIn from "../SignIn/SignIn"
import SignUp from "../SignUp/SignUp"
import ExperimentChoice from "../ExperimentChoice/ExperimentChoice"
import Experiment from "../Experiment/Experiment"
import Profile from "../Profile/Profile"
import History from "../History/History";
import Results from "../Results/Results";


const App = () => {
    const [token, setToken] = useState("")
    const [label, setLabel] = useState("Войдите")

    const updateToken = (t) => setToken("Bearer " + t)
    const deleteToken = () => setToken("")

    const cipherAlgorithms =
        [
            "Кузнечик", "AES128-GCM", "AES128-CFB",
            "DES-GCM", "DES-CFB", "RSA", "Blowfish-CFB",
        ]

    const hashAlgorithms =
        [
            "Streebog-256","Streebog-512","SHA-224",
            "SHA-256","SHA-384","SHA-512","RIPEMD-128",
            "RIPEMD-160","RIPEMD-256","RIPEMD-320","MD5"
        ]


    return (
        <div className={"main"}>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<NavBar label={label} />}>
                        <Route index element={<Intro />}/>
                        <Route path={"/c/profile"} element={<Profile token={token} logout={deleteToken}/>}/>
                        <Route path={"/c/history"} element={<History token={token}/>}/>

                        <Route path={"/c/fetch-result/"} element={<Results token={token}/>}/>

                        <Route path={"/c/experiment"} element={<ExperimentChoice/>}/>
                        <Route path={"/c/experiment/hashes"} element={<Experiment list={hashAlgorithms}  token={token}/>}/>
                        <Route path={"/c/experiment/ciphers"} element={<Experiment list={cipherAlgorithms}  token={token}/>}/>

                        <Route path={"/c/auth/register"} element={<SignUp/>}/>
                        <Route path={"/c/auth/login"} element={<SignIn updateToken={updateToken} updateLabel={setLabel}/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </div>

    )
}

export default App