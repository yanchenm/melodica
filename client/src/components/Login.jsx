import {api} from '../api';
import {useEffect} from "react";
import {useHistory} from "react-router-dom";
const Login = () => {
    let history = useHistory();

    
    useEffect(() => {
        const sendCode = async () => {
            const params = new URLSearchParams(window.location.search);
            const code = params.get("code");
            await api.post("/login", {access_code: code, redirect_uri: "https://melodica.tech/login"});
    
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