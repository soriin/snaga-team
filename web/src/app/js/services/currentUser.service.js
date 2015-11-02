(function() {
	'use strict';

	angular.module('app').service('$currentUser', currentUserSvc);

	function currentUserSvc() {
		this.myUser = {};

		this.SetCurrentUser = function(user) {
			this.myUser = user;
		};

		this.GetCurrentUser = function() {
			return this.myUser;
		};
	}
})();
