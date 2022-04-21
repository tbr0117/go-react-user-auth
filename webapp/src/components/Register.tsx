/* eslint-disable import/no-anonymous-default-export */
import React, { useState } from "react";
import {Link} from "react-router-dom"
import {Endpoints} from "./../api"

type UserType = {
  email: string;
  password: string;
  name: string;
};

export default () => {
  const [oUser, setUser] = useState<UserType>({
      email: "",
      password: "",
      name: ""
  });

  const [isSubmitting, setSubmit] = useState<boolean>(false)

  const fnHandleChange = (evnt:any)=>{
    setUser({...oUser, [evnt.target.name]: evnt.target.value})
  }

  const fnHandleSubmit = async (evnt:any) => {
    evnt.preventDefault()
    try {
      setSubmit(true)
      const res = await fetch(Endpoints.register, {
        method: 'POST',
        body: JSON.stringify(oUser),
        headers: {
          'Content-Type': 'application/json',
        },
      });
      const { success, errors = [] } = await res.json()

    //   if (success) history.push('/login')

    } catch (e) {
    //   setErrors([e.toString()])
    } finally {
    //   setIsSubmitting(false)
    }
  };

  return (
    <form onSubmit={fnHandleSubmit}>
      <div className="LoginFrom">
        <input
          className="input"
          type="name"
          placeholder="Name"
          value={oUser.name}
          name="name"
          onChange={fnHandleChange}
          required
        />
        <input
          className="input"
          type="email"
          placeholder="Email"
          value={oUser.email}
          name="email"
          onChange={fnHandleChange}
          required
        />
        <input
          className="input"
          type="password"
          placeholder="Password"
          value={oUser.password}
          name="password"
          onChange={fnHandleChange}
          required
        />

        <button disabled={isSubmitting} onClick={fnHandleSubmit}>
          {isSubmitting ? "....." : "Sign Up"}
        </button>
        <br />
        <Link to="login">{"login"}</Link>
        <br />
      </div>
    </form>
  );
};
