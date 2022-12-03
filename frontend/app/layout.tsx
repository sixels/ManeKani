import '@/styles/dist.css';
import React from 'react';
import Footer from '@/ui/Footer';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html>
      <head>
        <title>ManeKani</title>
      </head>
      <body className="bg-gray-100">
        <div className="app">{children}</div>
        <Footer />
      </body>
    </html>
  );
}
