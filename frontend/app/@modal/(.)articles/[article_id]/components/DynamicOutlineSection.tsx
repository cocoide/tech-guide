import { articleAPI } from '@/app/_functions/article';

export type Header = {
    content: string
    subHeader?: Header
}

function printHeaders(header?: Header, level: number = 1): void {
    const prefix = `header${level}`;
    console.log(prefix, header?.content);

    if (header?.subHeader) {
        printHeaders(header.subHeader, level + 1);
    }
}
export default async function DynamicOutlineSection() {
    const { data: headers } = await articleAPI.GetHeaders()
    if (headers && headers.length > 2) {
        printHeaders(headers[2])
    }
    return (
        <div className="relative flex flex-col  p-2 custom-border rounded-xl space-y-2">
            <div className="text-gray-500">Outlines</div>
        </div>
    )
}