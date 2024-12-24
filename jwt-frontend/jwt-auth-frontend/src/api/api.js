import axios from 'axios';

const API_URL ='http://localhost:8000';

const api = axios.create({
    baseURL : API_URL,
    headers :{
        'Content-Type' : 'application/json',

    }
});

export const setAuthToken = (token) =>{
    if(token){
        api.defaults.headers['Authorization'] = `Bearer ${token}`;

    }else{
        delete api.defaults.headers['Authorization'];

    }
};
export default api;
