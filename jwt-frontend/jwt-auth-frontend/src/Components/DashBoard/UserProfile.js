import React, {useEffect, useState} from 'react';
import axios from "../../api/api";

function UserProfile(){
    const [userData, setUserData] = useState(null);
    const[loading, setLoading] = useState(true);

    useEffect(() =>{
        const fetchUserData = async ()=>{
            try{
                const response = await axios.get('/users/me');
                setUserData(response.data);

            }catch(err){
                console.log("Error Fetching user data", err);

            }finally{
                setLoading(false);

            }

        };
        fetchUserData();
    }, []);

    if(loading) return <p> Loading...</p>;
    return(
        <div className='profile-container'>
            <h2> User Profile </h2>
            {userData ?(
                <div>
                    <p>Name :{userData.first_name}{userData.last_name}</p>
                    <p>Email: {userData.email}</p>
                    <p>Phone: {userData.phone}</p>
                    </div>
            ):(
                <p> Unable to fetch user data</p>
            )}
        </div>
    );

}
export default UserProfile;
