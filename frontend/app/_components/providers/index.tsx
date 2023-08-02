'use client'
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Provider } from 'jotai';
import { SessionProvider } from 'next-auth/react';
import { ReactNode } from 'react';
import { Toaster } from 'react-hot-toast';

export default function Providers({ children }: { children: ReactNode }) {
    const queryClient = new QueryClient();
    return (
        <>
            <Provider>
                <SessionProvider>
                    <QueryClientProvider client={queryClient}>
                    <Toaster />
                    {children}
                    </QueryClientProvider>
                </SessionProvider>
            </Provider>
        </>
    )
}