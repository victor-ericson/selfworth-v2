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
                <div className="login-page">
                    <h1 className="h1">SelfWorth</h1>
                    <div className="content-section">
                        <p className="p">Because nothing makes you save more money than knowing how you spend it</p>
                        <input className="input" type="email" placeholder="Email" value={email} name="email"
                               onChange={handleChange} required/>
                        <input className="input" type="password" placeholder="Password" value={password} name="password"
                               onChange={handleChange} required/>
                        <button disabled={isSubmitting} type="submit">{isSubmitting ? "....." : "Login"}</button>
                        <br/>
                        <a href="/register" className="p" style={{fontSize: '0.8rem'}}>Create Account</a>
                        <Errors errors={errors}/>
                    </div>
                </div>
            </div>

        </form>
    )
}
