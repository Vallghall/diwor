import React, {useEffect, useState} from 'react'
import classes from "./Results.module.css"
import Plot from "../Plot/Plot"
import { useNavigate } from "react-router-dom"
import ResultRow from "../ResultRow/ResultRow";

const HashResults = ({token, params}) => {
    const [results, setResults] = useState({})
    const [plotConfigs, setPlotConfigs] = useState({})
    const navigate = useNavigate()

    useEffect(() => {
        if (token === "") {
            navigate("/c/auth/login")
            return
        }

        fetch(`/api/profile/fetch-result?sorted-id=${params.id}&alg-type=${params.alg}`,
            {
                method: "GET",
                headers: {
                    Authorization: token,
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
    },[])

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
                        <ResultRow fst={"Продолжительность"} snd={res.duration}/>
                        <ResultRow fst={"Длина дайждеста"} snd={res.size}/>
                        <ResultRow fst={"Размер блока"} snd={res.blockSize}/>
                        <ResultRow fst={"Пример хэша"} snd={res.sample}/>
                    </table>
                </div>
            ))}

        </div>
    )
}

export default HashResults