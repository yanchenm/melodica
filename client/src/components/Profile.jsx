import '../styles/Profile.css'

const Profile = () => {
    return (
        <div>
            {/* Header idk if we have to keep */}
            <div id="header">
                <h1>Spotify</h1>
                <p>Check your daily mood here.</p>
            </div>
            <div id="user-profile">
                {/*User Profile image, name, followers, following, playlists*/}
                <img id="profile-picture">{}</img>
                <h1 id="user-name">Name</h1>
                <div className="info">
                    <p className="stat-number">{0}</p>
                    <p className="stat-name">FOLLOWERS</p>
                </div>
                <div className="info">
                    <p className="stat-number">{0}</p>
                    <p className="stat-name">FOLLOWING</p>
                </div>
                <div className="info">
                    <p className="stat-number">{0}</p>
                    <p className="stat-name">PLAYLISTS</p>
                </div>
            </div>

        </div>
    ); 
}

export default Profile;