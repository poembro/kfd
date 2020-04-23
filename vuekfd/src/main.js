import Vue from 'vue'
import App from './App.vue'
import router from './router'
import element from './plugins/element'
//导入全局样式表
import './assets/css/global.css'

import axios from 'axios'
axios.defaults.baseURL = 'http://www.591pubg.com/'
Vue.prototype.$http = axios

Vue.config.productionTip = false

new Vue({
  element,
  router,
  render: h => h(App),
}).$mount('#app')
