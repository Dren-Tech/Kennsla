import { flaschenpost, Logger } from 'flaschenpost';

const getLogger = function (): Logger {
  return flaschenpost.getLogger();
};

export { getLogger };
