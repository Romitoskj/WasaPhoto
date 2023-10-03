<script>
import Post from "../components/Post.vue";
import UsersModal from "../components/UsersModal.vue";

export default {
	components: {UsersModal, Post},
	data: function() {
		return {
			errormsg: null,
			loading: false,
			user_id: null,
			profile: null,
			changing: false,
			newName: null,
			users: [],
		}
	},
	methods: {
		async getProfile() {
			this.loading = true;
			this.errormsg = null;

			this.user_id = parseInt(this.$route.params.id)

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
		},
		async follow() {
			this.errormsg = null;

			try {
				let response = await this.$axios.put(`/users/${this.user_id}/followers/${this.$session.id}`)
				this.profile.followed = true
				this.profile.followers_n ++
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async unfollow() {
			this.errormsg = null;

			try {
				let response = await this.$axios.delete(`/users/${this.user_id}/followers/${this.$session.id}`)
				this.profile.followed = false
				this.profile.followers_n --
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async ban() {
			this.errormsg = null;

			try {
				let response = await this.$axios.put(`/users/${this.$session.id}/banned/${this.user_id}`)
				this.profile.banned = true
				if (this.profile.followed) {
					this.profile.followed = false
					this.profile.followers_n --
				}
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async unban() {
			this.errormsg = null;

			try {
				let response = await this.$axios.delete(`/users/${this.$session.id}/banned/${this.user_id}`)
				this.profile.banned = false
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		toggleChangeUsername() {
			this.newName = null
			this.changing = !this.changing
		},
		async changeUsername() {
			this.errormsg = null;

			try {
				let response = await this.$axios.put(`/users/${this.$session.id}/username`, {name: this.newName})
				this.profile.name = this.newName
				this.$session.username = this.newName
				this.toggleChangeUsername()
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
		async getUsers(url) {
			this.errormsg = null;

			try {
				let response = await this.$axios.get(url);
				this.users = response.data;
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
			<div class="card-body">
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner :loading="loading"></LoadingSpinner>
			</div>
		</div>

		<!--- User info -->
		<div class="card" style="width: 100%" v-if="profile">
			<div class="card-header d-flex justify-content-between">

				<!-- Change username form -->
				<div v-if="changing" class="card-body d-flex align-items-center">
					<form @submit.prevent="changeUsername" class="d-flex align-items-center gap-3">
						<input
							type="text"
							v-model="newName"
							class="form-control"
							placeholder="New username"
							required
							minlength="3"
							maxlength="16"
						>
						<div class="d-flex align-items-center justify-content-end gap-1">
							<button type="submit" class="btn btn-outline-primary d-flex align-items-end">Submit</button>
							<button type="button" class="btn btn-outline-secondary d-flex align-items-end" @click="toggleChangeUsername">Cancel</button>
						</div>
					</form>
				</div>

				<!-- Username -->
				<div v-else class="card-body d-flex align-items-center">
					<h2 class="card-title d-inline-block">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#user"/>
						</svg>
						{{ profile.name }}
					</h2>
					<button v-if="user_id === this.$session.id" type="button" class="btn btn-link link-warning d-flex align-items-center" @click="toggleChangeUsername">
						<svg class="feather">
							<use href="/feather-sprite-v4.29.0.svg#edit-3"/>
						</svg>
					</button>
				</div>

				<!-- Followers, following and photo count -->
				<div class="card-body d-flex align-items-center justify-content-end gap-1">
					<button type="button"
							class="btn btn-link d-flex align-items-end link-dark"
							@click="getUsers(`/users/${this.user_id}/followers/`)"
							data-bs-toggle="modal"
							data-bs-target="#Followers"
					>
						Followers: {{ profile.followers_n }}
					</button>
					<button
						type="button"
						class="btn btn-link d-flex align-items-end link-dark"
						@click="getUsers(`/users/${this.user_id}/following/`)"
						data-bs-toggle="modal"
						data-bs-target="#Following"
					>
						Following: {{ profile.following_n }}
					</button>
					<button  type="button" class="btn btn-link d-flex align-items-end link-dark text-decoration-none" style="pointer-events: none">Photos: {{profile.photos_n}}</button>
				</div>

				<!-- Follow/Unfollow and Ban/Unban buttons -->
				<div class="card-body d-flex align-items-center justify-content-end gap-1" v-if="user_id !== this.$session.id">
					<button v-if="profile.followed" @click="unfollow" type="button" class="btn btn-primary d-flex align-items-end">Unfollow</button>
					<button v-else-if="!profile.banned" @click="follow" type="button" class="btn btn-outline-primary d-flex align-items-end">Follow</button>
					<button v-if="profile.banned" @click="unban" type="button" class="btn btn-danger d-flex align-items-end">Unban</button>
					<button v-else @click="ban" type="button" class="btn btn-outline-danger d-flex align-items-end">Ban</button>
				</div>
			</div>

			<!-- Photos -->
			<div class="card-body d-flex flex-column align-items-center">
				<Post :photo="photo" v-for="photo in profile.photos" :key="photo.identifier"></Post>
				<h4 v-if="profile.photos.length === 0 && !loading && !errormsg">
					There's nothing here
					<svg class="feather">
						<use href="/feather-sprite-v4.29.0.svg#frown"/>
					</svg>
				</h4>
			</div>
		</div>

	</div>

	<div v-if="profile">
		<UsersModal header="Followers" :users="users"></UsersModal>
		<UsersModal header="Following" :users="users"></UsersModal>
	</div>
</template>
