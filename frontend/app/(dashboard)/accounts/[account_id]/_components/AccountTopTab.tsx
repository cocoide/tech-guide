'use client'
import { useAuth } from '@/hooks/useAuth';
import { clsx } from '@/utils/clsx';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

export default function AccountTopTab({ account_id }: { account_id: string }) {
    function checkAccountPathname(directory: string, pathname: string): boolean {
        const regex = new RegExp(`^/accounts/\\d+${directory}$`);
        return regex.test(pathname);
    }
    const { user } = useAuth()
    const pathname = usePathname()
    return (
        <div className="text-slate-400 w-full justify-center flex flex-row space-x-10 items-center p-1">
            <Link href={`/accounts/${account_id}`} className={clsx("p-1", checkAccountPathname("", pathname) ? "text-slate-600 border-b border-cyan-300" : "")}>投稿</Link>
            <Link href={`/accounts/${account_id}/collections`} className={clsx("p-1", checkAccountPathname("/collections", pathname) ? "text-slate-600 border-b border-cyan-300" : "")}>保存</Link>
            {
                user.uid === Number(account_id) &&
            <Link href={`/accounts/${account_id}/histories`} className={clsx("p-1", checkAccountPathname("/histories", pathname) ? "text-slate-600 border-b border-cyan-300" : "")}>履歴</Link>
            }
        </div >
    )
}


