import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '@/views/login/LoginView.vue';
import ViewIndex from '@/views/index/ViewIndex.vue';

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        component: LoginView
    },
    {
        path: '/dashboard',
        component: ViewIndex,
        name: 'dashboard' // 添加路由名称为 'dashboard'
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;