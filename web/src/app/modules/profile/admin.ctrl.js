(function() {
	'use strict';

	angular.module('app.profile').controller("AdminController", ['UserAccess', '$scope', '$currentUser', AdminCtrl]);

	function AdminCtrl(UserAccess, $scope, $currentUser){
		var adminVm = this;

		adminVm.addAdmin = addAdmin;
		adminVm.removeAdmin = removeAdmin;
		adminVm.currentUser = $currentUser.GetCurrentUser();
		adminVm.profileUser = $scope.profileUser;
		adminVm.IsAdmin = adminVm.profileUser.IsAdmin;
		adminVm.IsNotAdmin = !adminVm.profileUser.IsAdmin;

		/////////////////////////////////////
		function addAdmin() {
			UserAccess.addGroup($scope.profileUser.Id, "admin").then(updateData);
		}

		function removeAdmin() {
			UserAccess.removeGroup($scope.profileUser.Id, "admin").then(updateData);
		}

		function updateData(data) {
			var user = data;
			$scope.profileUser = user;
			$currentUser.SetCurrentUser(user);
			adminVm.IsAdmin = user.IsAdmin;
			adminVm.IsNotAdmin = !user.IsAdmin;
		}
	}

})();
