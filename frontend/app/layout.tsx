import CircleLoader from '@/ui/CircleLoader'
import { Suspense } from 'react'
import '../styles/globals.css'
import CollectionDialog from './_components/layouts/CollectionDialog'
import FeedFilterDialog from './_components/layouts/FeedFilterDialog'
import { Header } from './_components/layouts/Header'
import LoginDialog from './_components/layouts/LoginDialog'
import PostDialog from './_components/layouts/PostDialog'
import LeftSideVar from './_components/layouts/desktop/LeftSideVar'
import BottomNavigation from './_components/layouts/mobile/BottomNavigation'
import Providers from './_components/providers'

export const metadata = {
  title: 'TechGuide',
  description: 'tech feed web app',
}
interface Props {
  children: React.ReactNode
  modal: React.ReactNode
}
export default function RootLayout({ children, modal }: Props) {
  return (
    <html lang="ja">
      <body className="bg-white dark:bg-black">
        <Providers>
          <div className="flex md:w-[770px] lg:w-[1050px] xl:w-[1300px]  mx-auto relative">
            <div className="sticky top-0 h-screen">
              <LeftSideVar />
            </div>
            <div className='flex flex-col w-full'>
              <div className="flex md:hidden">
              <Header />
            </div>
            <Suspense fallback={<CircleLoader />}>
                <div className="md:border-x-[0.5px] custom-border-color min-h-screen">
            {children}
                </div>
            </Suspense>
          </div>
          </div>
          <div className="fixed bottom-0 w-[100%] z-30 flex md:hidden">
            <BottomNavigation />
          </div>
          <CollectionDialog />
          <LoginDialog />
          <PostDialog />
          <FeedFilterDialog />
          {modal}
        </Providers>
      </body>
    </html>
  )
}
