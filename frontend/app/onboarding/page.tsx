import RegisterContainer from './_containers/RegisterContainer';
import SignupContainer from './_containers/SignupContainer';
import TopicSelectContainer from './_containers/TopicSelectContainer';
import { getSignupSession } from './_functions/session';

interface Porps {
    searchParams: { "step": string }
}
export default async function OnboardingPage({ searchParams }: Porps) {
    const { data: session } = await getSignupSession()
    const step = session?.onboarding_Index
    return (
        <div className="w-full">
            {step === 1 ?
                    <RegisterContainer />
                    :
                step === 2 ?
                    <TopicSelectContainer />
                    :
                <SignupContainer />
            }
        </div>
    )
}