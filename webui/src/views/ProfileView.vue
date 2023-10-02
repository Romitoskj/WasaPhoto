<script>
import Post from "../components/Post.vue";

export default {
	components: {Post},
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
		if (this.$session.id === -1) {
			this.$router.push('/login')
		} else {
			this.getProfile()
		}
	}
}
</script>

<template>

	<!-- Page header -->
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1>User profile</h1>
	</div>

	<!-- Main content -->
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">

		<div class="card" style="width: 100%" v-if="errormsg || loading">
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner :loading="loading"></LoadingSpinner>
		</div>

		<!--- User info -->
		<div class="card" style="width: 100%" v-if="profile">
			<div class="card-header d-flex justify-content-between">

				<div class="card-body d-flex align-items-center gap-1">
					<h2 class="card-title d-inline-block">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#user"/>
						</svg>
						{{ profile.name }}
					</h2>
				</div>

				<div class="card-body d-flex align-items-center justify-content-end gap-1">
					<button type="button" class="btn btn-link d-flex align-items-end link-dark">Followers: {{ profile.followers_n }}</button>
					<button type="button" class="btn btn-link d-flex align-items-end link-dark">Following: {{ profile.following_n }}</button>
					<button  type="button" class="btn btn-link d-flex align-items-end link-dark" style="pointer-events: none">Photos: {{profile.photos_n}}</button>
				</div>

				<div class="card-body d-flex align-items-center justify-content-end gap-1">
					<button v-if="profile.followed" type="button" class="btn btn-primary d-flex align-items-end">Unfollow</button>
					<button v-else type="button" class="btn btn-outline-primary d-flex align-items-end">Follow</button>
					<button v-if="profile.banned" type="button" class="btn btn-danger d-flex align-items-end">Unban</button>
					<button v-else type="button" class="btn btn-outline-danger d-flex align-items-end">Ban</button>
				</div>
			</div>


			<div class="card-body d-flex flex-column align-items-center" v-if="profile.photos.length === 0 && !loading && !errormsg">
				<h4>
					There's nothing here
					<svg class="feather">
						<use href="/feather-sprite-v4.29.0.svg#frown"/>
					</svg>
				</h4>
			</div>

			<div class="card-body d-flex flex-column align-items-center">
				<Post :photo="photo" v-for="photo in profile.photos" :key="photo.identifier"></Post>
			</div>
		</div>

	</div>
</template>
