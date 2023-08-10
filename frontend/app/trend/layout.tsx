import RightSidevar from './_components/RightSidevar'

interface Props {
    children: React.ReactNode
    modal: React.ReactNode
}
const TrendLayout = ({ children }: Props) => {
    return (
        <>
            <div className="flex flex-row relative">
                {children}
                <div className="hidden lg:flex sticky top-0 h-screen"
                ><RightSidevar /></div>
            </div>
        </>
    )
}
export default TrendLayout