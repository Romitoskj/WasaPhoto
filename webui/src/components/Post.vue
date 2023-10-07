<script>
import displayDateAndTime from "../services/date";
import UsersModal from "./UsersModal.vue";
import CommentsModal from "./CommentsModal.vue";

export default {
	components: {CommentsModal, UsersModal},
	props: {
		photo: {type: Object}
	},
	data() {
		return {
			img : "",
			deleted: false,

			loading: false,
			photoError: null,
			errormsg: null,

			likes: [],
			comments: [],
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
		async deletePhoto(photo) {
			this.errormsg = null;

			try {
				let response = await this.$axios.delete(`/users/${photo.author.identifier}/photos/${photo.identifier}`)
				this.deleted = true
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async getLikes() {
			this.errormsg = null;

			try {
				let response = await this.$axios.get(`/users/${this.photo.author.identifier}/photos/${this.photo.identifier}/likes/`);
				this.likes = response.data;
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async getComments() {
			this.errormsg = null;

			try {
				let response = await this.$axios.get(`/users/${this.photo.author.identifier}/photos/${this.photo.identifier}/comments/`);
				this.comments = response.data;
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		commentsCount(value) {
			this.photo.comments_n += value
		}
	},

	mounted() {
		this.getPhoto(this.$props.photo)
	}
}
</script>

<template>
	<!-- Post card -->
	<div v-if="deleted" class="alert alert-success alert-dismissible fade show" role="alert" style="width: 60%">
		Photo deleted successfully!
		<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
	</div>
	<div v-else class="card mb-5" style="width: 60%">

		<!-- Image div -->

		<LoadingSpinner :loading="loading" />

		<ErrorMsg v-if="photoError" :msg="photoError"></ErrorMsg>
		<img :src="img" class="card-img-top" :alt="photoError">

		<div class="d-flex justify-content-between">

			<!-- Username and date -->
			<div class="card-body d-flex align-items-center">
				<h5 class="card-title d-inline-block" style="cursor: pointer">
					<RouterLink :to="'/profile/' + photo.author.identifier" class="link-dark">
						{{ photo.author.name }}
					</RouterLink>
				</h5>
				<small class="text-muted">&nbsp;- {{ displayDateAndTime(new Date(Date.parse(photo.created_at))) }}</small>
			</div>

			<!-- Comment and like buttons -->
			<div class="card-body d-flex justify-content-end gap-1">

				<!-- Delete button -->
				<div class="d-flex align-items-center gap-1" v-if="$session.id === photo.author.identifier">
					<button @click="deletePhoto(photo)" type="button" class="btn btn-link link-danger align-items-center">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#trash-2"/>
						</svg>
						Delete
					</button>
				</div>

				<!-- Comments icon an count -->
				<div class="d-flex align-items-center">
					<button @click="getComments" type="button" class="btn btn-link d-flex align-items-center gap-3 text-dark" data-bs-toggle="modal" :data-bs-target="`#Comments${photo.identifier}`">
						<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-chat-right" viewBox="0 0 16 16">
							<path d="M2 1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h9.586a2 2 0 0 1 1.414.586l2 2V2a1 1 0 0 0-1-1H2zm12-1a2 2 0 0 1 2 2v12.793a.5.5 0 0 1-.854.353l-2.853-2.853a1 1 0 0 0-.707-.293H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h12z"/>
						</svg>
						<span class="text-muted">{{ photo.comments_n }}</span>
					</button>
				</div>

				<!-- Likes icon and count -->

				<div class="d-flex align-items-center gap-1">
					<svg v-if="photo.liked" @click="unlike(photo)" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-heart-fill" viewBox="0 0 16 16" style="cursor: pointer">
						<path fill-rule="evenodd" d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314z"/>
					</svg>

					<svg v-else @click="like(photo)" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" class="bi bi-heart" viewBox="0 0 16 16" style="cursor: pointer">
						<path d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01L8 2.748zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15z"/>
					</svg>
					<button type="button" @click="getLikes" class="btn btn-link d-flex align-items-center text-muted"  data-bs-toggle="modal" :data-bs-target="`#Likes${photo.identifier}`">{{ photo.likes_n }}</button>
				</div>
			</div>


		</div>

		<div class="card-body" v-if="errormsg">
			<ErrorMsg :msg="errormsg"></ErrorMsg>
		</div>
	</div>
	<UsersModal :id="`Likes${photo.identifier}`" header="Likes" :users="likes"></UsersModal>
	<CommentsModal :id="`Comments${photo.identifier}`" :comments="comments" :photo_author="photo.author.identifier" :photo="photo.identifier" @comments-count="commentsCount"></CommentsModal>
</template>
