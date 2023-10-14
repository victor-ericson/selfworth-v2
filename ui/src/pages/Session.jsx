import React, {useEffect, useState} from "react"
import {Endpoints} from "../api"
import {deleteCookie} from "../utils"
import Errors from "../components/Errors"
import Dashboard from "./Dashboard";
//TODO: Do I need to import this (Link) to every page or can I simply use it in the App.js?
import {Link} from "react-router-dom";

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
            //
            setIsFetching(true)
            //awaits fetch from Endpoints.session which fetches only the user
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
                                <p className="p">{user && user.password}</p>
                                <br/>
                                <button onClick={logout}>logout</button>
                                <Link to="/dashboard">Dashboard!</Link>
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
