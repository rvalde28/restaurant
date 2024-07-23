import { createRouter, createWebHistory } from 'vue-router';
import App from '../App.vue';  // Replace with your component
import Checkout from '../components/Checkout.vue'; // Replace with your component

const routes = [
  {
    path: '/',
    name: 'Home',
    component: App
  },
  {
    path: '/checkout',
    name: 'Checkout',
    component: Checkout
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;