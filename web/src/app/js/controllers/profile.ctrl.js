(function() {
	'use strict';

	angular.module('app').controller('ProfileCtrl', ['$window', '$state', 'Users', '$currentUser', ProfileCtrl]);

	function ProfileCtrl($window, $state, Users, $currentUser) {
		var profileVm = this;
		if ($window.gapi.auth2 == undefined || $window.gapi.auth2.getAuthInstance().isSignedIn.get() == false) {
			$state.go("login");
			return;
		}

		profileVm.someText = "HI, " + $window.gapi.auth2.getAuthInstance().currentUser.get().getBasicProfile().getName();
		profileVm.update = update;

		var me = Users.save({}, function() {
			var user = me;

			profileVm.User = user;
			$currentUser.SetCurrentUser(user);
		});

		function update() {
			var currentUser = $currentUser.GetCurrentUser();
			Users.update({id: currentUser.Id},
			{
				FirstName : currentUser.FirstName,
				LastName : currentUser.LastName,
				DisplayName : currentUser.DisplayName,
				InGameName : currentUser.InGameName,
				Email : currentUser.Email
			})
		};
	}
})();
