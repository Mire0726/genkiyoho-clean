// pages/_app.js または pages/_app.tsx
import React from 'react';
import { AppProps } from 'next/app';

function MyApp({ Component, pageProps }: AppProps) {
  // アプリレベルでの状態やレイアウトをここに配置できます
  return <Component {...pageProps} />;
}

export default MyApp;