import {React, useState, useEffect}  from 'react';
import { BrowserRouter as Router, useLocation, Switch } from 'react-router-dom';
import axios from "axios";
import Graph from './Graph';


function Analyze() {

    const [user, setUser] = useState("");
    const [songData, setSongData] = useState([]);

    useEffect(() => {
        const getUser = async () => {           
            const req = await axios.get("https://us-central1-musemood.cloudfunctions.net/User", {withCredentials: true});
            const data = await req.data;
            setUser(data.names.split(" ")[0]);
            console.log(data);
        }

        const getSongData = async () => {
            const req = await axios.get("https://us-central1-musemood.cloudfunctions.net/GetRecentlyPlayed", {withCredentials: true})
            const data = await req.data;
            console.log(data);
            setSongData(data);
        }
        // const urlAddress = new URLSearchParams(props.location.search).get("code");
        // const a = urlAddress.get('code');

        getUser();
        getSongData();
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