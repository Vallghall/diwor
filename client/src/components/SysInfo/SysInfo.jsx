import React from 'react'
import classes from "./SysInfo.module.css";

const SysInfo = ({os, arch, proc}) => {
    return (
        <div className={classes.sysinfo}>
            {os} <br/> {arch} <br/> {proc}
        </div>
    )
}

export default SysInfo