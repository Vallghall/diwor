import React from 'react'
import {Link} from "react-router-dom"

const ProfileTable = ({digests}) => {
    if (!digests) return (
        <tr>
            <td>{'-'}</td>
            <td>{'-'}</td>
            <td>{'-'}</td>
        </tr>
    )
    return (Object.values(digests).map(({sorted_id, algorithm_type, started_at}) =>
        <tr>
            <td><Link to={`/c/fetch-result?id=${sorted_id}&alg=${algorithm_type}`}>{sorted_id}</Link></td>
            <td>{algorithm_type}</td>
            <td>{started_at}</td>
        </tr>
    ))
}

export default ProfileTable