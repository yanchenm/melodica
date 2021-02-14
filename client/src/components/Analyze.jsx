import {React, useEffect, useState}  from 'react';
import { BrowserRouter as Router, Switch, useLocation } from 'react-router-dom';
import Graph from "./Graph";
import {api} from '../api';

function Analyze() {

    const [user, setUser] = useState("");
    const [songData, setSongData] = useState([]);

    useEffect(() => {
        const getUser = async () => {           
            const req = await api.get('/User');
            const data = await req.data;
            setUser(data.name);
            console.log(data);
        }

        const getData = async () => {
            const req = await api.get('/GetRecentlyPlayed');
            const data = await req.data;
            console.log(data);
        }

        // const urlAddress = new URLSearchParams(props.location.search).get("code");
        // const a = urlAddress.get('code');

        getUser();
        getData();
    }, [])

    useEffect(() => {

    }, )
    const calculateNetPositivity = () => {
        let pos = 0;

        for (let i = 0; i < songData.length; i++) {
            const valence = songData[i].valence;
            pos += valence;
        }

        return Math.round(pos/songData.length);
    }
    return (
        <div>
            <h1> Hi, {user}! {calculateNetPositivity() > 50? "HELLO" : "NOT HELLO"}</h1>
            <Graph
                data={songData}
            />
        </div>
    )
}

export default Analyze;