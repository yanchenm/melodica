import Main from "./components/Main";
import Analyze from "./components/Analyze";
import { BrowserRouter as Router, Route, NavLink, Switch } from 'react-router-dom';
import Profile from "./components/Profile";
import './App.css';

// "/" --> Main
// "/analyze" --> Analyze
function App() {
  return (
      <div>
        <Switch>
          <Route path="/" exact component={Main} />
          <Route path="/analyze" exact component={Analyze} />
          <Route path="/profile" exact component={Profile} />
        </Switch>
      </div>
  );
}

export default App;