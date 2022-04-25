import React, {useEffect, useState} from 'react'
import classes from "./Profile.module.css"
import { useNavigate } from "react-router-dom"
import DigestTable from "../DigestTable/DigestTable";

const Profile = ({token, logout}) => {
    const [credentials, setCredentials] = useState({name: "", username: ""})
    const [digests, setDigests] = useState({})
    const navigate = useNavigate()

    useEffect(() => {
        if (!token) {
            navigate("/c/auth/login")
            return
        }

        fetch("/api/profile/", {
            method: "GET",
            headers: {
                Authorization: token
            }
        })
            .then(resp => resp.json())
            .then(body => {
                setCredentials(
                    {
                        username: body.username,
                        name: body.name,
                    })
                setDigests(body.Digests)

            })
            .catch(e => console.log(e))
    }, [])

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