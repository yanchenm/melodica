import React from 'react';
import * as qs from "query-string";

function Main() {

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
        <div className="App">
            <header className="App-header">
                <p>
                Edit <code>src/App.js</code> and save to reload.
                </p>
                <button> 
                    <a href={getQueryString()}>
                        CONNECT WITH SPOTIFY 
                    </a>
                    </button>
            </header>
        </div>
    )
}

export default Main;