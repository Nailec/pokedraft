import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Draft from '../views/Draft.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/draft',
    name: 'Draft',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: Draft,
    props: route => ({
        roomID: route.query.roomKey,
        userID: route.query.userID
    })
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
