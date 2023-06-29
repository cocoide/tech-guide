import "next-auth";

declare module "next-auth" {
  interface User{
    uid: number;
    token: string;
  }
  interface Session {
    user: {
      uid: number
      name: string
      image: string
    },
    token: string
    token_expires: number
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    token: string
    token_expires: number
    uid: number
  }
}
