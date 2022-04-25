import React, {useEffect, useState} from 'react'
import DigestTable from "../DigestTable/DigestTable"
import {useNavigate} from "react-router-dom"
import classes from "./History.module.css"

const History = ({token}) => {
    const [history, setHistory] = useState({})
    const navigate = useNavigate()

    useEffect(() => {
        if (token === "") {
            navigate("/c/auth/login")
            return
        }

        fetch("/api/profile/results/",
            {
                method: "GET",
                headers: {
                    Authorization: token,
                },
            })
            .then(resp => resp.json())
            .then(body => setHistory(body.History))
            .catch()
    },[])

    return (
        <div className={classes.wrapper}>
            <DigestTable digests={history}/>
        </div>
    )
}

export default History