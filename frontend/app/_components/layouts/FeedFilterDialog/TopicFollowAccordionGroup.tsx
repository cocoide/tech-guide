
import { Category, Topic } from '@/types/model';
import { clsx } from '@/utils/clsx';
import { ChevronDownIcon } from '@heroicons/react/24/outline';
import * as Accordion from '@radix-ui/react-accordion';
import React, { forwardRef } from 'react';
import TopicFollowSection from './TopicFollowSection';

interface Props {
    token?: string;
    categories?: Category[];
    following_topics?: Topic[];
}
const TopicFollowAccordionGroup = ({ categories, token, following_topics }: Props) => {
    const categorizedFollowingTopics: Record<string, Topic[]> = {};
    following_topics?.forEach((topic) => {
        const category = categories?.find((cat) => {
            return cat.topics.some((t) => t.id === topic.id);
        });

        if (category) {
            if (!categorizedFollowingTopics[category.id]) {
                categorizedFollowingTopics[category.id] = [];
            }
            categorizedFollowingTopics[category.id].push(topic);
        }
    });
    return (
        <Accordion.Root
            className="w-full text-gray-500 dark:text-gray-200 dark:bg-black"
            type="single"
            collapsible
        >
            {categories?.map((category, index) => (
                <>
                    <AccordionItem value={`items-${index}`} >
                        <AccordionTrigger>
                            <div className="flex flex-row justify-between items-center pr-2 w-full">
                                <> {category.name}</>
                                {categorizedFollowingTopics[category.id].length > 0 &&
                                    <div className="h-5 w-5 bg-cyan-300 text-white rounded-full custom-badge justify-center"
                                    >{categorizedFollowingTopics[category.id].length}</div>
                                }
                            </div>
                        </AccordionTrigger>
                        <AccordionContent>
                            <TopicFollowSection existingTopics={category.topics} token={token} followingTopics={categorizedFollowingTopics[category.id]} />
                        </AccordionContent>
                    </AccordionItem>
                </>
            ))}

        </Accordion.Root>
    )
}
export default TopicFollowAccordionGroup

interface AccordionItemProps {
    children: React.ReactNode;
    className?: string;
    value: string
}

// eslint-disable-next-line react/display-name
const AccordionItem = forwardRef<HTMLDivElement, AccordionItemProps>(({ children, className = '', ...props }, forwardedRef) => (
    <Accordion.Item
        className={clsx(
            ' ring-0 ring-transparent outline-none mt-px overflow-hidden first:mt-0 first:rounded-t last:rounded-b focus-within:relative focus-within:z-10',
            className,
        )}
        {...props}
        ref={forwardedRef}
    >
        {children}
    </Accordion.Item>
));

// eslint-disable-next-line react/display-name
const AccordionTrigger = forwardRef<HTMLButtonElement, Omit<AccordionItemProps, 'value'>>(({ children, className = '', ...props }, forwardedRef) => (
    <Accordion.Header className="flex">
        <Accordion.Trigger
            className={clsx(
                ' ring-transparent group flex h-[45px] flex-1 cursor-default items-center justify-between px-5 text-[15px] leading-none outline-none',
                className,
            )}
            {...props}
            ref={forwardedRef}
        >
            {children}
            <ChevronDownIcon
                className="ease-[cubic-bezier(0.87,_0,_0.13,_1)] transition-transform duration-300 group-data-[state=open]:rotate-180 h-5 w-5"
                aria-hidden
            />
        </Accordion.Trigger>
    </Accordion.Header>
))

// eslint-disable-next-line react/display-name
const AccordionContent = forwardRef<HTMLDivElement, Omit<AccordionItemProps, 'value'>>(({ children, className = '', ...props }, forwardedRef) => (
    <Accordion.Content
        className={clsx(
            'data-[state=open]:animate-slideDown data-[state=closed]:animate-slideUp overflow-hidden text-[15px]',
            className,
        )}
        {...props}
        ref={forwardedRef}
    >
        <div className="py-[5px] px-[5px]">{children}</div>
    </Accordion.Content>
));