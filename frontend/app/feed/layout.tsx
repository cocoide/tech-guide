import RightSidevar from './_components/RightSidevar'

interface Props {
    children: React.ReactNode
    modal: React.ReactNode
}
export default function FeedLayout({ children, modal }: Props) {
    return (
        <>
            <div className="flex flex-row relative">
                {children}
                <div className="hidden lg:flex sticky top-0 h-screen"
                ><RightSidevar /></div>
            </div>
            {modal}
        </>
    )
}