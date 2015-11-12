(function() {
	'use strict';

	angular.module('app.profile').directive('snagaAdminSwitch', ['UserAccess', '$currentUser', adminSwitch]);

	function adminSwitch(Users, $currentUser) {

		return {
			restrict: 'E',
			templateUrl: 'partials/profile/admin_switch.html',
			link: link,
			controller: 'AdminController as admin',
			scope: {
				profileUser: '='
			}
		};

		function link($scope, $element, $attrs) {
			$scope.currentUser = $currentUser.GetCurrentUser();
		}
	}
})();
