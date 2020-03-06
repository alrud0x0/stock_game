import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home.vue'
import Rank from '../components/Rank.vue'
import CurrentGame from '../components/CurrentGame.vue'

Vue.use(VueRouter)

//라우트 정의
const routes = [
  { path : '/home', component: Home, name: 'home' },
  { path : '/rank', component: Rank, name: 'rank' },
  { path : '/game/current', component: CurrentGame, name: 'currentGame' },
  { path : '/game/history', component: CurrentGame, name: 'gameHistory' }
]

//router 인스턴스 생성
export const router = new VueRouter({
  base: '/',
  routes: routes
})
