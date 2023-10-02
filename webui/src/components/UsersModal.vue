<script>
export default {
	props: ['header', 'url'],
	data: function() {
		return {
			loading: false,
			errormsg: null,
			users: [],
		}
	},
	methods: {
		close() {
			this.errormsg = null
		},
		async getUsers() {
			this.loading = true;
			this.errormsg = null;

			try {
				let response = await this.$axios.get(this.url);
				this.users = response.data;
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
			this.loading = false;
		}
	},
	mounted() {
		this.getUsers()
	}
}
</script>

<template>
	<!-- Upload modal -->
	<div class="modal fade" :id="this.header" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
		<div class="modal-dialog" role="document">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLongTitle">{{this.header}}</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="close"></button>
				</div>
				<div class="modal-body">
					<LoadingSpinner :loading="loading"></LoadingSpinner>
					<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
					<div v-else class="d-flex flex-column gap-2">
						<router-link
							v-for="user in this.users"
							:key="user.identifier"
							:to="'/profile/' + user.identifier"
							class="text-decoration-none text-dark border px-2 py-1 rounded"
						>
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#user"/>
							</svg>
							{{ user.name }}
						</router-link>
						<h6 v-if="users.length === 0 && !loading && !errormsg">
							There's nothing here
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#frown"/>
							</svg>
						</h6>
					</div>
				</div>
			</div>
		</div>
	</div>

</template>
