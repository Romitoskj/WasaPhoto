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
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false
			this.$router.push('/')
		}
	}
}
</script>

<template>
	<LoadingSpinner v-if="loading"></LoadingSpinner>
	<ErrorMsg v-else-if="errormsg" :msg="errormsg"></ErrorMsg>
	<div v-else class="container d-flex min-vh-100 align-items-center justify-content-center">
		<div class="card w-50">
			<div class="card-header">
				<h3 class="card-title mt-2">Login</h3>
			</div>
			<div class="card-body">
				<form @submit.prevent="login" class="d-flex flex-column gap-4">
					<div class="form-group">
						<label for="username">Username</label>
						<input
							type="text"
							v-model="this.name"
							class="form-control"
							placeholder="Your username"
							required
							minlength="3"
							maxlength="16"
						>
					</div>
					<button type="submit" class="btn btn-primary">Login</button>
				</form>
			</div>
		</div>
	</div>
</template>
