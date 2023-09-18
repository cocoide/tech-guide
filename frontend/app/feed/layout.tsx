
interface Props {
    children: React.ReactNode
}
export default function FeedLayout({ children }: Props) {
    return (
        <>
            <div className="flex flex-row relative">
                {children}
            </div>
        </>
    )
}