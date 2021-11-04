import { flaschenpost, Logger } from 'flaschenpost';

function getLogger(): Logger {
    return flaschenpost.getLogger();
}

export { getLogger }
