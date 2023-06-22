'use client'
import { ArrowRightOnRectangleIcon, Cog8ToothIcon, DocumentDuplicateIcon, PlusCircleIcon } from '@heroicons/react/24/outline'
import { signOut } from 'next-auth/react'
import Image from 'next/image'
import CustomDropDown, { Menue } from '../elements/CustomDropDown'

interface Props {
    uid: number
    img: string
}
const menues: Menue[] = [
    { icon: <PlusCircleIcon className='h-5 w-5 text-slate-400' />, label: "投稿する" },
    { icon: <DocumentDuplicateIcon className='h-5 w-5 text-slate-400' />, label: "保存リスト" },
    { icon: <Cog8ToothIcon className='h-5 w-5 text-slate-400' />, label: "アカウント設定" },
]
const UserDropDown = ({ img }: Props) => {
    return (
        <CustomDropDown
            menues={menues}
            button={
                <div className="rounded-full w-8 h-8 overflow-hidden animate-appear
                shadow-[0_8px_30px_rgb(0,0,0,0.12)]">
                    <Image src={img} alt='user_avatar' width={50} height={50} />
                </div>
            }
            option={
                <button onClick={() => signOut()}
                    className="text-sm border-t px-1 pt-2 flex flex-row items-center space-x-2"
                > <ArrowRightOnRectangleIcon className='h-5 w-5 text-slate-400' />,
                    <div className="">ログアウト</div>
                </button>
            }
        />
    )
}
export default UserDropDown