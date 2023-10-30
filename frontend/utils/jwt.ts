import { api } from '@/app/_functions/API';
import jwt, { VerifyOptions } from 'jsonwebtoken';

export async function decodeJwt(token: string): Promise<CustomClaims> {
    "use server"

    const option: VerifyOptions = {
        algorithms: ['HS256'],
    }
    const decoded = jwt.decode(token, option)
    if (decoded === null || typeof decoded === 'string' || !decoded["exp"] || !decoded["account_id"]) {
        throw new Error("Failed to decode accessToken")
    }
    return {
        account_id: decoded["account_id"],
        exp: decoded["exp"],
    }
}

type CustomClaims = {
    account_id: number,
    exp: number,
}

export async function refreshAccessToken(refreshToken: string) {
    "use server"

    try {
        const params = { "token": refreshToken }
        const { data: accessToken, error } = await api.pos<string>("/oauth/refresh", undefined, undefined, params)
        if (error || !accessToken) {
            throw new Error(`Failed to refresh token: ${error}`)
        }
        return accessToken
    } catch (error) {
        console.log(`Failed to refresh token: ${error}`)
        throw error
    }
}