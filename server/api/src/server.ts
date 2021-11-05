import express from 'express';
import { buildSchema } from 'graphql';
import { graphqlHTTP } from 'express-graphql';
import { getLogger } from '@kennsla/common';
import { getMiddleware } from 'flaschenpost';

type Status = 'running' | 'starting' | 'rebooting';

interface App {
  host: string;
  port: number;
  status: Status;
}

const appConfig: App = {
  host: 'localhost',
  port: 8_080,
  status: 'running'
};

const logger = getLogger();

logger.info('Booting API...');

const app = express();

// --------------
// Express config
// --------------
app.use(getMiddleware({ logOn: 'request' }));

app.get('/', (request, response): any => {
  response.send('Express + TypeScript Server');
});

// --------------
// GraphQL config
// --------------

// Construct a schema, using GraphQL schema language
const schema = buildSchema(`
  type Query {
    hello: String
  }
`);

// The root provides a resolver function for each API endpoint
const root = {
  hello: 'Hello world!'
};

app.use('/graphql', graphqlHTTP({
  schema,
  rootValue: root,
  graphiql: true
}));

logger.info('Start API of Kennlsa');

// Start server
app.listen(appConfig.port, appConfig.host, (): void => {
  logger.info(`⚡️[server]: Server is running at http://localhost:${appConfig.port}`);
  logger.info(`⚡️[server]: GraphQL endpoint is running at http://localhost:${appConfig.port}/graphql`);
});

// Const print = function (value: unknown): void {
//
// }
