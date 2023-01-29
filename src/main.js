import router from "./router";
import { createApp, provide, h } from "vue";
import { DefaultApolloClient } from "@vue/apollo-composable";
import { ApolloClient, InMemoryCache } from "@apollo/client/core";
import App from "./App.vue";
import "./index.css";
import store from './store/store' 
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

const cache = new InMemoryCache();
const token = localStorage.getItem("token");
const headers = {
  "Content-Type": "application/json",
};
if (token != null) {
  headers.Authorization = `Bearer ${token}`;
}
const apolloClient = new ApolloClient({
  cache,
  uri: "http://localhost:8081/v1/graphql",
  headers: headers,
});

const app = createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient);
  },

  render: () => h(App),
});

app.use(store).use(router).mount("#app");
