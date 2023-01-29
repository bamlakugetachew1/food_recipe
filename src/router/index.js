import { createRouter, createWebHistory } from 'vue-router'
import home from '../components/home.vue'
import singleproduct from '../components/singleproduct.vue'
import adminpage from '../components/adminpage.vue'
import loginpage from '../components/loginpage.vue'
import signup from '../components/signup.vue'
import addproduct from '../components/addproduct.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },
    {
        path: '/singleproduct/:id',
        name: 'singleproduct',
        component: singleproduct
      },
      {
        path: '/adminpage',
        name: 'adminpage',
        component: adminpage
      },
      {
        path: '/loginpage',
        name: 'loginpage',
        component: loginpage
      },
      {
        path: '/signup',
        name: 'signup',
        component: signup
      },
      {
        path: '/addproduct',
        name: 'addproduct',
        component: addproduct
      },
  ]
})

export default router
