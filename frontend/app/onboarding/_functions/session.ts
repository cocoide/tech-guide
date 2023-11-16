import { api } from '@/app/_functions/API'
import { cookies } from 'next/headers'

export function getSignupSession() {
    "use server"

    const store = cookies()
    const sessionId = store.get("signup.sessionId")?.value
    return api.get<SignupSession>("/onboarding/session", "no-store", undefined, undefined, sessionId)
}

type SignupSession = {
    onboarding_Index: number
}