const OutlineLoader = () => {
    return (
        <div className="flex flex-col  p-2 ring-1 ring-gray-300 rounded-md space-y-2">
            <div className="text-gray-500">Outlines</div>
            {Array.from({ length: 5 }).map((_, index) => (
                <div key={index} className="animate-pulse w-[100%] rounded-md bg-gray-200 h-[50px]" />
            ))}
        </div>
    )
}
export default OutlineLoader