
interface ApiService {
  get<T>(dirURL: string, cache?: Cache, token?: string, params?: Params): Promise<ApiResponse<T>>
  del(dirURL: string,  token?: string, params?: Params): Promise<ApiResponse<void>>
  pos<U>(dirURL: string, body: any, token?: string,params?: Params): Promise<ApiResponse<U>>
  put<U>(dirURL: string, body: U,  token?: string, params?: Params): Promise<ApiResponse<void>>
}

export type ApiResponse<T> = {
  data?: T
  error?: ApiErrorResponse
  status?: number
  ok: boolean
}
export type ApiErrorResponse = string | { message: string }

type ParamValue = string | number | boolean | string[] | number[]
type Params = {
  [key: string]: ParamValue
}
type Cache = 'no-store' | 'force-cache' | 'reload'|'reload'| number // SSR | SSG | SSG (from when reload) | ※default SSG
// 公式ドキュメント: https://nextjs.org/docs/app/building-your-application/data-fetching/caching

const headers: HeadersInit = {
  "Content-Type": 'application/json',
}
export const api: ApiService = {
  async get<T>(dirURL: string, cache?: Cache, token?: string, params?: Params): Promise<ApiResponse<T>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: "GET",
      headers: headers,
      credentials: 'include',
      mode: 'cors',
    }
    if (typeof cache === 'string') {
      options.cache=cache
    }else if (typeof cache === 'number'){
      options.next={...options.next, revalidate: cache}
    }
    return handleApiRequest(dirURL, options, params)
  },
  async pos<U>(dirURL: string, body: any, token?: string, params?: Params): Promise<ApiResponse<U>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(body)
    }
    return handleApiRequest(dirURL, options, params)
  },
  async del(dirURL: string,  token?: string,  params?: Params): Promise<ApiResponse<void>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: 'DELETE',
      headers: headers,
    }
    return handleApiRequest(dirURL, options, params)
  },
  async put<U>(dirURL: string, body: U,  token?: string, params?: Params): Promise<ApiResponse<void>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(body),
    }
    return handleApiRequest(dirURL, options, params)
  },
}

async function handleApiRequest<T>(dirURL: string, options: RequestInit, params?: Params): Promise<ApiResponse<T>> {
  const apiURL = buildApiURL(dirURL, "backend", params);
  console.log(apiURL)
  const res = await fetch(apiURL, options);
  try {
    return {
      data: res.ok ? await res.json() : undefined,
      error: res.ok ? undefined : await res.json(),
      status: res.status,
      ok: res.ok,
    };
  } catch (error) {
    return {
      error: await res.json() + `: ${error}`,
      status: 500,
      ok: false,
    };
  }
}

function buildApiURL(dirURL: string, mode: "backend" | "nextjs", params?: Params): string {
  var baseURL = ""
  if (mode === "nextjs") {
    baseURL = `${process.env.NEXT_PUBLIC_FRON_URL}/api`
  } else if (mode === "backend") {
    baseURL = process.env.NEXT_PUBLIC_API_BASE_URL
  }
  const queryParam = params
    ? '?' + Object.keys(params)
      .map((key) => {
        const paramValue = params[key];
        if (Array.isArray(paramValue)) {
          return paramValue.map((element) => `${encodeURIComponent(key)}=${encodeURIComponent(element)}`).join('&');
        }
        return `${encodeURIComponent(key)}=${encodeURIComponent(paramValue)}`;
      })
      .join('&')
    : '';
  return baseURL + dirURL + queryParam
}

export const apiRoute: ApiService = {
  async get<T>(dirURL: string, cache?: Cache, token?: string, params?: Params): Promise<ApiResponse<T>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: "GET",
      headers: headers,
      credentials: 'include',
      mode: 'cors',
    }
    if (typeof cache === 'string') {
      options.cache = cache
    } else if (typeof cache === 'number') {
      options.next = { ...options.next, revalidate: cache }
    }
    return handleApiRequest(dirURL, options, params)
  },
  async pos<U>(dirURL: string, body: any, token?: string, params?: Params): Promise<ApiResponse<U>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: "POST",
      headers: headers,
      body: JSON.stringify(body)
    }
    return handleApiRequest(dirURL, options, params)
  },
  async del(dirURL: string, token?: string, params?: Params): Promise<ApiResponse<void>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: 'DELETE',
      headers: headers,
    }
    return handleApiRequest(dirURL, options, params)
  },
  async put<U>(dirURL: string, body: U, token?: string, params?: Params): Promise<ApiResponse<void>> {
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
    const options: RequestInit = {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(body),
    }
    return handleApiRequest(dirURL, options, params)
  },
}
