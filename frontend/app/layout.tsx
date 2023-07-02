import CircleLoader from '@/ui/CircleLoader'
import { Suspense } from 'react'
import '../styles/globals.css'
import CollectionDialog from './_components/layouts/CollectionDialog'
import { Header } from './_components/layouts/Header'
import LoginDialog from './_components/layouts/LoginDialog'
import PostDialog from './_components/layouts/PostDialog'
import Providers from './_providers'

export const metadata = {
  title: 'Tech Guide',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <body >
        <Providers>
          <div className='flex flex-col'>
            <div className="sticky top-0 z-30">
              <Header />
            </div>
            <Suspense fallback={<CircleLoader />}>
            {children}
            </Suspense>
          </div>
          <CollectionDialog />
          <LoginDialog />
          <PostDialog />
        </Providers>
      </body>
    </html>
  )
}
