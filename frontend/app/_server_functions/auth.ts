import { VerifyJwt } from '@/utils/jwt';
import { cookies } from 'next/headers';

export const authServerFunc = {
    async GetAuth() {
        'use server'
        var token: string = ''
        var account_id: number = 0
        try {
        const cookieStore = cookies()
            const acceesTokenCookie = cookieStore.get("accessToken")
            if (acceesTokenCookie) {
                const verifyResponse = await VerifyJwt(acceesTokenCookie.value)
                if (verifyResponse?.updatedToken) {
                    token = verifyResponse.updatedToken
                }
                if (verifyResponse?.accountID) {
                    account_id = verifyResponse?.accountID
                }
            }
        } catch (e) {
            console.log(e)
        }
        return {
            token: token,
            account_id: account_id
        }
    }
}