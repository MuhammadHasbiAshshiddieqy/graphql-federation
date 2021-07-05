const {
    RemoteGraphQLDataSource
} = require('@apollo/gateway');

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
  willSendRequest({ request, context }) {
    request.http.headers.set('Content-Type', 'application/json');
    request.http.headers.set('Authorization', request.headers.authorization);
    request.http.headers.set('User-Agent', 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36');
  }
}

module.exports = AuthenticatedDataSource;
