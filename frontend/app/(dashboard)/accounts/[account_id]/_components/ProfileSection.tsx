import { GetPublicProfile } from '../_functions/account'

const ProfileSection = ({ profile }: { profile: GetPublicProfile }) => {
    return (
        <div className='flex flex-col space-y-3 w-full items-center'>
            <div className="rounded-full overflow-hidden h-15 w-15 aspect-square flex items-center bg-cyan-200">
                {/* eslint-disable-next-line @next/next/no-img-element */}
                <img src={profile.avatar_url} alt={profile.display_name}
                    width={200} className='' />
            </div>
            <div className="text-md">{profile.display_name}</div>
        </div>
    )
}
export default ProfileSection