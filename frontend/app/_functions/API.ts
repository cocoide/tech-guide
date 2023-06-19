
interface ApiService {
  get<T>(DIR_URL: string, cache?: Cache, token?: string, params?: Params): Promise<ApiResponse<T>>
  del(DIR_URL: string, params?: Params): Promise<ApiResponse<void>>
  pos<U>(DIR_URL: string, body: any, token?: string,params?: Params): Promise<ApiResponse<U>>
  put<U>(DIR_URL: string, body: U, params?: Params): Promise<ApiResponse<void>>
}

type ApiResponse<T> = {
  data?: T
  error?: ApiErrorResponse
  status?: number
  ok: boolean
}
type ApiErrorResponse = string | { message: string }

type Params = {
  [key: string]: any
}
type Cache = 'no-store' | 'force-cache' | 'reload'| number // SSR | SSG | SSG (from when reload) | ※default SSG
// 公式ドキュメント: https://nextjs.org/docs/app/building-your-application/data-fetching/caching

const headers: HeadersInit = {
  "Content-Type": 'application/json',
}
export const api: ApiService = {
  async get<T>(DIR_URL: string, cache?: Cache, token?: string, params?: Params): Promise<ApiResponse<T>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: "GET",
      headers: headers,
    }
    if (typeof cache === 'string') {
      options.cache=cache
    }else if (typeof cache === 'number'){
      options.next={...options.next, revalidate: cache}
    }
    return handleApiRequest(DIR_URL, options, params)
  },
  async pos<U>(DIR_URL: string, body: any, token?: string, params?: Params): Promise<ApiResponse<U>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    console.log(body)
    const options: RequestInit = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(body)
    }
    return handleApiRequest(DIR_URL, options, params)
  },
  async del(DIR_URL: string, params?: Params): Promise<ApiResponse<void>> {
    const options: RequestInit = {
      method: 'DELETE',
      headers: headers,
    }
    return handleApiRequest(DIR_URL, options, params)
  },
  async put<U>(DIR_URL: string, body: U, params?: Params): Promise<ApiResponse<void>> {
    const options: RequestInit = {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(body)
    }
    return handleApiRequest(DIR_URL, options, params)
  },
}

async function handleApiRequest<T>(DIR_URL: string, options: RequestInit, params?: Params): Promise<ApiResponse<T>> {
  const API_URL = buildApiURL(DIR_URL, params)
  const res = await fetch(API_URL, options);
  try {
    if (!res.ok) {
      return {
        error: await res.json(),
        status: res.status,
        ok: false,
      };
    } else {
      return {
        data: await res.json(),
        status: res.status,
        ok: true,
      };
    }
  } catch (error) {
    return {
      error: `unexpected error occured: ${error}`,
      status: 500,
      ok: false,
    };
  }
}

function buildApiURL(DIR_URL: string, params?: Params): string {
  const QueryParam = params
    ? '?' + Object.keys(params)
      .map((key) => params[key].map((value: any) => `${encodeURIComponent(key)}=${encodeURIComponent(value)}`).join('&'))
      .join('&')
    : '';
    // const BaseURL = process.env.NEXT_BASE_URL
  return  "http://localhost:8080" + DIR_URL + QueryParam
}