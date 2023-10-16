import { ChildrenProps, ElementProps } from '@/types/props'

interface Props extends ChildrenProps, ElementProps {
}

const WStack = ({ children, className, centerX, centerY }: Props) => {
    var style = 'flex flex-row '
    style += className
    if (centerX) {
        style += 'items-center '
    }
    if (centerY) {
        style += 'justify-center '
    }
    return (
        <div className={style}>{children}</div>
    )
}
export default WStack