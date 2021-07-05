require('dotenv').config()
const { ApolloServer } = require('apollo-server');
const { ApolloGateway, RemoteGraphQLDataSource } = require("@apollo/gateway");
const request = require('request-promise');

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

async function getData(req) {
  var payload = {
      uri: `${process.env.USER_URL}/dashboard_user`,
      headers: {
          'Content-Type': 'application/json',
          'Authorization': req.headers.authorization,
          'User-Agent': 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/38.0.2125.111 Safari/537.36'
      }
  };
  return await request.get(payload)
      .then(async function (data) {
          var data = JSON.parse(data);
          if(data.data.dashboard_user !== 'null'){
              return {datauser:data.data.dashboard_user}
          } else if (data.data.channel_user !== 'null') {
              return {datauser:data.data.dashboard_user}
          } else if (data.data.admin_user !== 'null') {
              return {datauser:data.data.admin_user}
          } else {
              throw new Error(`payload: ${JSON.stringify(payload)}, response: ${JSON.stringify(data)}`)
          }
      })
      .catch(function (err) {
          throw new Error(err.message)
      });
}

const server = new ApolloServer({
  gateway,
  subscriptions: false,
  context: ({ req }) => {
            return getData(req)
            .then(async function (data) {
              profileId = `{\"profile_id\": ${data.datauser.profile_id}}`
              return { profileId }
            })
          },
});

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});
