import redirects from '../../data/redirects.json';
import { IncomingMessage, ServerResponse } from 'http';

export default async (req: IncomingMessage, res: ServerResponse): Promise<void> => {
  // Find the redirect if it exists where the from === the requested url
  const redirect = redirects.find((item: any): boolean => item.from === req.url);

  // If it exists, redirect the page with a 301 response else carry on
  if (redirect) {
    res.writeHead(301, { Location: redirect.to });
    res.end();
  }
};
