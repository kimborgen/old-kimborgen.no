import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import './index.css';
import { Router, Route, browserHistory, IndexRoute } from 'react-router'
import { Articles } from "./components/Articles";
import { About } from "./components/About";
import { Login } from "./components/Login";
import { Article } from "./components/Article"
import { Home } from "./components/Home"

ReactDOM.render((
  <Router history={browserHistory}>
    <Route path="/" component={App}>
      <IndexRoute component={Home}/>
      <Route path="/about" component={About}/>
      <Route path="/articles" component={Articles}>
        <Route path="/articles/:articleId" component={Article}/>
      </Route>
      <Route path="/login" component={Login}/>
    </Route>
  </Router>
), document.getElementById('root'));

// browserHistory.push(path)
/*
ReactDOM.render(
  <App />,
  document.getElementById('root')
);
*/
