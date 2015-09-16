angular.module('app').controller('LoginCtrl', ['$scope', function LoginCtrl($scope) {
	var myScope = $scope;

	myScope.login = login;

	function login() {
		alert("Logged in!");
	};
}]);
