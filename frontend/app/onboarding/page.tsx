import RegisterContainer from './_containers/RegisterContainer';
import SignupContainer from './_containers/SignupContainer';
import TopicSelectContainer from './_containers/TopicSelectContainer';
import { getSessionID, getSignupSession } from './_functions/session';

export default async function OnboardingPage() {
    const sessionId = await getSessionID()
    const { data: session } = await getSignupSession(sessionId)
    const step = session?.onboarding_index
    return (
        <div className="w-full">
            {step === 1 ?
                <RegisterContainer sessionId={sessionId} avatar_url={session?.avatar_url} display_name={session?.display_name} />
                    :
                step === 2 ?
                    <TopicSelectContainer />
                    :
                <SignupContainer />
            }
        </div>
    )
}