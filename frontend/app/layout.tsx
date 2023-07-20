import CircleLoader from '@/ui/CircleLoader'
import { Suspense } from 'react'
import '../styles/globals.css'
import CollectionDialog from './_components/layouts/CollectionDialog'
import { Header } from './_components/layouts/Header'
import LoginDialog from './_components/layouts/LoginDialog'
import PostDialog from './_components/layouts/PostDialog'
import LeftSideVar from './_components/layouts/mobile/LeftSideVar'
import Providers from './_components/providers'

export const metadata = {
  title: 'Tech Guide',
  description: 'Generated by create next app',
}
interface Props {
  children: React.ReactNode
}
export default function RootLayout({ children }: Props) {
  return (
    <html lang="ja">
      <body >
        <Providers>
          <div className="flex justify-center md:w-[770px] lg:w-[1100px] xl:w-[1200px]  mx-auto relative">
            <div className="sticky top-0 h-screen">
              <LeftSideVar />
            </div>
            <div className='flex flex-col w-full'>
              <div className="sticky top-0 z-30 flex md:hidden">
              <Header />
            </div>
            <Suspense fallback={<CircleLoader />}>
                <div className="md:border-x-[0.5px] min-h-screen">
            {children}
                </div>
            </Suspense>
          </div>
          </div>
          <CollectionDialog />
          <LoginDialog />
          <PostDialog />
        </Providers>
      </body>
    </html>
  )
}
