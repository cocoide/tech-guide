import HStack from '@/app/_components/elements/ui/HStack';

export default async function RegisterContainer() {
    function register(formData: FormData) {
    }

    return (
        <HStack className="w-full items-center space-y-5 custom-text">
        <form action={register} className='w-full flex flex-row justify-center'>
            <input type="text" name="name" />
            <input type="text" name="avatar" />
            <button type="submit">登録</button>
        </form>
        </HStack>
    )
}