import { createApp } from 'vue';
import App from './App.vue';
import router from "./router";
import './assets/body.css';


// 创建应用实例
const app = createApp(App);

// 定义全局变量，并绑定到 Vue 的原型链上
localStorage.setItem('isAuthenticated', false);

app.use(router).mount('#app');