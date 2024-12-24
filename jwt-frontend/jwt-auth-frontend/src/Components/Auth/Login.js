import React , {useState} from 'react';
import { useNavigate } from 'react-router-dom'; 
import {setAuthToken} from '../../api/api';
import axios from "../../api/api";


function Login(){
    const navigate = useNavigate();

    const [formData , setFormData] = useState({
        email:'',
        password:'',

    });
    const [error, setError] = useState('');

    const handleInputChange =(e)=>{
        const {name,value} =  e.target;
        setFormData({
            ...formData ,
            [name]:value,

        });
    };

    const handleSubmit =async(e) =>{
        e.preventDefault();
        try{
            const response = await axios.post('/users/login', formData);
            if(response.data.token){
                localStorage.setItem('token', response.data.token);
                setAuthToken(response.data.token);
                navigate('/profile');
            }
        }catch(err){
            setError("Login failed, please try again");
        }
    };
    return(
        <div className='form-container'>
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <input 
                type ="email"
                name='email'
                value={formData.email}
                onChange={handleInputChange}
                placeholder="Email"
                required
                
                />
                 <input 
                type ="password"
                name='password'
                value={formData.password}
                onChange={handleInputChange}
                placeholder="Password"
                required
                
                />
                {error && <p className='error'>{error}</p>}
                <button type='Submit'>Login</button>

            </form>
            </div>
    );

}
export default Login;