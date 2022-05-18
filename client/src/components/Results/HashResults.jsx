import React, {useEffect, useState} from 'react'
import classes from "./Results.module.css"
import Plot from "../Plot/Plot"
import { useNavigate } from "react-router-dom"
import ResultRow from "../ResultRow/ResultRow"
import {tokenEffect} from "../../token"
import SysInfo from "../SysInfo/SysInfo"

const HashResults = ({token, params, renewToken}) => {
    const [results, setResults] = useState({})
    const [sysInfo, setSysInfo] = useState({os:"",arch:"", proc:""})
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
                        [res.algorithm]: res.plot,
                    }
                )
            ))
        })
        .catch(e => console.log(e))

    useEffect(() => tokenEffect(token, query, navigate, renewToken),[])

    return (
        <>
            {(sysInfo.os !== "" && sysInfo.arch !== "" && sysInfo.proc !== ""
                ? <SysInfo os={sysInfo.os} arch={sysInfo.arch} proc={sysInfo.proc}/>
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
                            <ResultRow fst={"Скорость"} snd={res.duration}/>
                            <ResultRow fst={"Длина дайждеста"} snd={res.size}/>
                            <ResultRow fst={"Размер блока"} snd={res.blockSize}/>
                            <ResultRow fst={"Пример хэша"} snd={res.sample}/>
                        </table>
                    </div>
                ))}

            </div>
        </>
    )
}

export default HashResults