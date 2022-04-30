import React, {useEffect, useState} from 'react'
import classes from "./Experiment.module.css"
import "./Experiment.module.css"
import Selection from "./Selection"
import SweetAlert from "sweetalert"
import {useNavigate} from "react-router-dom"
import {tokenEffect} from "../../token";



const Experiment = ({list, token, renewToken}) => {
    const [algs, setAlgs] = useState([""])
    const [startDisability, setStartDisability] = useState(true)
    const [plotInfo, setPlotInfo] = useState({from: 0, to: 0, step: 0, num_measurements: 0})
    const navigate = useNavigate()


    useEffect(() => {
        for (const alg of algs) {
            if (!alg)
                setStartDisability(true)
            else
                setStartDisability(false)
        }
    }, [algs])

    useEffect(() => tokenEffect(token, async () => {}, navigate, renewToken), [])
    useEffect(() => tokenEffect(token, async () => {}, navigate, renewToken), [startDisability])

    const addAlg = (alg, i) => {
        const arr =[...algs]
        arr[i] = alg
        setAlgs(arr)
    }

    const plus = () => {
        if (algs.length === 6) {
            SweetAlert("Стоп", "Больше шести нельзя!", "warning")
            return
        }

        setAlgs([...algs, ""])
    }
    const minus = () => {
        if (algs.length === 1) {
            SweetAlert("Стоп", "Меньше одного алгоритма выбрать нельзя!", "warning").catch()
            return
        }

        const arr = [...algs]
        arr.pop()
        setAlgs(arr)
    }

    const handleSubmit = e => {
        e.preventDefault()

        for (const pi of Object.values(plotInfo))
            if (!Number(pi)) {
                SweetAlert("Некорректные значения", "Введеные параметры графика являются некорректными!", "warning").catch()
                return
            }

        const query = JSON.stringify({
            ...plotInfo,
            algorithms: algs,
        })

        fetch((window.location.pathname === "/c/experiment/hashes")
            ? "/api/experiment/start-hash-experiment"
            : "/api/experiment/start-cipher-experiment",
            {
                method: "POST",
                headers: {
                    Authorization: token,
                },
                body: query,
            })
            .then(_ => {
                SweetAlert(
                    "Запуск расчетов",
                    "Вы сможете ознакомиться с результатами эксперимента через Ваш профиль, когда они будут готовы",
                    "success"
                    ).catch()
            })
            .catch()
    }

    return (
        <form className={classes.form} onSubmit={handleSubmit}>
            <h1>Выберите алгоритмы для сравнения</h1>

            <strong>Установка пределов для графика</strong>
            <div className={classes.inputs}>
                <input type="text" name="from" placeholder="От" onChange={(e) => {
                    setPlotInfo({...plotInfo, from: +e.target.value})
                }}/>
                <input type="text" name="to" placeholder="До" onChange={(e) => {
                    setPlotInfo({...plotInfo, to: +e.target.value})
                }}/>
                <input type="text" name="step" placeholder="Шаг" onChange={(e) => {
                    setPlotInfo({...plotInfo, step: +e.target.value})
                }}/>
                <input type="text" name="num" placeholder="Число замеров" onChange={(e) => {
                    setPlotInfo({...plotInfo, num_measurements: +e.target.value})
                }}/>
            </div>

            <div className={classes.select_wrapper}>
                <input disabled={startDisability} type="submit" id="begin" value="Начать"
                       className={classes.quantity_controller}
                />
                <div id="plus" onClick={plus}
                     className={[classes.quantity_controller, classes.operators].join(" ")}>+</div>
                <div id="minus" onClick={minus}
                     className={[classes.quantity_controller, classes.operators].join(" ")}>-</div>
            </div>

            {Object.keys(algs).map(i => <Selection list={list} id={i} addAlg={addAlg} key={i}/>)}
        </form>
    )
}

export default Experiment