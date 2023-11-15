import RegisterContainer from './_containers/RegisterContainer';
import SignupContainer from './_containers/SignupContainer';

interface Porps {
    searchParams: { "step": string }
}
export default async function OnboardingPage({ searchParams }: Porps) {
    const step = Number(searchParams.step);
    return (
        <div className="w-full">
            {step === 2 ?
                    <RegisterContainer />
                    :
                <SignupContainer />
            }
        </div>
    )
}