<script>
export default {
	data() {
		return {
			name: "",
			errormsg: null,
			loading: false,
		}
	},
	methods: {
		async login() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session", {name: this.name})
				this.$session.login(response.data.identifier, this.name)
				this.$axios.defaults.headers.common['Authorization'] = 'Bearer ' + response.data.identifier

				this.$router.push('/')
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
			this.loading = false
		}
	}
}
</script>

<template>
	<div class="container d-flex min-vh-100 align-items-center justify-content-center">

		<LoadingSpinner :loading="loading"></LoadingSpinner>

		<!-- Login card -->
		<div v-if="!loading" class="card w-50">
			<div class="card-header">
				<h3 class="card-title mt-2">Login</h3>
			</div>
			<div class="card-body">
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<form @submit.prevent="login" class="d-flex flex-column gap-4">
					<div class="form-group">
						<h6><label for="username">Username:</label></h6>
						<input
							type="text"
							v-model="this.name"
							class="form-control"
							placeholder="Your username"
							required
						>
					</div>
					<button type="submit" class="btn btn-primary">Login</button>
				</form>
			</div>
		</div>

	</div>
</template>
