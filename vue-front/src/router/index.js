import Vue from 'vue'
import VueRouter from 'vue-router'
import login from '@/views/Login_page.vue';
import home from '@/views/Home_page.vue';
import page1 from '@/views/home/Page1_page.vue';
import page2 from '@/views/home/Page2_page.vue';
Vue.use(VueRouter)

const routes = [
  {path: '*', redirect: {name: 'login'},},
  {path : '/', name : 'login', component : login, meta:{ requireLogin:false}},
  {path : '/home', name : 'home', component : home, meta:{ requireLogin:true}, redirect:{name:'page1'} , 
  children:[
    {path : '/page1', name : 'page1', component : page1, meta:{ requireLogin:true} },
    {path : '/page2', name : 'page2', component : page2, meta:{ requireLogin:true} },
  ]},
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})


router.beforeEach((to, from, next) => {
  const ls = localStorage.getItem('token');
  if (to.matched.some((record) => record.meta.requireLogin)) {
    if(ls=== undefined || ls === null){
      next({name: 'login'})
    }else{
      next()
    }
  }else{
    if(ls=== undefined || ls === null){
      next()
    }
    else{
      if (to.name==='login'){
        next({name: 'home'})
      }else{
        next()
      }
    }
  }
})

export default router
