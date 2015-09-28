angular.module('app').controller('LoginCtrl', ['$scope', '$window', '$location', '$state', '$cookieStore',
 function LoginCtrl($scope, $window, $location, $state, $cookieStore) {
	$scope.login = login;
	$scope.logout = logout;

	$window.gapi.signin2.render('googleSigninBtn', {
		'scope': 'profile',
    'width': 100,
    'height': 40,
    'longtitle': false,
    'theme': 'dark',
    'onsuccess': $scope.login,
		'onfailure': $scope.loginFail
  });

	function loginFail(e) {
		console.log(e);
	};

	function login(googleUser) {
		var profile = googleUser.getBasicProfile();
		console.log("Name: " + profile.getName());
		console.log("Image URL: " + profile.getImageUrl());
		console.log("Email: " + profile.getEmail());

		// The ID token you need to pass to your backend:
		var id_token = googleUser.getAuthResponse().id_token;
		console.log("ID Token: " + id_token);
    $cookieStore.put('token', id_token)

		$window.gapi.auth2.getAuthInstance().then(function () { $state.go("profile"); })
		$scope.$digest();
	};

	function logout() {
		var auth2 = $window.gapi.auth2.getAuthInstance();
    auth2.signOut().then(function () {
      console.log('User signed out.');
			$state.go("login");
    });
	}
}]);
