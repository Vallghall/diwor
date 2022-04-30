import React, {useEffect, useState} from 'react'
import classes from "./Results.module.css"
import Plot from "../Plot/Plot"
import { useNavigate } from "react-router-dom"
import ResultRow from "../ResultRow/ResultRow";
import {tokenEffect} from "../../token";

const CipherResults = ({token, params, renewToken}) => {
    const [results, setResults] = useState({})
    const [plotConfigs, setPlotConfigs] = useState({})
    const navigate = useNavigate()

    const query = (t = token) => fetch(`/api/profile/fetch-result?sorted-id=${params.id}&alg-type=${params.alg}`,
        {
            method: "GET",
            headers: {
                Authorization: t,
            }
        })
        .then(resp => resp.json())
        .then(body => {
            const Results = body["Results"].results
            console.log(Results)
            setResults(Results)
            Object.values(Results).forEach(res => setPlotConfigs(p => (
                    {
                        ...p,
                        [res.algorithm]: res.plot
                    }
                )
            ))

        })
        .catch(e => console.log(e))

    useEffect(() => tokenEffect(token, query, navigate, renewToken),[])

    return (
        <div className={classes.wrapper}>
            <Plot congigs={plotConfigs}/>
            {Object.values(results).map(res => (
                <div className={classes.result_wrapper}>
                    <table>
                        <tr>
                            <th>{"Параметры"}</th>
                            <th>{"Результаты"}</th>
                        </tr>
                        <ResultRow fst={"Алгоритм"} snd={res.algorithm}/>
                        <ResultRow fst={"Тип"} snd={res.type}/>
                        <ResultRow fst={"Продолжительность шифрования"} snd={res.ciphering_duration}/>
                        <ResultRow fst={"Продолжительность дешифрования"} snd={res.deciphering_duration}/>
                        <ResultRow fst={"Длина ключа"} snd={res.key_length}/>
                    </table>
                </div>
            ))}

        </div>
    )
}

export default CipherResults