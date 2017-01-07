import React from 'react'

import { Input, Button, Form } from 'semantic-ui-react'

export class Login extends React.Component {
    constructor(props) {
      super(props);
      this.state = {username: '', password: '' , reply:''};

      this.handleChangeUsername = this.handleChangeUsername.bind(this);
      this.handleChangePassword = this.handleChangePassword.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }


    handleChangeUsername(event) {
      this.setState({username: event.target.value});
    }
    handleChangePassword(event) {
      this.setState({password: event.target.value});
    }

    handleSubmit(event) {
      event.preventDefault();
      var request = new Request('http://146.185.153.19:8090/api/login', {
      	method: 'POST',
        mode: 'no-cors',
      	headers: new Headers({
      		'Content-Type': 'application/json'
      	}),
        body: JSON.stringify({
          username: this.state.username,
          password: this.state.password
        })
      });
      var rep = '';

      fetch(request).then(function(response) {
        console.log(response.ok)
        console.log(response.body)
        console.log(response.status)
        return response.json();
      }).then(function(response) {
        if (response.ok === true) {
          rep = "logged in!";
          localStorage.setItem( 'token', response);
        } else {
          rep = response.body;
        }
      }).catch(function(err) {
	       // Error :(
         console.log(err);
      });
      this.setState({reply: rep});
      console.log(request)
    }
    render() {
      return (
        <div>
          <p>{this.state.reply}</p>
          <Form>
            <Form.Field>
              <label>Username</label>
              <Input fluid placeholder='username' value={this.state.username} onChange={this.handleChangeUsername}/>
            </Form.Field>
            <Form.Field>
              <label>Password</label>
              <Input fluid type="password" placeholder='password' value={this.state.password} onChange={this.handleChangePassword} />
            </Form.Field>
            <Button fluid type='submit' value='submit' onClick={this.handleSubmit}>Submit</Button>
          </Form>
        </div>
      );

    }
}

/*
fetch('https://davidwalsh.name/demo/arsenal.json').then(function(response) {
	// Convert to JSON
	return response.json();
}).then(function(j) {
	// Yay, `j` is a JavaScript object
	console.log(j);
});
*/
/*
var request = new Request('https://davidwalsh.name/users.json', {
	method: 'POST',
	mode: 'cors',
	redirect: 'follow',
	headers: new Headers({
		'Content-Type': 'text/plain'
	})
});

// Now use it!
fetch(request).then(function() { /* handle response */ //});
