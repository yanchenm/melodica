import {useEffect} from "react";
import axios from "axios";
import {useHistory} from "react-router-dom";
const Login = () => {
    let history = useHistory();

    
    useEffect(() => {
        const sendCode = async () => {
            const params = new URLSearchParams(window.location.search);
            const code = params.get("code");
            await axios.post("https://us-central1-musemood.cloudfunctions.net/Login", {access_code: code, redirect_uri: "http://localhost:3000/login"}, );
    
            history.push("/analyze");
        }
        sendCode();
    }, [])

    return (
        <div>

        </div>
    )
}

export default Login;