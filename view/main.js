// ルートを定義
const routes = [
  {
    path: "/",
    name: "top",
    component: httpVueLoader("http://localhost:8080/view/components/Top.vue")
  },
  {
    path: "/list",
    name: "list",
    component: httpVueLoader("http://localhost:8080/view/components/List.vue")
  },
  {
    path: "/search",
    name: "search",
    component: httpVueLoader("http://localhost:8080/view/components/Search.vue")
  }
];

// ルーターインスタンスを作成し、ルートオプションを渡す
const router = new VueRouter({
  mode: "history",
  routes // `routes: routes` の短縮表記
});

// root となるインスタンスを作成してマウント
const app = new Vue({
  router
}).$mount("#app");
