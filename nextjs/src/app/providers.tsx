'use client';

import { AntdRegistry } from '@ant-design/nextjs-registry';
import { ConfigProvider } from 'antd';
import type { ReactNode } from 'react';
import { Toaster } from 'react-hot-toast';

import { TOAST_DEFAULT_OPTIONS } from '@/config/helpers/toast.helper';

export default function AppProvider({ children }: { children: ReactNode }) {
  return (
    <AntdRegistry>
      <ConfigProvider
        theme={{
          token: {
            borderRadius: 8,
          },
        }}
      >
        {children}
        <Toaster {...TOAST_DEFAULT_OPTIONS} />
      </ConfigProvider>
    </AntdRegistry>
  );
}
