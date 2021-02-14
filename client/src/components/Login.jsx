import {api} from '../api';
import {useEffect} from "react";
import {useHistory} from "react-router-dom";
const Login = () => {
    let history = useHistory();

    
    useEffect(() => {
        const sendCode = async () => {
            const params = new URLSearchParams(window.location.search);
            const code = params.get("code");
            console.log(api.baseUrl);
            await api.post("/Login", {access_code: code, redirect_uri: "http://localhost:3000/login"});
    
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