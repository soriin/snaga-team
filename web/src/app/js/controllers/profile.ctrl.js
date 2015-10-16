angular.module('app').controller('ProfileCtrl', ['$scope', '$window', '$state', 'Users', 'Ships',
	function ProfileCtrl($scope, $window, $state, Users, Ships) {
		if ($window.gapi.auth2 == undefined || $window.gapi.auth2.getAuthInstance().isSignedIn.get() == false) {
			$state.go("login");
			return;
		}

		$scope.someText = "HI, " + $window.gapi.auth2.getAuthInstance().currentUser.get().getBasicProfile().getName();
		$scope.update = update;

		var me = Users.save({}, function() {
			$scope.FirstName = me.FirstName;
			$scope.LastName = me.LastName;
			$scope.DisplayName = me.DisplayName;
			$scope.InGameName = me.InGameName;
			$scope.Email = me.Email;
			$scope.Id = me.Id;
		});

		function update() {
			Users.update({id: $scope.Id},
			{
				FirstName : $scope.FirstName,
				LastName : $scope.LastName,
				DisplayName : $scope.DisplayName,
				InGameName : $scope.InGameName,
				Email : $scope.Email
			})
		};
}]);
