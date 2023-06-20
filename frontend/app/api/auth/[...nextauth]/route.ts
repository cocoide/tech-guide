// https://remaster.com/blog/next-auth-jwt-session
import { api } from '@/app/_functions/API';
import { cookies } from '@/libs/next-auth';
import NextAuth, { User } from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';

const handler = NextAuth({
  session: {
    strategy: 'jwt',
    maxAge: 30 * 24 * 60 * 60, // 30 days
  },
  pages:{
    signIn: "/login"
  },
  cookies: cookies,
  secret: process.env.NEXTAUTH_SECRET,
  providers: [
    CredentialsProvider({
      name: 'Credentials',
      type: "credentials",
      credentials: {
        email: { label: "Email", type: "email", placeholder: "test@gmail.com" },
        password: { label: "Password", type: "password" }
      },
      async authorize(credentials, req):Promise<User | null> {
        const { data: user, ok } = await api.pos<User>("/login", credentials)
        if (ok && user) {
          return user
        }
        return null
      }
    })
  ],
  callbacks: {
    session: async ({ session, token }) => {
      session.token= token.token
      session.user.uid= token.uid
      return session
    },
    jwt: async ({ token, user }) => {
      if (user) {
        token.token = user.token
        token.uid= user.uid
      }
      return token
    },
  }
})

export { handler as GET, handler as POST };
