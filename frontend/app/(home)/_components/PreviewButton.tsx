"use client"
import { previewDialogAtom } from '@/stores/dialog';
import { PreviewDialogModel } from '@/types/dialog';
import { extractYoutubeID } from '@/utils/regex';
import { NewspaperIcon, PlayCircleIcon, PresentationChartBarIcon } from '@heroicons/react/24/outline';
import { useAtom } from 'jotai';

interface Props {
    url: string
    summary?: string
    domain: string
    youtube_id?: string
}
const PreviewButton = ({ url, domain, summary, youtube_id }: Props) => {
    const [_, setPreviewDialogOpen] = useAtom(previewDialogAtom)
    var speakerdeck_id: string | undefined = undefined
    if (domain === "speakerdeck.com") {
        speakerdeck_id = url
    }
    function handlePreviewButton() {
        youtube_id = extractYoutubeID(url)
        const model: PreviewDialogModel = {
            summary: summary,
            youtube_id: youtube_id,
            speakerdeck_id: speakerdeck_id,
        }
        setPreviewDialogOpen(model)
    }
    return (
        <button onClick={handlePreviewButton}>
            {domain === "youtube.com" ?
                <PlayCircleIcon className='h-6 w-6' />
                :
                domain === "speakerdeck.com" ?
                    <PresentationChartBarIcon className='h-6 w-6' />
                    :
                    summary &&
                    <NewspaperIcon className='h-6 w-6' />
            }
        </button>
    )
}
export default PreviewButton