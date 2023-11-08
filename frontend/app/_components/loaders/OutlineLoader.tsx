const OutlineLoader = () => {
    return (
        <div className="flex flex-col  p-2 custom-border rounded-xl space-y-2 w-full">
            <div className="text-gray-500">Outlines</div>
            {Array.from({ length: 5 }).map((_, index) => (
                <div key={index} className="animate-pulse w-[100%] rounded-md bg-gray-100 dark:bg-gray-600 h-[50px]" />
            ))}
        </div>
    )
}
export default OutlineLoader