import type { Metadata } from 'next';
import type { ReactNode } from 'react';

import MainLayout from '@/components/layout/MainLayout';

import AppProvider from './providers';

import '@/styles/globals.css';

export const metadata: Metadata = {
  title: 'Dreon Next.js Starter',
  description: 'Modern Dreon Next.js boilerplate with Ant Design and Tailwind CSS.',
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body>
        <AppProvider>
          <MainLayout>{children}</MainLayout>
        </AppProvider>
      </body>
    </html>
  );
}
