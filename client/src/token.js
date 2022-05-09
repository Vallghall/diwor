const isTokenExpired = token => Date.now() >= (JSON.parse(window.atob(token.split('.')[1]))).exp * 1000

const getNewToken = async () => await fetch('/auth/refresh', {method: "POST", credentials: 'same-origin'})

export const tokenEffect = (token, query, navigate, renewToken) => {
    if (token === "") {
        navigate("/c/auth/login")
        return
    }
    if (isTokenExpired(token)) {
        getNewToken()
            .then(resp => {
                if (resp.status > 200) throw new Error(`err: ${resp.error}`)
                console.log("access token expired")
                return resp.json()
            })
            .then(body => {
                renewToken(body.token)
                localStorage.setItem("toketoken", body.token)
                return `Bearer ${body.token}`
            })
            .catch(_ => {
                console.log("should redirect")
                navigate("/c/auth/login")
            })
            .then(t => query(t))
            .catch(e => console.log(`err: ${e}`))
    } else {
        console.log("normal fetch")
        query()
            .catch(e => console.log(`err: ${e}`))
    }
}
