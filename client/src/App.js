import Main from "./components/main";
import Analyze from "./components/analyze";
import { BrowserRouter as Router, Route, NavLink } from 'react-router-dom';

// "/" --> Main
// "/analyze" --> Analyze
function App() {
  return (
    <Main/>
  );
}

export default App;
