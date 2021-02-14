import '../styles/RecommendationList.css';

const getRecommendationList = (recs) => {
    let recList = recs.map((rec) => {
        return (
            <li key={rec.url} className="song"> 
                <a href={rec.url} target="_blank">
                    <img className="song-cover" src={rec.img}></img>
                    <div className="song-info">
                        <p className="song-title" style={{ color: "black" }}>{rec.title}</p>
                        <p className="song-artist" style={{ color: "black" }}>{rec.artist}</p>
                    </div>
                </a>
            </li>
        );
    });
    return recList;
}

const Recommendations = (props) => {
    return  (
        <div id="recommendations-container">
            <p id="recommendations-intro">
                Based on your choices, <br/> here are some recommendations! <br/>
                <span id="small-text">Click on one to open it in Spotify.</span>
            </p>
            <button id="reset-button" onClick={props.reset}>I WANT NEW SONGS</button>
            <ul id="recommendations">
                {getRecommendationList(props.recommendations)}
            </ul>
        </div>
    );
}

export default Recommendations;