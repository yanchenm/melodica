import '../styles/RecommendationList.css';

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

const getRecommendationList = (recs) => {
    let recList = recs.map((rec) => {
        // need to fix keys and pass a unique id
        return (
            <li key={rec.title} className="song"> 
                <img className="song-cover" src={rec.img}></img>
                <div className="song-info">
                    <p className="song-title">{rec.title}</p>
                    <p className="song-artist">{rec.artist}</p>
                </div>
            </li>
        );
    });
    return recList;
}

const Recommendations = () => {
    return  (
        <div id="recommendations-container">
            <p id="recommendations-intro">
                Based on your choices, here are some recommendations!
            </p>
            <div id="recommendations">
                {getRecommendationList(recs)}
            </div>
        </div>
    );
}

export default Recommendations;