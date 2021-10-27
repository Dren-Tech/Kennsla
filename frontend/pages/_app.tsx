import { AppProps } from 'next/app';
import React from 'react';
import '../styles/globals.css';
import 'tailwindcss/tailwind.css';

const MyApp = function ({ Component, pageProps }: AppProps): React.ReactElement {
  return <Component { ...pageProps } />;
};

export default MyApp;
