<script>

</script>

<template>
	<div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Upload photo</h1>
	</div>
	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">
		<form class="d-flex flex-column gap-3 card w-50" @submit.prevent="addPhoto">
			<label for="photo" class="card-header">Select a Photo</label>
			<div class="d-flex flex-column gap-3 p-3">
				<input type="file" ref="photo" accept="image/jpeg">
				<button type="submit" class="btn btn-outline-primary">Add Photo</button>
			</div>
		</form>
	</div>
<!--	<div class="container d-flex flex-column min-vh-100 align-items-center my-5 gap-3">-->
<!--		<h4 class="border rounded">-->
<!--			<input-->
<!--				type="text"-->
<!--				class="d-flex px-2 py-2"-->
<!--				v-model="this.query"-->
<!--				style="outline: none; border: none; background: none"-->
<!--				placeholder="Type an user name"-->
<!--			>-->
<!--		</h4>-->
<!--		<LoadingSpinner v-if="loading"></LoadingSpinner>-->
<!--		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>-->
<!--		<h5 v-if="users.length === 0 && query.length > 0">No user found</h5>-->
<!--		<div class="d-flex flex-column gap-2 w-50" v-else-if="query.length > 0">-->
<!--			<router-link-->
<!--				v-for="user in users"-->
<!--				:key="user.identifier"-->
<!--				:to="'/profile/' + user.identifier"-->
<!--				class="text-decoration-none text-dark border px-2 py-1 rounded"-->
<!--			>-->
<!--				<svg class="feather">-->
<!--					<use href="/feather-sprite-v4.29.0.svg#user"/>-->
<!--				</svg>-->
<!--				{{ user.name }}-->
<!--			</router-link>-->
<!--		</div>-->
<!--	</div>-->
</template>
