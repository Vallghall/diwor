import React from 'react'
import Plotly from "react-plotly.js"
import classes from "./Hyst.module.css"


const Hyst = ({configs}) => {
    if (!configs) return
    console.table(configs)
    const allocs = Object.values(configs).map(val => `${val.aloc}`)
    const ops = Object.values(configs).map(val => `${val.op}`)
    const algos = Object.values(configs).map(val => val.alg)
    const data = [
        {
            histfunc: "count",
            autobinx: false,
            x: algos,
            y: ops,
            type: "histogram",
            name: "Ops"
        },
        {
            autobinx: false,
            histfunc: "count",
            x: algos,
            y: allocs,
            type: "histogram",
            name: "Allocs"
        },
    ]

    const layout = {
        title : "Результаты эксперимента",
        xaxis: {
            title: 'Алгоритм',
            showgrid: true,
            zeroline: true
        },
        yaxis: {
            title: 'Херня, шт',
            showline: true,
            zeroline: true
        }
    }


    return (
        <div className={classes.plot}>
            <Plotly data={data} layout={layout}/>
        </div>

    )
}

export default Hyst