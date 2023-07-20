import AccountTopTab from './_components/AccountTopTab'
import ProfileSection from './_components/ProfileSection'
import { accountAPI } from './_functions/account'

interface Props extends AccountParams {
    children: React.ReactNode
}
export default async function AccountLayout({ children, params }: Props) {
    const { account_id } = params
    const { data: profile } = await accountAPI.getProfile(Number(account_id))
    return (
        <div className="flex flex-col w-full space-y-3 md:mt-5">
            {profile &&
                <ProfileSection profile={profile} />
            }
            <AccountTopTab account_id={account_id} />
            {children}
        </div>
    )
}