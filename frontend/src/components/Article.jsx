import React from 'react'

export class Article extends React.Component {
    render() {
        return (
          <h1>Hello, article {this.props.params.articleId}!</h1>
          );
    }
}
