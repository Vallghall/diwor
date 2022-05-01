import React from 'react'
import classes from "./SysInfo.module.css";

const SysInfo = ({os, arch}) => {
    return (
        <div className={classes.sysinfo}>
            {os} <br/> {arch}
        </div>
    )
}

export default SysInfo