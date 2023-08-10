import { ApiErrorResponse, ApiResponse } from '@/app/_functions/API';
import { DependencyList, useEffect, useState } from 'react';

interface CustomQueryResult<T> {
    data: T | undefined;
    loading: boolean;
    error: ApiErrorResponse | undefined;
}

function useCustomQuery<T>(
    fetchFunction: Promise<ApiResponse<T>>,
    shouldFetch: boolean,
    dependencies: DependencyList = []
): CustomQueryResult<T> {
    const [data, setData] = useState<T>();
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<ApiErrorResponse>();

    useEffect(() => {
        if (shouldFetch) { 
            (async () => {
                setLoading(true);
                const { data, error, ok } = await fetchFunction;
                if (!ok) {
                    setError(error);
                } else {
                    setData(data);
                }
                setLoading(false);
            })()
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [shouldFetch, ...dependencies]);

    return { data, loading, error };
}

export default useCustomQuery;