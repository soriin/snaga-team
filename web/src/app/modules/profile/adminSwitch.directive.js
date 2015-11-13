(function() {
	'use strict';

	angular.module('app.profile').directive('snagaAdminSwitch', adminSwitch);

	function adminSwitch() {

		return {
			restrict: 'E',
			templateUrl: 'partials/profile/admin_switch.html',
			controller: 'AdminController as admin',
			scope: {
				profileUser: '='
			}
		};
	}
})();
