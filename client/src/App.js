import Main from './components/Main';
import Analyze from './components/Analyze';
import Mood from './components/Mood';
import { BrowserRouter as Router, Route, NavLink, Switch } from 'react-router-dom';
import './App.css';

// "/" --> Main
// "/analyze" --> Analyze
const App = () => {
  return (
      <div>
        <Switch>
          <Route path="/" exact component={Main} />
          <Route path="/analyze" exact component={Analyze} />
          <Route path="/mood" exact component={Mood} />
        </Switch>
      </div>
  );
}

export default App;