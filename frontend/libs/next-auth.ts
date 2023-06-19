import { CookiesOptions } from 'next-auth';

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