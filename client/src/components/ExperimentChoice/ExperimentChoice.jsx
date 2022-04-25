import React from 'react'
import { useNavigate } from 'react-router-dom'
import classes from "./ExperimentChoice.module.css"

const ExperimentChoice = () => {
    const navigate = useNavigate()
    return (

        <div className={classes.choice_panel}>
            <h1>Выберите тип алгоритма для эксперимента</h1>

            <button type="submit" className={[classes.choose_between,classes.hash].join(" ")}
                        onClick={()=>navigate('/c/experiment/hashes')}>Алгоритмы хэширования</button>

            <button type="submit" className={[classes.choose_between,classes.cipher].join(" ")}
                        onClick={()=>navigate('/c/experiment/ciphers')}>Алгоритмы шифрования</button>
        </div>

    )
}

export default ExperimentChoice