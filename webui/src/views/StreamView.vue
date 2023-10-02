<script>
import {RouterLink} from "vue-router";
import Post from "../components/Post.vue";

export default {
	components: {Post},
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
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
			<h1>Welcome {{ this.$session.username }}!</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<!-- Button for upload photo modal -->
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
			<LoadingSpinner :loading="loading"></LoadingSpinner>

			<div class="container d-flex flex-column min-vh-100 align-items-center" v-if="stream.length === 0 && !loading && !errormsg">
				<h4>
					There's nothing here
					<svg class="feather">
						<use href="/feather-sprite-v4.29.0.svg#frown"/>
					</svg>
				</h4>
				<h6>You need to follow someone to see posts in your stream</h6>
			</div>

			<Post :photo="photo" v-for="photo in stream" :key="photo.identifier"></Post>
		</div>
	</div>

	<UploadModal></UploadModal>

</template>
