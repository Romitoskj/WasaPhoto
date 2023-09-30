<script>
import Post from "../components/Post.vue";
import LoadingSpinner from "../components/LoadingSpinner.vue";

export default {
	components: {LoadingSpinner, Post},
	data: function() {
		return {
			errormsg: null,
			loading: false,
			user_id: null,
			profile: null,
		}
	},
	methods: {
		async getProfile() {
			this.loading = true;
			this.errormsg = null;

			this.user_id = this.$route.params.id

			try {
				let response = await this.$axios.get("/users/" + this.user_id);
				this.profile = response.data;
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

		console.log("starting mounting")
		if (this.$session.id === -1) {
			this.$router.push('/login')
		} else {
			this.getProfile()
		}
	}
}
</script>

<template>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	<LoadingSpinner :loading="loading"></LoadingSpinner>
	<div v-if="profile">
		<!-- Page header -->
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">{{profile.name}}</h1>
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
<!--			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>-->
<!--			<LoadingSpinner :loading="loading"></LoadingSpinner>-->

<!--			<div class="container d-flex flex-column min-vh-100 align-items-center" v-if="stream.length === 0 && !loading && !errormsg">-->
<!--				<h4>-->
<!--					There's nothing here-->
<!--					<svg class="feather">-->
<!--						<use href="/feather-sprite-v4.29.0.svg#frown"/>-->
<!--					</svg>-->
<!--				</h4>-->
<!--				<h6>You need to follow someone to see posts in your stream</h6>-->
<!--			</div>-->

<!--			<Post :photo="photo" v-for="photo in stream"></Post>-->
		</div>
	</div>
</template>
