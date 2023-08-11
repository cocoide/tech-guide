import { articleAPI } from '@/app/_functions/article'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'

export default async function OutlineSection({ articleURL }: { articleURL: string }) {
    const { data: overview } = await articleAPI.GetOverview(articleURL)
    if (!overview) {
        return null
    }
    return (
        <div className="relative flex flex-col  p-2 custom-border rounded-xl space-y-2">
            <div className="text-gray-500">Outlines</div>
            <ReactMarkdown remarkPlugins={[remarkGfm]} className='markdown'
            >{overview}</ReactMarkdown>
        </div>
    )
}