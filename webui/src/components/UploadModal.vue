<script>

export default {
	data: function() {
		return {
			loading: false,
			errormsg: null,
			uploadSuccessful: false,
		}
	},
	methods: {
		close() {
			this.errormsg = null
			this.uploadSuccessful = false
		},
		async uploadPhoto() {
			this.loading = true;
			this.errormsg = false;
			const photo = this.$refs.photo.files[0]
			if (photo) {
				try {
					await this.$axios.post("/users/" + this.$session.id + "/photos/", photo)
					this.errormsg = null
					this.uploadSuccessful = true
					this.$refs.upload.reset()
				} catch (e) {
					if (e.response.data !== undefined) {
						this.errormsg = e.response.data.message
					} else {
						this.errormsg = e.toString()
					}
				}
				this.loading = false;
			} else {
				this.uploadSuccessful = false
				this.errormsg = "No photo selected."
			}
		}
	},
}
</script>

<template>
	<!-- Upload modal -->
	<div class="modal fade" id="uploadModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
		<div class="modal-dialog" role="document">
			<div class="modal-content">
				<form class="d-flex flex-column gap-3" ref="upload" @submit.prevent="uploadPhoto">
					<div class="modal-header">
						<h5 class="modal-title" id="exampleModalLongTitle">Upload photo</h5>
						<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="close"></button>
					</div>
					<div class="modal-body">
						<div class="alert alert-success" v-if="uploadSuccessful">Photo successfully uploaded!</div>
						<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
						<div class="d-flex flex-column gap-3 p-3">
							<input type="file" ref="photo" accept="image/jpeg">
						</div>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-secondary" data-bs-dismiss="modal" @click="close">Close</button>
						<button type="submit" class="btn btn-primary">Add Photo</button>
					</div>

				</form>
			</div>
		</div>
	</div>

</template>
