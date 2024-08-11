import type { UploadedFile } from "@/pages/home/types";
import { http } from "./http"
import type { User } from "@/stores/account";

export const api = {
  async upload(form: FormData, onProgress?: (progress: number) => void) {
    const { data } = await http.post<UploadedFile>('/v1/compress', form, {
      onUploadProgress: (e) => {
        const progress = Math.round((e.loaded * 100) / (e.total || 0));
        onProgress && onProgress(progress)
      }
    })

    return data
  },

  async getAccount() {
    const { data } = await http.get<User>('/v1/account/session')
    return data
  }
}