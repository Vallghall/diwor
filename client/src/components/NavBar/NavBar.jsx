import React from 'react'
import { Outlet, Link } from 'react-router-dom'
import classes from "./NavBar.module.css"
import logo from './logo.png'
import UserBadge from "../UserBadge/UserBadge"

const NavBar = ({label}) => {

    return (
        <>
            <footer className={classes.footer}>
                <nav className={classes.navigation}>
                    <img className={classes.logo} src={logo} alt="Logo"/>
                    <ul className={classes.navigation__items}>
                        <li className={classes.navigation__item}>
                            <Link to={"/"}>Главная</Link>
                        </li>
                        <li className={classes.navigation__item}>
                            <Link to={"/c/profile"}>Профиль</Link>
                        </li>
                        <li className={classes.navigation__item}>
                            <Link to={"/c/experiment"}>Эксперимент</Link>
                        </li>
                    </ul>
                    <UserBadge label={label}/>
                </nav>
            </footer>

            <Outlet/>
        </>
    )
}

export default NavBar