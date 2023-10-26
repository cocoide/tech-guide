import { api } from '@/app/_functions/API';
import { AccountSession } from '@/types/model';
import { NextResponse } from 'next/server';
import { serverAuthFunc } from '../../_server_actions/auth';

export async function GET() {
    let response: AccountSession = {
        account_id: 0,
        display_name: '',
        avatar_url: '',
        features: []
    }
    try {
        const token = await serverAuthFunc.GetAccessToken()
        if (token) {
            throw new Error(`Error getting token`)
        }
        const { data: session, error } = await api.get<AccountSession>("/account/session", "no-store", token)
        if (!session || error) {
            throw new Error(`Error getting session: ${error}`)
        }
        response = session
    } catch (error) {
        console.error(error)
        return NextResponse.json(error, { status: 403 })
    }
    return NextResponse.json(response, { status: 200 })
}