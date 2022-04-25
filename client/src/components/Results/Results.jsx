import React from 'react'
import {useSearchParams} from "react-router-dom"
import HashResults from "./HashResults"
import CipherResults from "./CipherResults"

const Results = ({token}) => {
    const [params, _] = useSearchParams()

    return ( (params.get("alg") === "Алгоритм шифрования")
            ? <CipherResults token={token} params={{id: params.get("id"), alg: params.get("alg")}}/>
            : <HashResults token={token} params={{id: params.get("id"), alg: params.get("alg")}}/>
    )
}

export default Results