import Vue from 'vue'
import Router from 'vue-router'

import Login from '@/views/Login'
import Registro from '@/views/Register'
import Home from '@/views/Page'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
        path: '/',
        name: 'Login',
        component: Login
    },
    {
        path: '/registro',
        name: 'Registro',
        component: Registro
    },
    {
        path: '/home',
        name: 'home',
        component: Home
    }
    ]
})