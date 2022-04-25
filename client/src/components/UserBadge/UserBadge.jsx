import React from 'react'
import { Link } from "react-router-dom"
import classes from "./UserBadge.module.css"


const UserBadge = ({label}) => {
    return (
        <div className={classes.badge}>
            {(label === "Войдите")
                ? <Link to={"/c/auth/login"}>{label}</Link>
                : <Link to={"/c/profile"}>{label}</Link>}

        </div>
    )
}

export default UserBadge