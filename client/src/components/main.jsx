import React from 'react'
import { BrowserRouter as Router, Route, NavLink } from 'react-router-dom'
import axios from "axios"

function Main() {

    const onSpotifyClick = () => {
        axios.get("https://accounts.spotify.com/authorize")
    }
    return (
        <div className="App">
            <header className="App-header">
                <p>
                Edit <code>src/App.js</code> and save to reload.
                </p>
                <a
                className="App-link"
                href="https://reactjs.org"
                target="_blank"
                rel="noopener noreferrer"
                >
                
                </a>
            </header>
        </div>
    )
}

export default Main;