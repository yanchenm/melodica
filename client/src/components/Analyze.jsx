import {React, useState, useEffect}  from 'react';
import { BrowserRouter as Router, useLocation, Switch } from 'react-router-dom';
import axios from "axios";


function Analyze() {

    const [user, setUser] = useState("");

    useEffect(() => {
        const getUser = async () => {           

        const req = await axios.get("https://us-central1-musemood.cloudfunctions.net/User", {withCredentials: true});
        const data = await req.data;
        setUser(data.Name);
        console.log(data);
        }
        // const urlAddress = new URLSearchParams(props.location.search).get("code");
        // const a = urlAddress.get('code');

        getUser();
        
    }, [])

    return (
        <div>
            <h1> Hi, {} </h1>
        </div>
    )
}

export default Analyze;