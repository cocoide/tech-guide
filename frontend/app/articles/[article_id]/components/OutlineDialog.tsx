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
    const { data: overview, isLoading, isRefetching, isFetching } = useQuery({
        queryFn: async () => (await articleAPI.GetOverview(dialogState)).data,
        queryKey: [`article_outline.${dialogState}`],
        enabled: typeof dialogState !== 'boolean',
    })
    if (!isLoading && !overview) {
        return null
    }
    return (
        <CustomDialog layout='mx-[20px] my-[150px] bg-white z-50 sm:mx-[15%] md:my-[80px] md:mx-[20%] lg:mx-[25%] rounded-xl'
            openAtom={outlineDialogAtom}
            content={
                <div className="relative flex flex-col  p-2 space-y-2 overflow-y-scroll h-full w-full">
                    <button onClick={() => setDialogState(false)} className="absolute right-3 top-3 bg-gray-400 custom-badge p-2 rounded-full justify-center h-7 w-7">
                        <XMarkIcon className='h-5 w-5 text-gray-200' />
                    </button>
                    <div className="text-gray-500">Outlines</div>
                    <div className="relative flex flex-col  p-5 custom-border rounded-xl space-y-2">
                        {isLoading || isRefetching || isFetching ?
                        <OutlineLoader />
                        :
                        overview &&
                        <ReactMarkdown remarkPlugins={[remarkGfm]} className='markdown'
                        >{overview}</ReactMarkdown>
                    }
                    </div>
                </div>
            }
        />
    )
}
export default OutlineDialog