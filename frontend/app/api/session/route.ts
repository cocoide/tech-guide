import { api } from '@/app/_functions/API';
import { AccountSession } from '@/types/model';
import { NextResponse } from 'next/server';
import { serverAuthFunc } from '../../_server_actions/auth';

export async function GET() {
    const token = await serverAuthFunc.GetAccessToken()
    if (!token) {
        return NextResponse.json({ status: 403 })
    }
    const { data: session, ok } = await api.get<AccountSession>("/account/session", "no-store", token)
    if (!ok) {
        return NextResponse.json({ status: 403 })
    }
    return NextResponse.json(session, { status: 200 })
}