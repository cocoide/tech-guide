import { clsx } from '@/utils/clsx'

interface Porps {
    openFirst: () => void
    openSecond: () => void
    openThird: () => void
    openSection: number
}
export default function SettingMenue({ openFirst, openSecond, openThird, openSection }: Porps) {
    return (
        <div className="flex flex-row items-center justify-center w-[300px] custom-border rounded-r-xl rounded-l-xl
        text-gray-500">
            <button onClick={openFirst}
                className={clsx("custom-badge w-[33%]  justify-center p-1 rounded-l-xl duration-500",
                    openSection === 1 ? "text-gray-600 bg-gray-100" : "")}>トピック</button>
            <button onClick={openSecond}
                className={clsx("custom-badge w-[33%]  justify-center p-1 border-x-[0.5px] duration-500",
                    openSection === 2 ? "text-gray-600 bg-gray-100" : "")}>ドメイン</button>
            <button onClick={openThird}
                className={clsx("custom-badge w-[33%]  justify-center p-1 rounded-r-xl duration-500",
                    openSection === 3 ? "text-gray-600 bg-gray-100" : "")}>高度な設定</button>
        </div>
    )
}