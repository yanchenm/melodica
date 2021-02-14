import Main from './components/Main';
import Analyze from './components/Analyze';
import Mood from './components/Mood';
import { Route, Switch } from 'react-router-dom';
import './App.css';
import axios from "axios";
import Login from './components/Login';
// "/" --> Main
// "/analyze" --> Analyze
const App = () => {
  axios.defaults.withCredentials = true;

  return (
      <div>
        <Switch>
          <Route path="/" exact component={Main} />
          <Route path="/analyze" exact component={Analyze} />
          <Route path="/mood" exact component={Mood} />
          <Route path="/login" exact component={Login} />
        </Switch>
      </div>
  );
}

export default App;