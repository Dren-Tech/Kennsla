import express from 'express';
import { flaschenpost, getMiddleware } from 'flaschenpost';

type Status = 'running' | 'starting' | 'rebooting';

interface App {
  host: string;
  port: number;
  status: Status;
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const appConfig: App = {
  host: 'localhost',
  port: 8_080,
  status: 'running'
};

const logger = flaschenpost.getLogger();

logger.info('Booting API...');

const app = express();

// --------------
// Express config
// --------------
app.use(getMiddleware({ logOn: 'request' }));

app.get('/', (request, response): any => {
  response.send('Express + TypeScript Server');
});

logger.info('Start API of Kennlsa');

// Start server
app.listen(appConfig.port, appConfig.host, (): void => {
  logger.info(`⚡️[server]: Server is running at http://localhost:${appConfig.port}`);
  logger.info(`⚡️[server]: GraphQL endpoint is running at http://localhost:${appConfig.port}/graphql`);
});

// Const print = function (value: unknown): void {
//
// }
