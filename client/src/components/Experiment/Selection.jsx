import React from 'react'
import classes from "./Experiment.module.css"


const Selection = ({id, addAlg, list}) => {
    return (
        <div className={classes.select_wrapper}>
            <select id={`alg-${id}`} className="select-alg" onChange={e => addAlg(e.target.value, id)}>
                <option selected disabled>Выберите алгоритм</option>
                {list.map(alg => <option value={alg} key={alg+id}>{alg}</option>)}
            </select>
        </div>
    )
}

export default Selection