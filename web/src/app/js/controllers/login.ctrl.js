angular.module('app').controller('LoginCtrl', ['$scope', '$window', '$location', '$state',
 function LoginCtrl($scope, $window, $location, $state) {
	$scope.login = login;
	$scope.logout = logout;

	$window.gapi.signin2.render('googleSigninBtn', {
                'scope': 'https://www.googleapis.com/auth/plus.login',
                'width': 200,
                'height': 50,
                'longtitle': true,
                'theme': 'dark',
                'onsuccess': $scope.login
            });

	function login(googleUser) {
		var profile = googleUser.getBasicProfile();
		console.log("Name: " + profile.getName());
		console.log("Image URL: " + profile.getImageUrl());
		console.log("Email: " + profile.getEmail());

		// The ID token you need to pass to your backend:
		var id_token = googleUser.getAuthResponse().id_token;
		console.log("ID Token: " + id_token);

		$window.gapi.auth2.getAuthInstance().then(function () { $state.go("profile"); })

	};

	function logout() {
		var auth2 = $window.gapi.auth2.getAuthInstance();
    auth2.signOut().then(function () {
      console.log('User signed out.');
    });
	}
}]);
