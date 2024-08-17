import type { UploadedFile } from "@/pages/home/types";
import { http } from "../http";

export const fileApi = {
  async compress(form: FormData, onProgress?: (progress: number) => void) {
    return http.post<UploadedFile>('/v1/compress', form, {
      onUploadProgress: (e) => {
        onProgress && onProgress((e.progress || 0) * 100)
      }
    })
    .then(res => res.data)
    .catch(() => undefined)
  },
}