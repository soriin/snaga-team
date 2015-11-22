(function() {
	'use strict';

	angular.module('app.layout').controller('AppCtrl', ['$rootScope', AppCtrl]);

	function AppCtrl($rootScope) {
		// this is the controller for the whole page
		var appVm = this;
		$rootScope.isLoggedIn = false;
	}
})();
