import React from 'react'

export class Articles extends React.Component {
    render() {
        return (
        <div>
          <h1>Insert articles here</h1>
          {this.props.children}
        </div>
        );
    }
}
