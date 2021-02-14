import '../styles/Mood.css';

const Mood = () => {
    return (
        <div id="mood-container">
            <div id="mood-text-container">
                <p id="mood-intro">How are you feeling today?</p>
                <form id="mood-prompt">
                    <p>
                        I'm 
                        <select name="current-mood" id="current-mood" defaultValue="">
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
                        <select name="desired-mood" id="desired-mood" defaultValue="">
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
                </form>
            </div>
        </div>
    );
}

export default Mood;