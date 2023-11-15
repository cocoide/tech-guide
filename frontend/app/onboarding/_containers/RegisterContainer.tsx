
export default async function RegisterContainer() {
    function register(formData: FormData) {
    }

    return (
        <form action={register} className='w-full flex flex-row justify-center'>
            <input type="text" name="name" />
            <input type="text" name="avatar" />
            <button type="submit">登録</button>
        </form>
    )
}