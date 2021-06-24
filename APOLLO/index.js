const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
  serviceList: [
    { name: 'customer', url: 'http://localhost:4001/query' },
    { name: 'order', url: 'http://localhost:4002/query' },
    { name: 'user', url: 'http://localhost:4003/query' },
  ],
});

const server = new ApolloServer({
  gateway,
  subscriptions: false,
});

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});
