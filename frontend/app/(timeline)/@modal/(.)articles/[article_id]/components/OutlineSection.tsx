import { articleAPI } from '@/app/(timeline)/_functions/article'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'

export default async function OutlineSection({ articleURL }: { articleURL: string }) {
    const { data: overview } = await articleAPI.GetOverview(articleURL)
    if (!overview) {
        return null
    }
    return (
        <div className="flex flex-col  p-2 ring-1 ring-gray-300 rounded-md space-y-2">
            <div className="text-gray-500">Outlines</div>
            <ReactMarkdown remarkPlugins={[remarkGfm]} className='markdown'
            >{overview}</ReactMarkdown>
        </div>
    )
}