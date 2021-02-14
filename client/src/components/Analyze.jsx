import {React, useEffect, useState}  from 'react';
import { BrowserRouter as Router, Switch, useLocation } from 'react-router-dom';
import Graph from "./Graph";
import {api} from '../api';
import "../styles/Analyze.css";

function Analyze() {

    const [user, setUser] = useState("");
    const [songData, setSongData] = useState([]);
    const [positivity, setPositivity] = useState(0);
    const [energy, setEnergy] = useState(0);

    useEffect(() => {
        const getUser = async () => {           
            const req = await api.get('/user');
            const data = await req.data;
            setUser(data.name.split(" ")[0]);
        }

        const getData = async () => {
            const req = await api.get('/recents');
            const data = await req.data;
            let songs = []
            data.songs.forEach(song => {
                const s = {
                    pos: song.Attribute.valence,
                    energy: song.Attribute.energy,
                    name: song.Name,
                    artist: song.Artists[0]
                };
                songs.push(s);
            })
            setSongData(songs);
            calculateNetPositivity();
        }
        const calculateNetPositivity = () => {
            let pos = 0;
            let energy = 0;
            for (let i = 0; i < songData.length; i++) {
                pos += songData[i].pos;
                energy += songData[i].energy;
            }
            setPositivity(Math.round(pos/songData.length * 100));
            setEnergy(Math.round(energy/songData.length * 100));
        }

        // const urlAddress = new URLSearchParams(props.location.search).get("code");
        // const a = urlAddress.get('code');

        getUser();
        getData();
    }, [positivity, energy])

    useEffect(() => {

    }, )
    
    return (
        <div id="analysis-container">
            <h1 id="analysis-header">
                Hi, {user}! Your average music positivity is <span className="positivity">{positivity}%</span>, and <br/> 
                your average musical energy is <span className="energy">{energy}%</span>. Hope you're doing okay!
            </h1>
            <div id="analysis-content-container">
                <Graph data={songData} />
                <div id="button-container">
                    <a id="recommend-button" href="../mood">FIND ME SOME MUSIC</a>
                    <a href="../" id="back-button">GO BACK</a>
                </div>
            </div>
        </div>
    )
}

export default Analyze;