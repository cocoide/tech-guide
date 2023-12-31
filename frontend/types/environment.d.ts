namespace NodeJS {
    interface ProcessEnv extends NodeJS.ProcessEnv {
      NEXTAUTH_GOOGLE_CLIENT_ID: string;
      NEXTAUTH_GOOGLE_CLIENT_SECRET: string;
      NEXT_PUBLIC_API_BASE_URL: string;
      NEXT_PUBLIC_FRON_URL: string;
      NEXTJWT_KEY: string;
      NEXT_PUBLIC_GTM_ID: string;
    }
  }