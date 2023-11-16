import RegisterContainer from './_containers/RegisterContainer';
import SignupContainer from './_containers/SignupContainer';
import TopicSelectContainer from './_containers/TopicSelectContainer';
import { getSessionID, getSignupSession } from './_functions/session';

interface Porps {
    searchParams: { "step": string }
}
export default async function OnboardingPage({ searchParams }: Porps) {
    const sessionId = await getSessionID()
    const { data: session } = await getSignupSession(sessionId)
    const step = session?.onboarding_index
    return (
        <div className="w-full">
            {step === 1 ?
                <RegisterContainer sessionId={sessionId} />
                    :
                step === 2 ?
                    <TopicSelectContainer />
                    :
                <SignupContainer />
            }
        </div>
    )
}