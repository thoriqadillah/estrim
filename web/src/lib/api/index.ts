import { accountApi } from "./account.api";
import { fileApi } from "./file.api";

export const api = {
  ...accountApi,
  ...fileApi,
}