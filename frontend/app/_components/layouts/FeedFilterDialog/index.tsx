"use client"
import { topicDialogAtom } from '@/stores/dialog'
import CustomDialog from '../../elements/CustomDialog'
import DialogContent from './DialogContent'

const FeedFilterDialog = () => {
    return (
        <CustomDialog
            openAtom={topicDialogAtom}
            layout='mt-[120px] sm:mb-[120px]  bg-white z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] rounded-xl'
            content={
                <DialogContent />
            }
        />
    )
}
export default FeedFilterDialog
