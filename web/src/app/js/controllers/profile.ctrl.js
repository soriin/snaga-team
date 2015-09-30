angular.module('app').controller('ProfileCtrl', ['$scope', '$window', '$state', 'Users',
	function ProfileCtrl($scope, $window, $state, Users) {
		if ($window.gapi.auth2 == undefined || $window.gapi.auth2.getAuthInstance().isSignedIn.get() == false) {
			$state.go("login");
			return;
		}

		$scope.someText = "HI, " + $window.gapi.auth2.getAuthInstance().currentUser.get().getBasicProfile().getName();
		//var me = Users.get({id: })
}]);
