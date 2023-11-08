const OutlineLoader = () => {
    return (
        <>
            {Array.from({ length: 5 }).map((_, index) => (
                <div key={index} className="animate-pulse w-[100%] rounded-md bg-gray-100 dark:bg-gray-600 h-[50px]" />
            ))}
        </>
    )
}
export default OutlineLoader