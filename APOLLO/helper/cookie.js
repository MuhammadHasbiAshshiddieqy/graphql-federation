const {
    RemoteGraphQLDataSource
} = require('@apollo/gateway');

class CookieDataSource extends RemoteGraphQLDataSource {
  didReceiveResponse({ response, request, context }) {
    const cookie = request.http.headers.get('Cookie');
    if (cookie) {
      context.responseCookies.push(cookie);
    }

    // Return the response, even when unchanged.
    return response;
  }
}

module.exports = CookieDataSource;
