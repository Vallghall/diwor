import React, {useEffect, useState} from 'react'
import classes from "./Results.module.css"
import Plot from "../Plot/Plot"
import { useNavigate } from "react-router-dom"
import ResultRow from "../ResultRow/ResultRow"
import {tokenEffect} from "../../token"
import SysInfo from "../SysInfo/SysInfo"

const CipherResults = ({token, params, renewToken}) => {
    const [results, setResults] = useState({})
    const [sysInfo, setSysInfo] = useState({os:"",arch:""})
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
            const Results = body["Results"]
            setSysInfo({
                os: Results.os,
                arch: Results.arch,
            })
            setResults(Results.results)

            Object.values(Results.results).forEach(res => setPlotConfigs(p => (
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
        <>
            {(sysInfo.os !== "" && sysInfo.arch !== ""
                ? <SysInfo os={sysInfo.os} arch={sysInfo.arch}/>
                : "")}
            <div className={classes.wrapper}>

                <Plot configs={plotConfigs}/>
                {Object.values(results).map(res => (
                    <div className={classes.result_wrapper}>
                        <table>
                            <tr>
                                <th>{"Параметры"}</th>
                                <th>{"Результаты"}</th>
                            </tr>
                            <ResultRow fst={"Алгоритм"} snd={res.algorithm}/>
                            <ResultRow fst={"Тип"} snd={res.type}/>
                            <ResultRow fst={"Скорость шифрования"} snd={res.ciphering_duration}/>
                            <ResultRow fst={"Скорость дешифрования"} snd={res.deciphering_duration}/>
                            <ResultRow fst={"Длина ключа"} snd={res.key_length}/>
                        </table>
                    </div>
                ))}

            </div>
        </>
    )
}

export default CipherResults