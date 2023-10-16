import { ChildrenProps, ElementProps } from '@/types/props'

interface Props extends ChildrenProps, ElementProps {
}

const HStack = ({ children, className, centerX, centerY }: Props) => {
    var style = 'flex flex-col '
    style += className
    if (centerX) {
        style += 'justify-center '
    }
    if (centerY) {
        style += 'items-center '
    }
    return (
        <div className={style}>{children}</div>
    )
}
export default HStack