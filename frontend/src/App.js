import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import { NavLink } from './components/NavLink';
import { IndexLink } from 'react-router'

class App extends Component {
  render() {
    return (
      <div className="App">
        <div>
          <h1>React Router Tutorial</h1>
          <ul role="nav">
            <li><IndexLink to="/" activeStyle={{color:"red"}}>Home</IndexLink></li>
            <li><NavLink to="/about">About</NavLink></li>
            <li><NavLink to="/articles">Articles</NavLink></li>
            <li><NavLink to="/login">Login</NavLink></li>
          </ul>
        </div>
        {this.props.children}
      </div>
    );
  }
}

export default App;
