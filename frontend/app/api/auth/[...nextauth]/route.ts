// https://remaster.com/blog/next-auth-jwt-session
import { authAPI } from '@/app/(authentication)/_function/auth';
import { cookies } from '@/libs/next-auth';
import NextAuth from 'next-auth';
import GoogleProvider from 'next-auth/providers/google';

const handler = NextAuth({
  session: {
    strategy: 'jwt',
    maxAge: 30 * 24 * 60 * 60,
  },
  pages: {
    signIn: "/login"
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
      session.token = token.token
      session.user.uid = token.uid
      return session
    },
    jwt: async ({ token, user }) => {
        if (user) {
        const isRegisterd = await authAPI.IsEmailUsed(user.email!)
        if (isRegisterd===false){
          const { data, ok } = await authAPI.SignUp({ "email": user.email!, "image":user.image!, "name": user.name! })
          if(!ok){
            throw new Error("failed to signup")
          }
          token.token = data?.token!
          token.uid = data?.uid!
          return token
      }
      const res = await authAPI.Login({ "email": user.email!})
      if (res===null){
        throw new Error("failed to login")
      }else{
        token.token = res?.token
        token.uid = res.uid
        user.name = res.name
        user.image=res.image
      }
      }
      return token
    },
  }
})

export { handler as GET, handler as POST };
