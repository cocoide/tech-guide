import { api } from '@/app/_functions/API'
import { cookies } from 'next/headers'

export async function getSignupSession() {
    "use server"

    const store = cookies()
    const sessionId = store.get("signup.sessionId")?.value
    return api.get<SignupSession>("/onboarding/session", "no-store", undefined, undefined, sessionId)
}

type SignupSession = {
    account_id: number;
    display_name: string;
    avatar_url: string;
    email: string;
    follow_topic_ids: number[] | null;
    onboarding_index: number;
};