"use client"
import { sourceAPI } from '@/app/(dashboard)/settings/_functions/source'
import { topicAPI } from '@/app/(dashboard)/settings/_functions/topic'
import { useAuth } from '@/hooks/useAuth'
import { topicDialogAtom } from '@/stores/dialog'
import { XMarkIcon } from '@heroicons/react/24/outline'
import { useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import { useState } from 'react'
import DetailSettingSection from './DetailSettingSection'
import DomainFollowSection from './DomainFollowSection'
import SettingMenue from './SetttingMenue'
import TopicFollowAccordionGroup from './TopicFollowAccordionGroup'

export default function DialogContent() {
    function handleClose() {
        window.location.reload()
        setDialogOpen(false)
    }
    const { token } = useAuth();
    const [openSection, setOpenSection] = useState(1)
    const [_, setDialogOpen] = useAtom(topicDialogAtom)
    const { data: followingTopics } = useQuery({
        queryFn: async () => (await topicAPI.GetFollowingTopics(token)).data,
        queryKey: ["following_topics"],
        enabled: openSection === 1
    })
    const { data: existingCategories } = useQuery({
        queryFn: async () => (await topicAPI.GetAllCategories()).data,
        queryKey: ["existing_categories"],
        enabled: openSection === 1
    })
    const { data: followingSources } = useQuery({
        queryFn: async () => (await sourceAPI.GetFollowingSources()).data,
        queryKey: ["following_sources"],
        enabled: openSection === 2
    })
    const { data: existingSources } = useQuery({
        queryFn: async () => (await sourceAPI.GetAllSources()).data,
        queryKey: ["existing_sources"],
        enabled: openSection === 2
    })
    return (
        <div className='flex flex-col w-full h-full overflow-y-scroll p-3 space-y-5 relative'>
            <button onClick={handleClose}
            ><XMarkIcon className='h-6 w-6 text-gray-500 absolute right-3 top-3' /></button>
            <div className="flex flex-row justify-center">
                <SettingMenue
                    openFirst={() => setOpenSection(1)}
                    openSecond={() => setOpenSection(2)}
                    openThird={() => setOpenSection(3)}
                    openSection={openSection} />
            </div>
            {openSection === 1 ? (
                <TopicFollowAccordionGroup categories={existingCategories} following_topics={followingTopics} token={token} />
            ) : openSection === 2 ? (
                <DomainFollowSection token={token} followingSources={followingSources} existingSources={existingSources} />
            ) : (
                <DetailSettingSection />
            )}
        </div>
    )
}