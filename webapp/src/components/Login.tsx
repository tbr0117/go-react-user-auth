/* eslint-disable import/no-anonymous-default-export */
import React, {Fragment} from "react"
import { useState } from "react"
import {CreateCookie} from "./../utils"
import {Endpoints} from "./../api"
import {Link } from "react-router-dom"


type LoginData = {
    email:string
    password:string
}
const Login = () => {
    const [oLoginData, setLoginData] = useState<LoginData>(()=>{
        return {
            email: "",
            password: ""
        }
    })
    const [isSubmitting, setIsSubmitting] = useState(false)
    const fnHandleChange = (e: any) =>{
        setLoginData({
            ...oLoginData, [e.target.name]: e.target.value
        })
    }

    const fnHandleLogin = async (evnt:any)=>{
        // history.push({ pathname: "/session", state: "user" })
        evnt.preventDefault()
        try {
          setIsSubmitting(true)
          const res = await fetch(Endpoints.login, {
            method: "POST",
            body: JSON.stringify(oLoginData),
            headers: {
              "Content-Type": "application/json",
            },
          })
    
          const { token, success, errors = [], user } = await res.json()
          if (success) {
            // creating a cookie expire in 30 minutes(same time as the token is invalidated on the backend)
            // ordinarily the setcookie from the server should suffice, however it has to be created here manually to bypass browsers
            // restriction on cross-site/non secure cookies on localhost.
            CreateCookie("token", token, 0.5)
          }
        //   setErrors(errors)
        } catch (e) {
        //   setErrors([e.toString()])
        } finally {
        //   setIsSubmitting(false)
        }
    }

    return( 
    <Fragment>
        <form onSubmit={fnHandleLogin}>
            <div className="LoginFrom">
            <input
                className="input"
                type="email"
                placeholder="email"
                value={oLoginData.email}
                name="email"
                onChange={fnHandleChange}
                required
                />
            <input
                className="input"
                type="password"
                placeholder="password"
                value={oLoginData.password}
                name="password"
                onChange={fnHandleChange}
                required
                />
                <button disabled={isSubmitting} type="submit">
                    {isSubmitting ? "....." : "login"}
                </button>
                <br />
                <Link to="register">{"create account"}</Link>
                <br />
                {/* <Error errors={Error} /> */}
            </div>
        </form>
        </Fragment>
    );
}


export default Login