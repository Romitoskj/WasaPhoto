<script>
import Comment from "./Comment.vue";
import comment from "./Comment.vue";

export default {
	computed: {
		comment() {
			return comment
		}
	},
	components: {Comment},
	props: ['id', 'comments', 'photo_author', 'photo'],
	data: function() {
		return {
			errormsg: null,
			newComment: null,
		}
	},
	methods: {
		clearComment() {
			this.newComment = null
			this.errormsg = null
		},
		async addComment() {
			try {
				let response = await this.$axios.post(`/users/${this.photo_author}/photos/${this.photo}/comments/`, {'content': this.newComment})
				this.errormsg = null
				this.comments.push(response.data)
				this.newComment = null
			} catch (e) {
				if (e.response.data !== undefined) {
					this.errormsg = e.response.data.message
				} else {
					this.errormsg = e.toString()
				}
			}
		},
	},
}
</script>

<template>
	<!-- Comments modal -->
	<div class="modal fade" :id="id" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
		<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
			<div class="modal-content" style="height: 40%">

				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLongTitle">Comments</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="clearComment"></button>
				</div>

				<div class="modal-body">
					<div class="d-flex flex-column gap-2">
						<Comment
							v-for="comment in comments"
							:key="comment.identifier"
							:comment="comment"
							:photo="photo"
							:photo_author="photo_author"
						></Comment>

						<h6 v-if="comments.length === 0">
							There's nothing here
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#frown"/>
							</svg>
						</h6>
					</div>
				</div>

				<div class="modal-footer d-flex flex-column">
					<form @submit.prevent="addComment" class="d-flex align-items-center justify-content-between gap-1" style="width: 100%;">
						<input
							type="text"
							v-model="newComment"
							class="form-control d-flex"
							placeholder="Type comment..."
							required
						>
						<button type="submit" class="btn btn-outline-primary d-flex align-items-center">
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#send"/>
							</svg>
							Add
						</button>
					</form>
					<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				</div>
			</div>
		</div>
	</div>

</template>
