import React, {useEffect, useState} from 'react'
import DigestTable from "../DigestTable/DigestTable"
import {useNavigate} from "react-router-dom"
import classes from "./History.module.css"
import {tokenEffect} from "../../token";

const History = ({token, renewToken}) => {
    const [history, setHistory] = useState({})
    const navigate = useNavigate()

    const query = (t = token) => fetch("/api/profile/results/",
            {
                method: "GET",
                headers: {
                    Authorization: t,
                },
            })
            .then(resp => resp.json())
            .then(body => setHistory(body.History))
            .catch()

    useEffect(() => tokenEffect(token, query, navigate, renewToken),[])

    return (
        <div className={classes.wrapper}>
            <DigestTable digests={history}/>
        </div>
    )
}

export default History