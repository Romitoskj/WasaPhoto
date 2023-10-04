<script>
import displayDateAndTime from "../services/date";

export default {
	props: ['comment', 'photo_author', 'photo'],
	data: function() {
		return {
			deleted: false,
		}
	},
	methods: {
		displayDateAndTime,
		async deleteComment(id) {
			this.errormsg = null;

			try {
				await this.$axios.delete(`/users/${this.photo_author}/photos/${this.photo}/comments/${id}`)
				this.deleted = true
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		}
	}
}
</script>

<template>
	<div v-if="deleted" class="alert alert-success alert-dismissible fade show" role="alert">
		Comment deleted successfully!
		<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
	</div>
	<div v-else class="d-flex align-items-center justify-content-between border px-2 py-1 rounded">
		<div class="d-flex flex-column gap-2">
			<div class="text-decoration-none text-dark">
				<svg class="feather">
					<use href="/feather-sprite-v4.29.0.svg#user"/>
				</svg>
				{{ this.comment.author.name }}
				<small class="text-muted">&nbsp;- {{ displayDateAndTime(new Date(Date.parse(this.comment.created_at))) }}</small>
			</div>
			<h6>
				{{ this.comment.content }}
			</h6>
		</div>

		<div class="d-flex align-items-center" v-if="this.comment.author.identifier === this.$session.id">
			<button @click="deleteComment(this.comment.identifier)" type="button" class="btn btn-link link-danger align-items-center">
				<svg class="feather">
					<use href="/feather-sprite-v4.29.0.svg#trash-2"/>
				</svg>
			</button>
		</div>
	</div>
</template>
