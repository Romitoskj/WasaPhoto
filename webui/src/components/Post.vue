<script>
export default {
	props: {
		photo: {type: Object}
	},
	data() {
		return {
			img : "",

			loading: false,
			errormsg: "",
		}
	},
	methods: {
		async getPhoto(photo) {
			this.loading = true;
			this.errormsg = null;

			try {
				let response = await this.$axios.get(`/users/${photo.author.identifier}/photos/${photo.identifier}`, {responseType: "blob"})
				this.img = URL.createObjectURL(response.data)
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
		this.getPhoto(this.$props.photo)
	}
}
</script>

<template>
	<!-- Post card -->
	<div class="card mb-5" style="width: 30rem;">

		<!-- Card header -->
		<div class="card-header">
			<h5 class="card-title d-inline-block">{{this.$props.photo.author.name}}</h5>
			<p class="card-text">{{ new Date(Date.parse(this.$props.photo.created_at)).toDateString() }}</p>
		</div>

		<!-- Image div -->
		<div class="card-body">
			<LoadingSpinner :loading="loading" />
			<img
				:src="img"
				:alt='"photo ID: " + this.$props.photo.identifier'
				class="card-img"
			>
		</div>

<!--		<div class="container">-->
<!--			<div class="row">-->

<!--				&lt;!&ndash; Username and date &ndash;&gt;-->
<!--				<div class="col-10">-->
<!--					<div class="card-body">-->
<!--						<h5 @click="visitUser" class="card-title d-inline-block" style="cursor: pointer">{{ name }}</h5>-->
<!--						<p class="card-text">{{ new Date(Date.parse(date)) }}</p>-->
<!--					</div>-->
<!--				</div>-->

<!--				&lt;!&ndash; Comment and like buttons &ndash;&gt;-->
<!--				<div class="col-2">-->
<!--					<div class="card-body d-flex justify-content-end" style="display: inline-flex">-->
<!--						<a @click="showHideComments">-->
<!--							<h5><i class="card-title bi bi-chat-right pe-1"></i></h5>-->
<!--						</a>-->
<!--						<h6 class="card-text d-flex align-items-end text-muted">{{ post_comments_cnt }}</h6>-->
<!--						<a v-if="!post_liked" @click="like">-->
<!--							<h5><i class="card-title bi bi-suit-heart ps-2 pe-1 like-icon"></i></h5>-->
<!--						</a>-->
<!--						<a v-if="post_liked" @click="unlike">-->
<!--							<h5><i class="card-title bi bi-heart-fill ps-2 pe-1 like-icon like-red"></i></h5>-->
<!--						</a>-->
<!--						<h6 class="card-text d-flex align-items-end text-muted">{{ post_like_cnt }}</h6>-->
<!--						<h5></h5>-->
<!--					</div>-->
<!--				</div>-->
<!--			</div>-->

<!--			&lt;!&ndash; Comments section &ndash;&gt;-->
<!--			<div v-if="comments_shown">-->
<!--				<div v-for="item of comments_data" class="row" v-bind:key="item.comment_id">-->
<!--					<div class="col-7 card-body border-top">-->
<!--						<b>{{ item.name }}:</b> {{ item.comment }}-->
<!--					</div>-->
<!--					<div class="col-5 card-body border-top text-end text-secondary">-->
<!--						{{ new Date(Date.parse(item.date)).toDateString() }}-->
<!--					</div>-->
<!--				</div>-->

<!--				&lt;!&ndash; Show more comments label &ndash;&gt;-->
<!--				<div v-if="!data_ended" class="col-12 card-body text-end pt-0 pb-1 px-0">-->
<!--					<a @click="getComments" class="text-primary">Show more comments...</a>-->
<!--				</div>-->

<!--				&lt;!&ndash; New comment form &ndash;&gt;-->
<!--				<div class="row">-->

<!--					&lt;!&ndash; Comment input &ndash;&gt;-->
<!--					<div class="col-10 card-body border-top text-end">-->
<!--						<input v-model="commentMsg" type="text" class="form-control" placeholder="Commenta...">-->
<!--					</div>-->

<!--					&lt;!&ndash; Comment publish button &ndash;&gt;-->
<!--					<div class="col-1 card-body border-top text-end ps-0 d-flex">-->
<!--						<button style="width: 100%" type="button" class="btn btn-primary"-->
<!--								@click="postComment">Go</button>-->
<!--					</div>-->
<!--				</div>-->
<!--			</div>-->
<!--		</div>-->
	</div>
</template>
