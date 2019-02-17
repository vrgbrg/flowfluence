import Vue from 'vue';
import VueRouter from 'vue-router';
import Login from './components/Login.vue'
import Articles from './components/Articles.vue'
import MyArticles from './components/MyArticles.vue'
import Article from './components/Article.vue'


Vue.use(VueRouter);
// TODO auth
// Vue.use(Auth)

export default new VueRouter({
    mode: 'history',
    routes: [
        // TODO login
        { name: 'root', path: '/', component: Login },
        { name: 'articles', path: '/articles', component: Articles },
        { name: 'my-articles', path: '/my-articles', component: MyArticles },
        { name: 'new', path: '/new', component: Article },
        { name: 'edit',path: '/edit/:id', component: Article },
    ]
});