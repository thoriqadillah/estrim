import { http } from "../http"
import type { User } from "@/stores/account";

export const accountApi = {
  async getAccount() {
    return http.get<User>('/v1/account')
      .then(res => res.data)
      .catch(() => undefined)
  }
}