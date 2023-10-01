<script>
import LoadingSpinner from "../components/LoadingSpinner.vue";

export default {
	components: {LoadingSpinner},
	data() {
		return {
			users: [],
			query: "",
			errormsg: null,
			loading: false,
		}
	},
	mounted() {
		if (this.$session.id === -1) {
			this.$router.push('/login')
		}
	},
	methods: {
		async refresh(searchQuery) {
			this.loading = true
			this.errormsg = null
			try {
				if (this.query.length > 0) {
					let response = await this.$axios.get("/users/?search=" + searchQuery)
					this.users = response.data
				} else {
					this.users = []
				}
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
			this.loading = false
		},
	},
	watch: {
		query(newQuery) {
			this.refresh(newQuery)
		}
	}
}

</script>

<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>Search user</h1>
	</div>
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<h4 class="border rounded">
			<input
				type="text"
				class="d-flex px-2 py-2"
				v-model="this.query"
				style="outline: none; border: none; background: none"
				placeholder="Type an user name"
			>
		</h4>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<h5 v-if="users.length === 0 && query.length > 0">No user found</h5>
		<div class="d-flex flex-column gap-2 w-50" v-else-if="query.length > 0">
			<router-link
				v-for="user in users"
				:key="user.identifier"
				:to="'/profile/' + user.identifier"
				class="text-decoration-none text-dark border px-2 py-1 rounded"
			>
				<svg class="feather">
					<use href="/feather-sprite-v4.29.0.svg#user"/>
				</svg>
				{{ user.name }}
			</router-link>
		</div>
		<LoadingSpinner :loading="loading"></LoadingSpinner>
	</div>
</template>
