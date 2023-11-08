import { ChildrenProps } from '@/types/props'

interface Props extends ChildrenProps { }
const ArticlesContainer = ({ children }: Props) => {
    return (
        <div>ArticlesContainer</div>
    )
}
export default ArticlesContainer