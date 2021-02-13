import React from 'react';
import * as qs from "query-string";
import '../styles/Main.css';

function Main () {

    const getQueryString = () => {
        const url = "https://accounts.spotify.com/authorize?";

        const params = {
            "client_id": process.env.REACT_APP_SPOTIFY_CLIENT_ID,
            "response_type": "code",
            "redirect_uri": "http://localhost:3000/analyze",
            "scope": "user-read-email user-read-recently-played",
        };

        console.log(qs.stringify(params));
        return url + qs.stringify(params);
    }

    return (
        <div id="main-container">
            <div id="intro">
                <p>Curious to see how your music tastes effects your mood?</p>
                <p>Find out here!</p>
            </div>
            <a href={getQueryString()} id="login-button">CONNECT WITH SPOTIFY</a>
        </div>
    )
}

export default Main;