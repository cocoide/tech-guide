"use client"

import CustomDialog from '@/app/_components/elements/CustomDialog'
import OutlineLoader from '@/app/_components/loaders/OutlineLoader'
import { articleAPI } from '@/app/_functions/article'
import { outlineDialogAtom } from '@/stores/dialog'
import { useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'

const OutlineDialog = () => {
    const [dialogState, _] = useAtom(outlineDialogAtom)
    const { data: overview, isLoading, isFetching } = useQuery({
        queryFn: async () => (await articleAPI.GetOverview(dialogState)).data,
        queryKey: [`article_outline.${dialogState}`],
        enabled: typeof dialogState !== 'boolean',
    })
    return (
        <CustomDialog layout='my-[100px] bg-white z-50 sm:mx-[15%] md:my-[80px] md:mx-[20%] lg:mx-[25%] sm:rounded-xl overflow-y-hidden'
            openAtom={outlineDialogAtom}
            content={
                <div className="relative flex flex-col  p-2 space-y-2 ">
                    <div className="text-gray-500">Outlines</div>
                    {isLoading && isFetching ?
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