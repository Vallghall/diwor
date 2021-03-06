import React from 'react'
import classes from "./ResultRow.module.css"

const ResultRow = ({fst, snd}) => {
    return (
        <tr>
            <td>{fst}</td>
            <td className={(fst?.startsWith("Скорость") ? classes.duration : "")}>{snd}</td>
        </tr>
    )
}

export default ResultRow