<script>
export default {
	props: ['id', 'header', 'users'],
	data: function() {
		return {

		}
	}
}
</script>

<template>
	<!-- Users modal -->
	<div class="modal fade" :id="id" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">
		<div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
			<div class="modal-content" style="height: 40%">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLongTitle">{{header}}</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<div class="d-flex flex-column gap-2">
						<div
							v-for="user in users"
							:key="user.identifier"
							class="text-decoration-none text-dark "
						>
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#user"/>
							</svg>
							{{ user.name }}
						</div>
						<h6 v-if="users.length === 0">
							There's nothing here
							<svg class="feather">
								<use href="/feather-sprite-v4.29.0.svg#frown"/>
							</svg>
						</h6>
					</div>
				</div>
			</div>
		</div>
	</div>

</template>
