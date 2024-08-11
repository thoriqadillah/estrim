import { ref, computed } from "vue";
import { defineStore } from "pinia";
import { api } from "@/lib";

export interface User {
  id: string;
  name?: string;
  email?: string;
  password?: string;
  source: string;
  created_at: Date;
}

export const useAccount = defineStore("account", () => {
  const user = ref<User>();
  const authenticated = computed(() => user.value && user.value.source !== 'session');

  async function getAccount() {
    const _user = await api.getAccount()
    user.value = _user
  }

  function logout() {
    user.value = undefined
  }

  return { 
    user,
    authenticated,
    getAccount,
    logout,
  }
});
