import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import Login from "./views/LoginView.vue";

import './assets/dashboard.css'
import './assets/main.css'

const session = {
	id: -1,
	username: "",

	login(id, username) {
		this.id = id;
		this.username = username;
	},

	logout() {
		this.id = -1;
		this.username = "";
	}
}

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$session = session;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("Login", Login)
app.use(router)
app.mount('#app')
