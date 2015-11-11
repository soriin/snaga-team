(function() {
	'use strict';

	angular.module('app').controller('ProfileCtrl', ['$window', '$state', 'UserAccess', '$currentUser', ProfileCtrl]);

	function ProfileCtrl($window, $state, UserAccess, $currentUser) {
		var profileVm = this;

		profileVm.someText = "HI, " + $window.gapi.auth2.getAuthInstance().currentUser.get().getBasicProfile().getName();
		profileVm.update = update;

		activate();

		///////////////// Functions ////////////////////////

		function activate() {
			UserAccess.createUser().then(updateData);
		}

		function update() {
			var currentUser = $currentUser.GetCurrentUser();
			UserAccess.updateUser(currentUser.Id, {
				FirstName : currentUser.FirstName,
				LastName : currentUser.LastName,
				DisplayName : currentUser.DisplayName,
				InGameName : currentUser.InGameName,
				Email : currentUser.Email
			}).then(updateData);
		}

		function updateData(data) {
			var user = data;
			profileVm.User = user;
			$currentUser.SetCurrentUser(user);
		}
	}
})();
