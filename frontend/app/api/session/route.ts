import { api } from '@/app/_functions/API';
import { AccountSession } from '@/types/model';
import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function GET() {
    try {
        var accessToken = cookies().get("accessToken")?.value
        if (!accessToken) {
            throw new Error(`Error getting token`)
        }
        const { data: session, error } = await api.get<AccountSession>("/account/session", "no-store", accessToken)
        if (!session || error) {
            throw new Error(`Error getting session: ${error}`)
        }
        return NextResponse.json(session, { status: 200 })
    } catch (error) {
        console.error(error)
        return NextResponse.json(error, { status: 403 })
    }
}