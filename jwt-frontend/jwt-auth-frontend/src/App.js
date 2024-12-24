import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Signup from './Components/Auth/Signup';
import Login from './Components/Auth/Login';
import UserProfile from './Components/DashBoard/UserProfile';
import './App.css';


function App(){
  return(
    
    <BrowserRouter>
      <div className='app'>
        <h1> JWT Authentication</h1>
        <Routes>
        <Route path="/" element={<h2>Welcome to the Homepage</h2>} /> 
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/profile" element={<UserProfile />} />
      </Routes>
      </div>
      </BrowserRouter>
 
  );
}

export default App;


