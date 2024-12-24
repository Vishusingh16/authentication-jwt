import React , {useState} from 'react';
import axios from '../../api/api';
import { useNavigate } from 'react-router-dom';


function Signup(){
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        first_name:'',
        last_name:'',
        email:'',
        password:'',
    });
    const [error, setError] = useState('');

    const handleInputChange =(e) =>{
        const {name , value} = e.target;
        setFormData({
            ...formData ,
            [name]: value,

        });
    };
    const handleSubmit = async (e) =>{
        e.preventDefault();
        try{
            const response = await axios.post('/users/signup', formData);
            if(response.data.success){
                navigate('/login');

            }
        } catch(err){
            setError('Signup failed, please try again.');

        }
    };
    return(
        <div className ="form-container">
            <h2>Signup</h2>
            <form onSubmit={handleSubmit}>
            <input 
            type ="text"
            name="first_name"
            value={formData.first_name}
            onChange={handleInputChange}
            placeholder='First Name'
            required
             /> 
               <input 
               type ="text"
               name="last_name"
               value={formData.last_name}
               onChange={handleInputChange}
               placeholder='Last Name'
               required

             />
             <input 
               type ="email"
               name="email"
               value={formData.email}
               onChange={handleInputChange}
               placeholder='Email'
               required

             />
             <input 
               type ="text"
               name="phone"
               value={formData.phone}
               onChange={handleInputChange}
               placeholder='Phone'
               required

             />
             <input 
               type ="password"
               name="password"
               value={formData.password}
               onChange={handleInputChange}
               placeholder='Password'
               required

             />
             {error && <p className="error">{error}</p>}
             <button type='submit'>Sign Up</button>
            </form>
        </div>
    );
}
export default Signup;
