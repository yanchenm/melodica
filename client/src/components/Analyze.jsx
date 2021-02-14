import {React, useEffect, useState}  from 'react';
import { BrowserRouter as Router, Switch, useLocation } from 'react-router-dom';

import {api} from '../api';

function Analyze() {

    const [user, setUser] = useState("");

    useEffect(() => {
        const getUser = async () => {           

        const req = await api.get('/User');
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