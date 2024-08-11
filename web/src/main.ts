import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import { AuthGuard, router } from "./router";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(AuthGuard);

app.mount("#app");
