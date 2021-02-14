import '../styles/RecommendationList.css';

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

const Recommendations = (props) => {
    return  (
        <div id="recommendations-container">
            <p id="recommendations-intro">
                Based on your choices, <br/> here are some recommendations!
            </p>
            <button id="reset-button" onClick={props.reset}>I WANT NEW SONGS</button>
            <ul id="recommendations">
                {getRecommendationList(props.recommendations)}
            </ul>
        </div>
    );
}

export default Recommendations;