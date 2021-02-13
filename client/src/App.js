import Main from "./components/Main";
import Analyze from "./components/Analyze";
import { BrowserRouter as Router, Route, NavLink } from 'react-router-dom';
import Profile from "./components/Profile";
import './App.css';

// "/" --> Main
// "/analyze" --> Analyze
function App() {
  return (
    <div>
        <Main />
    </div>
  );
}

export default App;