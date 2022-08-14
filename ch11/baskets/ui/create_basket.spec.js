const {PactV3, MatchersV3} = require('@pact-foundation/pact');
const chai = require('chai');
const expect = chai.expect;
const asPromised = require('chai-as-promised');
const axios = require("axios");

const {Client} = require('./client');

chai.use(asPromised);

describe('Baskets UI', () => {
  let provider;

  before(async () => {
    provider = new PactV3({
      consumer: 'baskets-ui',
      provider: 'baskets-api',
      logLevel: "warn",
//      dir: '../../pacts'
    });
  });

  context('calling startBasket', () => {
    describe('with a customer ID', () => {
      let basketId;

      before(() => {
        basketId = '3109004a-2b7f-4184-81da-daba46efe8fd';

        provider.given('I have been given a customer ID')
          .uponReceiving('a request to start a basket')
          .withRequest({
            method: 'POST',
            path: '/api/baskets',
            body: {
              customerId: 'customer-id'
            },
            headers: {Accept: 'application/json'},
          })
          .willRespondWith({
            body: MatchersV3.like({
              id: basketId
            }),
            headers: {'Content-Type': 'application/json'},
            status: 200,
          });
      });

      it('should start a new basket for the customer', () => {
        return provider.executeTest((mockServer) => {
          const client = new Client(mockServer.url);
          return client.startBasket('customer-id')
            .then((response) => {
              expect(response.data).to.have.property('id');
              expect(response.data.id).to.eq(basketId);
            });
        });
      });
    })
    describe('without any customer ID', () => {
      before(() => {
        provider.given('I have not been given a customer ID')
          .uponReceiving('a request to start a basket')
          .withRequest({
            method: 'POST',
            path: '/api/baskets',
            body: {
              customerId: ''
            },
            headers: {Accept: 'application/json'},
          })
          .willRespondWith({
            body: MatchersV3.like({
              "message": "the customer id cannot be blank",
            }),
            headers: {'Content-Type': 'application/json'},
            status: 400,
          });
      });

      it('should not start a new basket for the customer', async () => {
        return provider.executeTest((mockServer) => {
          const client = new Client(mockServer.url);
          return client.startBasket('')
            .catch(({response}) => {
              expect(response.status).to.eq(400);
            });
        });
      });
    })
  })

})
