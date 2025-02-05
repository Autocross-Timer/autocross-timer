import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/leaderboard',
      redirect: '/leaderboard/pax'
    },
    {
      path: '/leaderboard/:sortClass',
      name: 'leaderboard',
      component: () => import('../views/LeaderboardView.vue'),
    },
    {
      path: '/watchlist',
      name: 'watchlist',
      component: () => import('../views/WatchlistView.vue'),
    },
    {
      path: '/car/:carId',
      name: 'car',
      component: () => import('../views/CarView.vue'),
    },
    {
      path: '/car',
      redirect: '/'
    },
    {
      path: '/admin',
      name: 'admin',
      component: () => import('../views/AdminView.vue'),
    }
  ],
})

export default router
