import {useEffect} from "react";
import * as qs from "query-string";
import '../styles/Main.css';
import { Redirect } from "react-router-dom";

const Main = () => {

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);

        if (params.get("code")) {
            console.log(params.get("code"));
            return <Redirect to={{
                pathname: "/analyze",
                state: {code: params.get("code")}
            }} />
        }
    }, [])

    const getQueryString = () => {
        const url = "https://accounts.spotify.com/authorize?";

        const params = {
            "client_id": process.env.REACT_APP_SPOTIFY_CLIENT_ID,
            "response_type": "code",
            "redirect_uri": "http://localhost:3000/login",
            "scope": "user-read-email user-read-recently-played",
        };
        return url + qs.stringify(params);
    }

    return (
        <div id="main-container">
            <div id="intro">
                <p>Curious to see how your music tastes affects your mood?</p>
                <p>Find out here!</p>
            </div>
            <a href={getQueryString()} id="login-button">CONNECT WITH SPOTIFY</a>
        </div>
    )
}

export default Main;