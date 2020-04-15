import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
        path: '/',
        name: 'Login',
        component: ()=> import(/* webpackChunkName: "Login" */ '@/views/Login')
    },
    {
        path: '/admin',
        name: 'admin',
        component: ()=> import(/* webpackChunkName: "admin" */'@/views/Admin')
    },
    {
        path: '/user',
        name: 'empleado',
        component: ()=> import(/* webpackChunkName: "user" */'@/views/Empleado')
    },
    {
        path: '*',
        name: 'Error',
        component: ()=> import(/* webpackChunkName: "Error" */'@/views/Error')
    }
    ]
})