import Vue from 'vue';
import Router from 'vue-router';

import Accounts from '@/pages/Accounts.vue';
import Login from '@/pages/Login.vue';
import Projects from '@/pages/Projects.vue';
import ProjectIndex from '@/pages/projects/Index.vue';

import auth from '@/services/auth';

Vue.use(Router);

function requires_auth(to, from, next) {
  if (!auth.is_authenticated()) {
    next({
      path: '/login',
      query: { redirect: to.fullPath },
    });
  }

  next();
}

function cant_when_authenticated(to, from, next) {
  console.log(auth.is_authenticated());
  if (auth.is_authenticated()) {
    next({
      path: '/',
    });
  } else {
    next();
  }
}


export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      beforeEnter: requires_auth,
      redirect: '/projects',
    },
    {
      path: '/login',
      name: 'login',
      beforeEnter: cant_when_authenticated,
      component: Login,
    },
    {
      path: '/accounts',
      name: 'accounts',
      beforeEnter: requires_auth,
      component: Accounts,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './pages/About.vue'),
    },
    {
      path: '/projects',
      name: 'projects',
      beforeEnter: requires_auth,
      component: Projects,
    },
    {
      path: '/projects/:project_id',
      beforeEnter: requires_auth,
      children: [
        {
          path: '',
          name: 'project_home',
          component: ProjectIndex,
        },
      ],
    },
    { path: '*', redirect: '/' },
  ],
});
