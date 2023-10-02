<script>
import displayDateAndTime from "../services/date";

export default {
	props: {
		photo: {type: Object}
	},
	data() {
		return {
			img : "",

			loading: false,
			photoError: null,
			errormsg: null,
		}
	},
	methods: {
		displayDateAndTime,
		async getPhoto(photo) {
			this.loading = true;
			this.photoError = null;

			try {
				let response = await this.$axios.get(`/users/${photo.author.identifier}/photos/${photo.identifier}`, {responseType: "blob"})
				this.img = URL.createObjectURL(response.data)
			} catch (e) {
				let response = JSON.parse(await e.response.data.text())
				if (e.response.data !== undefined) {
					this.photoError = response.message
				} else {
					this.photoError = e.toString()
				}
			}


			this.loading = false;
		},
		async like(photo) {
			this.errormsg = null;

			try {
				let response = await this.$axios.put(`/users/${photo.author.identifier}/photos/${photo.identifier}/likes/${this.$session.id}`)
				photo.liked = true
				photo.likes_n++
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async unlike(photo) {
			this.errormsg = null;

			try {
				let response = await this.$axios.delete(`/users/${photo.author.identifier}/photos/${photo.identifier}/likes/${this.$session.id}`)
				photo.liked = false
				photo.likes_n--
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
	},
	mounted() {
		this.getPhoto(this.$props.photo)
	}
}
</script>

<template>
	<!-- Post card -->
	<div class="card mb-5" style="width: 60%">

		<!-- Image div -->

		<LoadingSpinner :loading="loading" />

		<ErrorMsg v-if="photoError" :msg="photoError"></ErrorMsg>
		<img v-else :src="img" class="card-img-top" :alt="photoError">

		<div class="container">
			<div class="row">

				<!-- Username and date -->
				<div class="col-8">
					<div class="card-body">
						<h5 class="card-title d-inline-block" style="cursor: pointer">
							<RouterLink :to="'/profile/' + this.photo.author.identifier" class="link-dark">
								{{ this.photo.author.name }}
							</RouterLink>
						</h5>
						<small class="text-muted"> - {{ displayDateAndTime(new Date(Date.parse(this.photo.created_at))) }}</small>
					</div>
				</div>

				<!-- Comment and like buttons -->
				<div class="col-4">
					<div class="card-body d-flex justify-content-end gap-3">

						<!-- Comments icon an count -->
						<div class="d-flex align-items-center gap-1">
							<svg @click="showCommentForm" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-chat-right" viewBox="0 0 16 16" style="cursor: pointer">
								<path d="M2 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9.586a2 2 0 0 1 1.414.586l2 2V2a1 1 0 0 0-1-1H2zm12-1a2 2 0 0 1 2 2v12.793a.5.5 0 0 1-.854.353l-2.853-2.853a1 1 0 0 0-.707-.293H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12z"/>
							</svg>
							<button @click="showComments" type="button" class="btn btn-link d-flex align-items-end text-muted">{{ this.photo.comments_n }}</button>
						</div>

						<!-- Likes icon and count -->

						<div class="d-flex align-items-center gap-1">
							<svg v-if="this.photo.liked" @click="unlike(this.photo)" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-heart-fill" viewBox="0 0 16 16" style="cursor: pointer">
								<path fill-rule="evenodd" d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314z"/>
							</svg>

							<svg v-else @click="like(this.photo)" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-heart" viewBox="0 0 16 16" style="cursor: pointer">
								<path d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z"/>
							</svg>
							<button type="button" class="btn btn-link d-flex align-items-end text-muted">{{ this.photo.likes_n }}</button>
						</div>
					</div>
				</div>


			</div>
		</div>
		<div class="card-body" v-if="errormsg">
			<ErrorMsg :msg="errormsg"></ErrorMsg>
		</div>
	</div>
</template>
