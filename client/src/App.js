import Main from "./components/Main";
import Analyze from "./components/Analyze";
import { BrowserRouter as Router, Route, NavLink } from 'react-router-dom';
import Profile from "./components/Profile";

// "/" --> Main
// "/analyze" --> Analyze
function App() {
  return (
    <div>
        <Profile />
    </div>
  );
}

export default App;