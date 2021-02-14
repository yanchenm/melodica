import '../styles/Mood.css';
import React, { useState } from 'react';
import Recommendations from './RecommendationList';

// rec = {album cover, song title, artist, link to spotify}
const recs = [
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
    {
        img: "https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Porter_Robinson_-_Worlds.jpg/220px-Porter_Robinson_-_Worlds.jpg",
        title: "Shelter",
        artist: "Porter Robinson"
    },
]

const Mood = () => {

    const [currentMood, setCurrentMood] = useState("");
    const [desiredMood, setDesiredMood] = useState("");
    const [recommendations, setRecommendations] = useState([]);

    const handleCurrentMoodChange = (e) => setCurrentMood(e.target.value);
    const handleDesiredMoodChange = (e) => setDesiredMood(e.target.value);

    const handleSubmit = (e) => {
        e.preventDefault();
        setRecommendations(recs);
    }

    const resetRecommendations = () => {
        setRecommendations([]);
    }

    if (recommendations.length === 0)
        return (
            <div id="mood-container">
                <div id="mood-text-container">
                    <p id="mood-intro">How are you feeling today?</p>
                    <form id="mood-prompt" onSubmit={handleSubmit}>
                        <p>
                            I'm 
                            <select name="current-mood" id="current-mood" defaultValue="" onChange={handleCurrentMoodChange}>
                                <option value="" disabled>(pick a mood)</option>
                                <option value="happy">happy</option>
                                <option value="sad">sad</option>
                                <option value="scared">scared</option>
                                <option value="stressed">stressed</option>
                                <option value="depressed">depressed</option>
                                <option value="excited">excited</option>
                                <option value="overwhelmed">overwhelmed</option>
                                <option value="relaxed">relaxed</option>
                            </select>
                            and I want to feel 
                            <select name="desired-mood" id="desired-mood" defaultValue="" onChange={handleDesiredMoodChange}>
                                <option value="" disabled>(pick a mood)</option>
                                <option value="happy">happy</option>
                                <option value="sad">sad</option>
                                <option value="scared">scared</option>
                                <option value="stressed">stressed</option>
                                <option value="depressed">depressed</option>
                                <option value="excited">excited</option>
                                <option value="overwhelmed">overwhelmed</option>
                                <option value="relaxed">relaxed</option>
                            </select>
                        </p>
                        <input type="submit" value="FIND ME SOME SONGS!" id="submit-button" />
                    </form>
                    <a href="../analyze" id="back-button">GO BACK</a>
                </div>
            </div>
        );
    else {
        return <Recommendations recommendations={recommendations} reset={resetRecommendations} />
    }
}

export default Mood;