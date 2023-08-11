import { ReactNode } from 'react'

interface Props {
    height: number
    children: ReactNode
}
const ScrollButton = ({ height, children }: Props) => {
    function scroll() {
        window.scrollTo({
            top: height,
            behavior: 'smooth'
        })
    }
    return (
        <button onClick={scroll}>{children}</button>
    )
}
export default ScrollButton