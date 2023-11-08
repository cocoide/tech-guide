"use client"

import CustomDialog from '@/app/_components/elements/CustomDialog'
import OutlineLoader from '@/app/_components/loaders/OutlineLoader'
import { articleAPI } from '@/app/_functions/article'
import { outlineDialogAtom } from '@/stores/dialog'
import { XMarkIcon } from '@heroicons/react/24/outline'
import { useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'

const OutlineDialog = () => {
    const [dialogState, setDialogState] = useAtom(outlineDialogAtom)
    const { data: overview, isLoading, isRefetching } = useQuery({
        queryFn: async () => (await articleAPI.GetOverview(dialogState)).data,
        queryKey: [`article_outline.${dialogState}`],
        enabled: typeof dialogState !== 'boolean',
    })
    return (
        <CustomDialog layout='my-[100px] bg-white z-50 sm:mx-[15%] md:my-[80px] md:mx-[20%] lg:mx-[25%] sm:rounded-xl'
            openAtom={outlineDialogAtom}
            content={
                <div className="relative flex flex-col  p-2 space-y-2 overflow-y-scroll h-full w-full">
                    <button onClick={() => setDialogState(false)} className="absolute right-5 top-5 bg-gray-500 custom-badge p-2 rounded-full justify-center h-7 w-7">
                        <XMarkIcon className='h-5 w-5 text-gray-200' />
                    </button>
                    <div className="text-gray-500">Outlines</div>
                    {isLoading && isRefetching ?
                        <OutlineLoader />
                        :
                        overview &&
                        <ReactMarkdown remarkPlugins={[remarkGfm]} className='markdown'
                        >{overview}</ReactMarkdown>
                    }
                </div>
            }
        />
    )
}
export default OutlineDialog