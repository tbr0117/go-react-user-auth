import React from 'react';
import {Routes, Route} from "react-router-dom"
import Login from "./components/Login"
import Register from './components/Register';
import './App.css';

function App() {
  return (
      <Routes>
          <Route  path="/" element={ <div /> } />
          <Route path='login' element={<Login />} />
          <Route path="register" element={<Register />} />
      </Routes>
  );
}

export default App;
