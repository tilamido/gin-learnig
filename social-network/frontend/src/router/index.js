import { createRouter, createWebHashHistory } from "vue-router"
import LoginPage from "@/views/LoginPage.vue"
// import SuccessPage from "@/views/SuccessPage.vue"
import HomePage from "@/views/HomePage.vue";
import { requireAuth } from './auth';

const routes = [
    // 每个路由都需要映射到一个组件
    {
        path: '/',
        redirect: '/login', // 默认重定向到登录页面
    },
    {
        path: '/login',
        component: LoginPage,
    },
    {
        path: '/home',
        component: HomePage,
        meta: { requiresAuth: true },// 添加一个 meta 字段，用来标记需要登录后才能访问的页面
        children: [
            {
                path: '',
                redirect: '/home/index',
            },
            {
                path: 'index',
                name: 'index',
                component: () => import('../components/contents/HomeComponent.vue'),
            },
            {
                path: 'hotest',
                name: 'hotest',
                component: () => import('../components/contents/HotestComponent.vue'),
            },
            {
                path: 'publish',
                name: 'publish',
                component: () => import('../components/contents/PublishComponent.vue'),
            },
            {
                path: 'individualcomponent',
                name: 'individual',
                component: () => import('../components/contents/IndividualComponent.vue'),
            },
        ]
    },
]
//3、创建路由实例
const router = createRouter({
    routes,
    history: createWebHashHistory("./")
})

// 使用导航守卫来检查用户登录状态
router.beforeEach((to, from, next) => {
    // 如果页面需要登录权限，则调用 requireAuth 导航守卫函数检查用户登录状态
    if (to.meta.requiresAuth) {
        requireAuth(to, from, next);
    } else {
        // 如果页面不需要登录权限，则直接允许访问
        next();
    }
});
export default router