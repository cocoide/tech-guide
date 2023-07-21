interface Props {
    children: React.ReactNode
    modal: React.ReactNode
}
const HomeLayout = ({ children, modal }: Props) => {
    return (
        <>
            {children}
            {modal}
        </>
    )
}
export default HomeLayout