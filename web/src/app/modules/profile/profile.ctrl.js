(function() {
	'use strict';

	angular.module('app.profile').controller('ProfileCtrl',
	['$window', '$state', 'UserAccess', '$rootScope', 'profileData', '$stateParams', ProfileCtrl]);

	function ProfileCtrl($window, $state, UserAccess, $rootScope, profileData, $stateParams) {
		var profileVm = this;

		// profileVm.someText = "HI, " + $window.gapi.auth2.getAuthInstance().currentUser.get().getBasicProfile().getName();
		profileVm.update = update;
		profileVm.user = {};

		activate(profileData);

		///////////////// Functions ////////////////////////

		function activate(data) {
			setProfileData(data);
		}

		function update() {
			UserAccess.updateUser(profileVm.user.Id, {
				FirstName : profileVm.user.FirstName,
				LastName : profileVm.user.LastName,
				DisplayName : profileVm.user.DisplayName,
				InGameName : profileVm.user.InGameName,
				Email : profileVm.user.Email
			}).then(setProfileData);
		}

		function setProfileData(data) {
			profileVm.user = data;

			if ($rootScope.currentUser.UserId == data.UserId) {
				$rootScope.currentUser = data;
			}
		}
	}
})();
