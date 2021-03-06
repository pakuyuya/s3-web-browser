import Vue from 'vue';
import Router from 'vue-router';
import Home from '../../views/Home.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL + (process.env.NODE_ENV === 'production' ? 'browser/' : ''),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/s3/:profileid',
      name: 's3',
      component: () => import('../../views/S3.vue'),
    },
    {
      path: '/s3/:profileid/:path*',
      name: 's3pathed',
      component: () => import('../../views/S3.vue'),
    },
  ],
});
