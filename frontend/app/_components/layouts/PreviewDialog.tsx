"use client"
import YouTubeEmbed from '@/app/(home)/_components/YoutubeEmbed'
import SpeakerDeckEmbed from "@/app/_components/layouts/SpeakerDeckEmbed"
import SpeakerDeckLoader from "@/app/_components/loaders/SpeakerDeckLoader"
import { previewDialogAtom } from '@/stores/dialog'
import { useAtom } from 'jotai'
import { Suspense } from 'react'
import CustomDialog from '../elements/CustomDialog'
const PreviewDialog = () => {
    const [previewDialog, _] = useAtom(previewDialogAtom)

    var summary, youtube_id, speakerdeck_id: string | undefined
    if (typeof previewDialog !== "boolean") {
        summary = previewDialog?.summary
        youtube_id = previewDialog?.youtube_id
        speakerdeck_id = previewDialog?.speakerdeck_id
    }
    return (
        <CustomDialog
            openAtom={previewDialogAtom}
            layout='my-[150px] bg-white z-50 sm:mx-[15%] md:my-[100px] md:mx-[20%] lg:mx-[25%] sm:rounded-xl'
            content={
                <div className='p-5 md:p-10 flex flex-col space-y-3 text-gray-600 dark:text-gray-100'>
                    {speakerdeck_id &&
                        <Suspense fallback={<SpeakerDeckLoader />}>
                            <SpeakerDeckEmbed url={speakerdeck_id} />
                        </Suspense>
                    }
                    {youtube_id &&
                        <YouTubeEmbed youtube_id={youtube_id} />
                    }
                    {!speakerdeck_id && !youtube_id &&
                        <>
                            <div className="">Summary</div>
                            <div className="rounded-xl p-1 bg-gray-50 overflow-x-hidden h-[200px]">
                                {summary}
                            </div>
                        </>
                    }
                </div>
            }
        />
    )
}
export default PreviewDialog