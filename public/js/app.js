import Vue from 'vue';
import AppComponent from './components/app.vue';
import css from "./app.css";

Vue.config.productionTip = false

new Vue({
    css: css,
    render: h => h(AppComponent),
}).$mount('#app') 