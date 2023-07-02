import { authAPI } from '@/app/_functions/auth';
import { CookiesOptions, NextAuthOptions } from 'next-auth';
import GoogleProvider from 'next-auth/providers/google';

export const cookies: Partial<CookiesOptions> = {
    sessionToken: {
        name: `next-auth.session-token`,
        options: {
            httpOnly: true,
            sameSite: "none",
            path: "/",
            domain: process.env.NEXT_PUBLIC_DOMAIN,
            secure: true,
        },
    },
    callbackUrl:{
        name: `next-auth.callback-url`,
        options: {
        },
    },
    csrfToken:{
        name: "next-auth.csrf-token",
        options:{
        },
    },
}

export const authOptions: NextAuthOptions = 
  {
    session: {
      strategy: 'jwt',
      maxAge: 30 * 24 * 60 * 60,
    },
    cookies: cookies,
    secret: process.env.NEXTAUTH_SECRET,
    providers: [
      GoogleProvider({
        clientId: process.env.NEXTAUTH_GOOGLE_CLIENT_ID,
        clientSecret: process.env.NEXTAUTH_GOOGLE_CLIENT_SECRET,
      },
      ),
    ],
    callbacks: {
      session: async ({ session, token }) => {
        if (Date.now() < token.token_expires * 1000) {
          session.token = token.token
        } else {
          const { data } = await authAPI.RefreshToken(token.token)
          if (!data?.token && !data?.token_expires) {
            throw new Error("failed to refresh token")
          }
          session.token = data.token
          session.token_expires = data.token_expires
        }
        session.user.uid = token.uid
        return session
      },
      jwt: async ({ token, user }) => {
        if (user) {
          const isRegisterd = await authAPI.IsEmailUsed(user.email!)
          if (isRegisterd === false) {
            const { data, ok } = await authAPI.SignUp({ "email": user.email!, "image": user.image!, "name": user.name! })
            if (!ok || !data?.token || !data.token_expires) {
              throw new Error("failed to signup")
            }
            token.token = data.token!
            token.token_expires = data.token_expires
            token.uid = data.uid!
            return token
          }
          const { ok, data } = await authAPI.Login({ "email": user.email! })
          if (!ok || !data?.token || !data.token_expires) {
            throw new Error("failed to login")
          } else {
            token.token = data?.token
            token.token_expires = data.token_expires
            token.uid = data.uid
            user.name = data.name
            user.image = data.image
          }
        }
        return token
      },
    }
}

