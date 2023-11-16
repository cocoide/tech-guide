import { api } from '@/app/_functions/API'
import { cookies } from 'next/headers'

export async function getSignupSession(sessionId: string | undefined) {
    "use server"

    return api.get<SignupSession>("/onboarding/session", "no-store", undefined, undefined, sessionId)
}

export async function getSessionID() {
    "use server"

    const store = cookies()
    return store.get("signup.sessionId")?.value
}

type SignupSession = {
    account_id: number;
    display_name: string;
    avatar_url: string;
    email: string;
    follow_topic_ids: number[] | null;
    onboarding_index: number;
};