const { ApolloServer } = require('apollo-server');
const { ApolloGateway, RemoteGraphQLDataSource } = require("@apollo/gateway");

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
  willSendRequest({ request, context }) {
    request.http.headers.set('context', context.profileId);
  }
}

const gateway = new ApolloGateway({
  serviceList: [
    { name: 'customer', url: 'http://localhost:4001/query' },
    { name: 'order', url: 'http://localhost:4002/query' },
    { name: 'user', url: 'http://localhost:4003/query' },
    // { name: 'customer', url: 'https://customer-gql.forstok.com/query' },
    // { name: 'order', url: 'https://order-gql.forstok.com/query' },
    // { name: 'user', url: 'https://user-gql.forstok.com/query' },
    ],
  buildService({ name, url }) {
    return new AuthenticatedDataSource({ url });
    },
});

const server = new ApolloServer({
  gateway,
  subscriptions: false,
  context: ({ req }) => {
            // // Get the user token from the headers
            // const token = req.headers.authorization || '';
            // // Try to retrieve a user with the token
            // const userId = getUserId(token);
            // // Add the user ID to the context
            // return { userId };
            const profileId = "{\"profile_id\": 186}"
            return { profileId }
          },
});

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});
