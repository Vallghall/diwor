import React from 'react'
import classes from "./Intro.module.css"

const Intro = () => {
    return (
        <section className={classes.intro}>
            Данное приложение является дипломным проектом студента <br/>
            Московского Технического Университета Связи и Информатики <br/>
            Гусева Р.М. из группы БПЗ 1802
        </section>
    )
}

export default Intro