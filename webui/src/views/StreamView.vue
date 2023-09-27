<script>
import {RouterLink} from "vue-router";
import LoadingSpinner from "../components/LoadingSpinner.vue";

export default {
	components: {LoadingSpinner, RouterLink},
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
			uploadError: null,
		}
	},
	methods: {
		async getStream() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.$session.id + "/stream");
				this.stream = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async uploadPhoto() {
			this.loading = true;
			this.uploadError = false;
			const photo = this.$refs.photo.files[0]
			if (photo) {
				try {
					await this.$axios.post("/users/" + this.$session.id + "/photos/")
					this.uploadError = null
				} catch (e) {
					this.uploadError = e.toString();
				}
				this.loading = false;
			}
		}
	},
	mounted() {
		if (this.$session.id === -1) {
			this.$router.push('/login')
		} else {
			this.getStream()
		}
	}
}
</script>

<template>
	<div>
		<!-- Page header -->
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Welcome {{ this.$session.username }}!</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<!-- Button trigger for upload photo modal -->
					<button  type="button"  class="btn btn-sm btn-outline-primary" data-bs-toggle="modal" data-bs-target="#uploadModal">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#upload"/>
						</svg>
						Upload Photo
					</button>
				</div>
			</div>
		</div>

		<!-- Main content -->
		<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<div class="container d-flex flex-column min-vh-100 align-items-center" v-if="stream.length === 0">
				<h4>
					There's nothing here
					<svg class="feather">
						<use href="/feather-sprite-v4.29.0.svg#frown"/>
					</svg>
				</h4>
				<h6>You need to follow someone to see posts in your stream</h6>
			</div>

			<div v-for="photo in stream">
				{{ photo }}
			</div>
		</div>
	</div>

	<!-- Upload modal -->
	<div class="modal fade" id="uploadModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
		<div class="modal-dialog" role="document">
			<div class="modal-content">
				<form class="d-flex flex-column gap-3" @submit.prevent="uploadPhoto">
					<div class="modal-header">
						<h5 class="modal-title" id="exampleModalLongTitle">Upload photo</h5>
						<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
					</div>
					<div class="modal-body">
							<div class="d-flex flex-column gap-3 p-3">
								<input type="file" ref="photo" accept="image/jpeg">
							</div>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-outline-secondary" data-bs-dismiss="modal">Close</button>
						<button type="submit" class="btn btn-primary">Add Photo</button>
					</div>

				</form>
			</div>
		</div>
	</div>

</template>
