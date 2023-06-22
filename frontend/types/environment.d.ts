namespace NodeJS {
    interface ProcessEnv extends NodeJS.ProcessEnv {
      NEXTAUTH_GOOGLE_CLIENT_ID: string;
      NEXTAUTH_GOOGLE_CLIENT_SECRET: string;
      NEXT_API_BASE_URL: string;
    }
  }