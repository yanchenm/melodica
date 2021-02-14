import {React, useState, useEffect}  from 'react';
import { BrowserRouter as Router, useLocation, Switch } from 'react-router-dom';



function Analyze() {

    const [code, setCode] = useState("");

    useEffect(() => {
    
        // const urlAddress = new URLSearchParams(props.location.search).get("code");
        // const a = urlAddress.get('code');

        const url = window.location.href;
        const params = new URLSearchParams(url);
        const text = params.toString();
        const code = text.substring(text.indexOf("code=") + 5);
        console.log(code);
        setCode(code);

    }, [])

    return (
        <div>
            <h1> {code} </h1>
        </div>
    )
}

export default Analyze;