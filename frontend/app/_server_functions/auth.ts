import jwt, { VerifyOptions } from 'jsonwebtoken';
import { cookies } from 'next/headers';

export const authServerFunc = {
    async GetAuth() {
        'use server'
        var token: string = ''
        var account_id: number = 0
        try {
        const cookieStore = cookies()
        const cookie = cookieStore.get('accessToken')
            if (cookie?.value) {
                token = cookie.value
                const claims = verifyToken(token)
                if (claims) {
                    account_id = claims
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


function verifyToken(token: string): number | null {
    const option: VerifyOptions = {
        algorithms: ['HS256'],
    }
    const decoded = jwt.decode(token, option)
    if (decoded === null || typeof decoded === 'string' || !decoded.exp) {
        return null
    }
    if (Date.now() < decoded.exp * 1000) {
        // Add refresh token later
        console.log("expire")
    }
    const account_id = decoded["account_id"] as number
    if (!account_id) {
        return null
    }
    return account_id
}