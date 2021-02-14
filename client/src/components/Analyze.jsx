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
            // const req = await axios.get("https://us-central1-musemood.cloudfunctions.net/GetRecentlyPlayed", {withCredentials: true})
            // const data = await req.data;
            const data = [
                {x: 0.1, y: 0.2, title: "Omega"},
                {x: 0.6, y: 0.5, title: "Omeg"},
                {x: -0.4, y: 0.6, title: "Ome"},
                {x: 0.25, y: 0.2, title: "Om"},
                {x: 0.89, y: -0.7, title: "O"},
                {x: -0.44, y: -0.2, title: "OmegaO"},

            ]
            setSongData(data);
        }
        // const urlAddress = new URLSearchParams(props.location.search).get("code");
        // const a = urlAddress.get('code');

        getUser();
        getSongData();
    }, [])
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