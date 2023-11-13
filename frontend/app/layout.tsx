import CircleLoader from '@/ui/CircleLoader'
import { Suspense } from 'react'
import '../styles/globals.css'
import CollectionDialog from './_components/layouts/CollectionDialog'
import CommentDialog from './_components/layouts/CommentDialog'
import FeedFilterDialog from './_components/layouts/FeedFilterDialog'
import GoogleAnalytics from './_components/layouts/GoogleAnalitics'
import { Header } from './_components/layouts/Header'
import LoginDialog from './_components/layouts/LoginDialog'
import PostDialog from './_components/layouts/PostDialog'
import PreviewDialog from './_components/layouts/PreviewDialog'
import LeftSideVar from './_components/layouts/desktop/LeftSideVar'
import Providers from './_components/providers'

export const metadata = {
  title: 'TechGuide',
  description: 'tech feed web app',
}
interface Props {
  children: React.ReactNode
}
export default function RootLayout({ children }: Props) {
  return (
    <html lang="ja">
      <body className="bg-white dark:bg-black">
        <Suspense>
          <GoogleAnalytics />
        </Suspense>
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
          <CommentDialog />
          <CollectionDialog />
          <LoginDialog />
          <PostDialog />
          <FeedFilterDialog />
          <PreviewDialog />
        </Providers>
      </body>
    </html>
  )
}
