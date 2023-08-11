"use client"
import { topicDialogAtom } from '@/stores/dialog'
import { useQueryClient } from '@tanstack/react-query'
import CustomDialog from '../../elements/CustomDialog'
import DialogContent from './DialogContent'

const FeedFilterDialog = () => {
    const queryClient = useQueryClient();
    function refetchQuery() {
        queryClient.invalidateQueries({ queryKey: ['feeds_query'] })
        window.location.reload()
    }
    return (
        <CustomDialog
            openAtom={topicDialogAtom}
            layout='mt-[120px] sm:mb-[120px]  bg-white z-50 sm:mx-[15%] md:mx-[20%] lg:mx-[25%] rounded-xl'
            content={
                <DialogContent />
            }
            closeFunc={refetchQuery}
        />
    )
}
export default FeedFilterDialog
