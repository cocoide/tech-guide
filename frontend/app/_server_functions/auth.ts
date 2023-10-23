import { cookies } from 'next/headers'

export const authServerFunc = {
    async GetAuth() {
        'use server'
        const cookieStore = cookies()
        const cookie = cookieStore.get('accessToken')
        return {
            token: cookie?.value,
        }
    }
}