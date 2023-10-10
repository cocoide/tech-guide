import { NewspaperIcon, PlayCircleIcon, PresentationChartBarIcon } from '@heroicons/react/24/outline'

const PreviewButton = ({ domain }: { domain: string }) => {

    return (
        <button>
            {domain === "youtube.com" ?
                <PlayCircleIcon className='h-6 w-6' />
                :
                domain === "speakerdeck.com" ?
                    <PresentationChartBarIcon className='h-6 w-6' />
                    :
                    <NewspaperIcon className='h-6 w-6' />
            }
        </button>
    )
}
export default PreviewButton