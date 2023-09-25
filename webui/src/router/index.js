import {createRouter, createWebHashHistory} from 'vue-router'
import Stream from "../views/StreamView.vue";
import LoginView from "../views/LoginView.vue";
import UploadView from "../views/UploadView.vue";
import SearchView from "../views/SearchView.vue";
import ProfileView from "../views/ProfileView.vue";

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView, name:'login'},
		{path: '/stream', component: Stream},
		{path: '/upload-photo', component: UploadView},
		{path: '/search', component: SearchView},
		{path: '/profile/:username', component: ProfileView},
	]
})

export default router
