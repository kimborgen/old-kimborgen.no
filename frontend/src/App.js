import React, { Component } from 'react';
import { NavLink } from './components/NavLink';
import { IndexLink } from 'react-router'
import { Menu } from 'semantic-ui-react'
import { Item } from 'semantic-ui-react'

class App extends Component {
  render() {
    return (
      <div className="App">
        <div>
          <Menu>
            <Item><IndexLink to="/" activeStyle={{color:"red"}}>Home</IndexLink></Item>
            <Item><NavLink to="/about">About</NavLink></Item>
            <Item><NavLink to="/articles">Articles</NavLink></Item>
            <Item><NavLink to="/login">Login</NavLink></Item>
          </Menu>
        </div>
        {this.props.children}
      </div>
    );
  }
}

export default App;
