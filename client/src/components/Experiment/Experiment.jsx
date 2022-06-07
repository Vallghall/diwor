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
    const [plotInfo, setPlotInfo] = useState({from: 0, to: 128, step: 16})
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
        if (algs.length === 10) {
            SweetAlert("Стоп", "Больше десяти экспериментов установить нельзя.", "warning")
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
            if (!Number(pi) && pi !== 0) {
                SweetAlert("Некорректные значения", "Введеные параметры графика являются некорректными!", "warning").catch()
                return
            }

        const query = JSON.stringify({
            from: plotInfo.from * 1024,
            to: plotInfo.to * 1024,
            step: plotInfo.step * 1024,
            algorithms: Array.from(new Set(algs)),
        })

        SweetAlert(
            "Запуск расчетов",
            "Вы сможете ознакомиться с результатами эксперимента через Ваш профиль, когда они будут готовы",
            "success"
        ).catch()

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
            .catch()
    }

    return (
        <form className={classes.form} onSubmit={handleSubmit}>
            <h1>Выберите алгоритмы для сравнения</h1>

            <strong>Установка пределов для графика (в килобайтах)</strong>
            <div className={classes.inputs}>
                <span>От</span><span>До</span><span>Шаг</span>
                <input type="text" name="from" defaultValue={plotInfo.from} onChange={(e) => {
                    setPlotInfo({...plotInfo, from: +e.target.value})
                }}/>
                <input type="text" name="to" defaultValue={plotInfo.to} onChange={(e) => {
                    setPlotInfo({...plotInfo, to: +e.target.value})
                }}/>
                <input type="text" name="step" defaultValue={plotInfo.step} onChange={(e) => {
                    setPlotInfo({...plotInfo, step: +e.target.value})
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