(function() {
	'use strict';

	angular.module('app').controller('ProfileCtrl', ['$scope', '$window', '$state', 'Users', 'Ships', '$currentUser', ProfileCtrl]);

	function ProfileCtrl($scope, $window, $state, Users, Ships, $currentUser) {
		if ($window.gapi.auth2 == undefined || $window.gapi.auth2.getAuthInstance().isSignedIn.get() == false) {
			$state.go("login");
			return;
		}

		$scope.someText = "HI, " + $window.gapi.auth2.getAuthInstance().currentUser.get().getBasicProfile().getName();
		$scope.update = update;

		var me = Users.save({}, function() {
			var user = me;

			$scope.User = user;
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
