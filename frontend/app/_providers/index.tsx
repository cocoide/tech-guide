'use client'
import { Provider } from 'jotai';
import { SessionProvider } from 'next-auth/react';
import { ReactNode } from 'react';
import { Toaster } from 'react-hot-toast';

export default function Providers({ children }: { children: ReactNode }) {
    return (
        <>
            <Provider>
            <SessionProvider>
                <Toaster />
                {children}
            </SessionProvider>
            </Provider>
        </>
    )
}