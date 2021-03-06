import '../styles/Main.css';

import * as qs from "query-string";

import { Redirect } from "react-router-dom";
import {useEffect} from "react";

const Main = () => {

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);

        if (params.get("code")) {
            return <Redirect to={{
                pathname: "/analyze",
                state: {code: params.get("code")}
            }} />
        }

        document.body.style.backgroundColor = "#181818";
    }, [])

    const getQueryString = () => {
        const url = "https://accounts.spotify.com/authorize?";

        const params = {
            "client_id": process.env.REACT_APP_SPOTIFY_CLIENT_ID,
            "response_type": "code",
            "redirect_uri": "https://melodica.tech/login",
            "scope": "user-read-email user-read-recently-played",
        };

        return url + qs.stringify(params);
    }

    return (
        <div id="main-container">
            <img id="logo" src="/brand.png" alt="logo"/>
            <div id="intro">
                <p>Curious to see how your music taste affects your mood?</p>
                <p>Find out here!</p>
            </div>
            <a href={getQueryString()} id="login-button">CONNECT WITH SPOTIFY</a>
        </div>
    )
}

export default Main;