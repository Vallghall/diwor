import React from 'react'
import Plotly from "react-plotly.js"
import classes from "./Plot.module.css"


const Plot = ({configs}) => {
    const data = Object.entries(configs).map(([alg, points]) =>
    ({
        x: points.x,
        y: points.y,
        type: "scatter",
        name: alg
    }))

    const layout = {
        title : "Результаты эксперимента",
        xaxis: {
            title: 'Длина сообщения, Б',
            showgrid: true,
            zeroline: true
        },
        yaxis: {
            title: 'Продолжительность, нс',
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

export default Plot