import Main from './components/Main';
import Analyze from './components/Analyze';
import Mood from './components/Mood';
import Recommendations from './components/RecommendationList';
import { BrowserRouter as Router, Route, NavLink, Switch } from 'react-router-dom';
import Profile from "./components/Profile";
import './App.css';

// "/" --> Main
// "/analyze" --> Analyze
const App = () => {
  return (
      <div>
        <Switch>
          <Route path="/" exact component={Main} />
          <Route path="/analyze" exact component={Analyze} />
          <Route path="/profile" exact component={Profile} />
          <Route path="/mood" exact component={Mood} />
          <Route path="/recommendations" exact component={Recommendations} />
        </Switch>
      </div>
  );
}

export default App;