import React, {useState} from "react"

import {Endpoints} from "../api"
import Errors from "../components/Errors"
import {createCookie} from "../utils"

export default ({history}) => {
    const [login, setLogin] = useState({
        email: "",
        password: "",
    })

    const [isSubmitting, setIsSubmitting] = useState(false)
    const [errors, setErrors] = useState([])

    const {email, password} = login

    const handleChange = (e) =>
        setLogin({...login, [e.target.name]: e.target.value})

    const handleSubmit = async (e) => {
        e.preventDefault()
        const {email, password} = login
        try {
            setIsSubmitting(true)
            //fetching my endpoints, here I fetch my Login from
            const res = await fetch(Endpoints.login, {
                method: "POST",
                body: JSON.stringify({
                    email,
                    password,
                }),
                headers: {
                    "Content-Type": "application/json",
                },
            })

            const {token, success, errors = [], user} = await res.json()
            if (success) {
                createCookie("token", token, 0.5)
                history.push({pathname: "/session", state: user})
            }
            setErrors(errors)
        } catch (e) {
            setErrors([e.toString()])
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <div className="wrapper">
                <h1>Login</h1>

                <input
                    className="input"
                    type="email"
                    placeholder="email"
                    value={email}
                    name="email"
                    onChange={handleChange}
                    required
                />

                <input
                    className="input"
                    type="password"
                    placeholder="password"
                    value={password}
                    name="password"
                    onChange={handleChange}
                    required
                />

                <button disabled={isSubmitting} type="submit">
                    {isSubmitting ? "....." : "login"}
                </button>
                <br/>
                <a href="/register">{"create account"}</a>
                <br/>
                <Errors errors={errors}/>
            </div>
        </form>
    )
}
