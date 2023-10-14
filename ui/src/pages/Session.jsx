import React, {useEffect, useState} from "react"
import {Endpoints} from "../api"
import {deleteCookie} from "../utils"
import Errors from "../components/Errors"
import Dashboard from "./Dashboard";

//TODO: FIX THIS
/*Error: Invalid hook call. Hooks can only be called inside of the body of a function component. This could happen for one of the following reasons:
1. You might have mismatching versions of React and the renderer (such as React DOM)
2. You might be breaking the Rules of Hooks
3. You might have more than one copy of React in the same app
See https://fb.me/react-invalid-hook-call for tips about how to debug and fix this problem.*/
const Session = ({history}) => {
    const [user, setUser] = useState(null)
    const [isFetching, setIsFetching] = useState(false)
    const [errors, setErrors] = useState([])

    const headers = {
        Accept: "application/json",
        Authorization: document.cookie.split("token=")[1],
    }

    const getUserInfo = async () => {
        try {
            setIsFetching(true)
            const res = await fetch(Endpoints.session, {
                method: "GET",
                credentials: "include",
                headers,
            })

            if (!res.ok) logout()

            const {success, errors = [], user} = await res.json()
            setErrors(errors)
            if (!success) history.push("/login")
            setUser(user)
        } catch (e) {
            setErrors([e.toString()])
        } finally {
            setIsFetching(false)
        }
    }

    const logout = async () => {
        const res = await fetch(Endpoints.logout, {
            method: "GET",
            credentials: "include",
            headers,
        })

        if (res.ok) {
            deleteCookie("token")
            history.push("/login")
        }
    }
    const dashboard = async () => {
        const res = await fetch(Endpoints.dashboard, {
            method: "GET",
            credentials: "include",
            headers,
        })

        if (res.ok) {
            history.push("/dashboard")
        }
    }

    useEffect(() => {
        getUserInfo()
    }, [])

    return (
        <div className="wrapper">
            <div>
                {isFetching ? (
                    <div>fetching details...</div>
                ) : (
                    <div>
                        {user && (
                            <div>
                                <h1 className="h1">Welcome, {user && user.name}</h1>
                                <p className="p">{user && user.email}</p>
                                <br/>
                                <button onClick={logout}>logout</button>
                                <button onClick={Dashboard}>Dashboard!</button>
                            </div>
                        )}
                    </div>
                )}

                <Errors errors={errors}/>
            </div>
        </div>
    )
}

export default Session
