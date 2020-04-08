import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/views/Login'
import admin from '@/views/Admin'
import cliente from '@/views/Empleado'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
        path: '/',
        name: 'Login',
        component: Login
    },
    {
        path: '/admin',
        name: 'admin',
        component: admin
    },
    {
        path: '/user',
        name: 'empleado',
        component: cliente
    }
    ]
})