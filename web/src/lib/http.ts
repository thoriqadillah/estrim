import axios, {
  AxiosError,
  type AxiosResponse,
  type InternalAxiosRequestConfig,
} from "axios";
import { eventBus } from "./event-bus";
import type { Toast } from "@/components/ui/toast/use-toast";

const REFRESH_TOKEN = "/v1/account/refresh-token";

const http = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL + "/api",
  withCredentials: true
});

async function onError(error: AxiosError) {
  if (!error.response) return;

  const res = error.response;
  const origin = error.config as any;
  
  if (res.status === 401 && origin.url !== REFRESH_TOKEN && !origin._retry) {
    origin._retry = true;
    try {
      const response = await http.get(REFRESH_TOKEN);
      localStorage.setItem("access-token", response.data.token);
      return http(origin);
    } catch (error) {
      if (error instanceof AxiosError) {
        const err = error.toJSON() as AxiosError;
        if (err.status === 401) {
          localStorage.removeItem('access-token')
          eventBus.emit("logout");
        }
      }
    }
  }

  if (res.status >= 400 && res.status < 600) {
    const message: string = (res.data as any).message
    eventBus.emit<Toast>('toast', {
      title: 'Error',
      description: message[0].toUpperCase() + message.substring(1),
      variant: 'destructive',
    })

    return Promise.reject(error);
  }

  return res;
}

function onRequest(config: InternalAxiosRequestConfig) {
  const token = localStorage.getItem("access-token");
  if (!token) return config;

  config.headers.Authorization = `Bearer ${token}`;
  return config;
}

function onResponse(response: AxiosResponse): AxiosResponse {
  return response;
}

http.interceptors.request.use(onRequest, onError);
http.interceptors.response.use(onResponse, onError);

export { http };
