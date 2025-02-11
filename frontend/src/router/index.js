import { createRouter, createWebHistory } from 'vue-router'
import EventsView from '../views/EventsView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: EventsView,
    },
    {
      path: '/event/:eventId',
      name: 'event',
      component: () => import('../views/EventView.vue'),
    },
    {
      path: '/event/:eventId/leaderboard/:sortClass',
      name: 'leaderboard',
      component: () => import('../views/LeaderboardView.vue'),
    },
    {
      path: '/event/:eventId/leaderboard/',
      name: 'leaderboard-redirect',
      component: () => import('../views/LeaderboardView.vue'),
    },
    {
      path: '/event/:eventId/car/:carId',
      name: 'car',
      component: () => import('../views/CarView.vue'),
    },    
    {
      path: '/event/:eventId/watchlist',
      name: 'watchlist',
      component: () => import('../views/WatchlistView.vue'),
    },
    {
      path: '/admin',
      name: 'admin',
      component: () => import('../views/AdminView.vue'),
    }
  ],
})

export default router
