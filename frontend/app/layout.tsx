import '../styles/globals.css'
import { Header } from './_components/layouts/Header'
import LoginDialog from './_components/layouts/LoginDialog'
import PostDialog from './_components/layouts/PostDialog'
import BottomNavigation from './_components/layouts/mobile/BottomNavigation'
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
            {children}
            <div className="fixed w-full bottom-0 md:hidden">
              <BottomNavigation />
            </div>
          </div>
          <LoginDialog />
          <PostDialog />
        </Providers>
      </body>
    </html>
  )
}
