angular.module('app').service('$currentUser', function() {
	this.myUser = {};

	this.SetCurrentUser = function(user) {
		this.myUser = user;
	};

	this.GetCurrentUser = function() {
		return this.myUser;
	};
});
