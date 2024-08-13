import { createRouter, createWebHistory } from 'vue-router';
import RegisterPage from '/src/views/RegisterPage.vue';
import LoginPage from '/src/views/LoginPage.vue';
import Choose from '../views/ChoosePage.vue';
import LectioPage from '../views/LectioPage.vue';
import UpdateStatus from '../views/UpSatusPage.vue';
import UpdatePage from '@/views/UpPage.vue';
import WriteQuote from '@/views/WriteQuotePage.vue';
import WriteReview from '@/views/WriteReviewPage.vue';
import DeletePage from '@/views/DeletePage.vue';
import AddBook from '../views/AddbookPage.vue';

const routes = [
  { path: '/', component: LoginPage },
  { path: '/registrar', component: RegisterPage },
  { path: '/choose', component: Choose, meta: { requiresAuth: true } },
  { path: '/lectio', component: LectioPage, meta: { requiresAuth: true } },
  { path: '/registerbook', component: AddBook, meta: { requiresAuth: true } },
  { path: '/atualizar/status/:title', name: 'status', component: UpdateStatus, meta: { requiresAuth: true } },
  { path: '/atualizar/:title', name: 'atualizar', component: UpdatePage, meta: { requiresAuth: true } },
  { path: '/citacao/:title', name: 'citacao', component: WriteQuote, meta: { requiresAuth: true } },
  { path: '/review/:title', name: 'review', component: WriteReview, meta: { requiresAuth: true } },
  { path: '/apagar/:title', name: 'apagar', component: DeletePage, meta: { requiresAuth: true } }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
  const token = localStorage.getItem('token');

  if (requiresAuth && !token) {
    next('/');
  } else {
    next();
  }
});

export default router;
