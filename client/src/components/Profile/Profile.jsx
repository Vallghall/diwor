import React, {useEffect, useState} from 'react'
import classes from "./Profile.module.css"
import { useNavigate } from "react-router-dom"
import DigestTable from "../DigestTable/DigestTable"
import {tokenEffect} from "../../token"

const Profile = ({token, logout, renewToken}) => {
    const [credentials, setCredentials] = useState({name: "", username: ""})
    const [digests, setDigests] = useState({})
    const navigate = useNavigate()

    const query = (t = token) => fetch("/api/profile/", {
        method: "GET",
        headers: {
            Authorization: t
        }
    })
        .then(resp => {
            if (resp.status > 200) throw new Error(`StatusError: ${resp.error}`)
            return resp.json()
        })
        .then(body => {
            setCredentials(
                {
                    username: body.username,
                    name: body.name,
                })
            setDigests(body.Digests)

        })
        .catch(e => console.log(`FetchError: ${e}`))

    useEffect(() => tokenEffect(token, query, navigate, renewToken),[])

    return (
        <div className={classes.wrapper}>

            <div className={classes.profile_header}>
                <strong>Имя пользователя:</strong>
                <br/>
                {credentials.name}
                <br/>
                <strong>Никнейм:</strong>
                <br/>
                {credentials.username}
                <br/><br/>
                <button className={classes.logout}
                        type="button"
                        onClick={_ => logout()}>Выйти</button>
                <button className={classes.history}
                        type="button"
                        onClick={_ => navigate("/c/history")}>История</button>
            </div>

            <div className={classes.profile_component_footer}>
                <DigestTable digests={digests}/>
            </div>
        </div>)
}

export default Profile