import React from 'react';
// import {BrowserRouter as Router , Route, Switch} from 'react-router-dom';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Signup from './Components/Auth/Signup';
import Login from './Components/Auth/Login';
import UserProfile from './Components/DashBoard/UserProfile';
import './App.css';


function App(){
  return(
    <Router>
      <div className='app'>
        <h1> JWT Authentication</h1>
        <Routes>
          <Route path='/signup' Component={Signup} />
          <Route path='/login' Component={Login} />
          <Route path="/profile" Component={UserProfile} />

        </Routes>
      </div>
    </Router>
  );
}

export default App;
