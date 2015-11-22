(function() {
	'use strict';

	angular.module('app.core').controller('LoginCtrl',
	 ['$window', '$location', '$state', '$cookieStore', '$rootScope', LoginCtrl]);

  function LoginCtrl($window, $location, $state, $cookieStore, $rootScope) {
		var loginVm = this;
  	loginVm.login = login;
  	loginVm.logout = logout;
		loginVm.renderLogin = renderLogin;
		loginVm.notRendered = true;

		$rootScope.$watch('isLoggedIn', function(val) { loginVm.isLoggedIn = val;});

		///////////////// Functions ////////////////////////

		function renderLogin() {
			console.log("rendering login button");
	  	$window.gapi.signin2.render('googleSigninBtn', {
	      'scope': 'profile',
	      'width': 100,
	      'height': 40,
	      'longtitle': false,
	      'theme': 'dark',
	      'onsuccess': login,
	      'onfailure': loginFail
	    });

			loginVm.notRendered = false;
		}

  	function loginFail(e) {
  		console.log(e);
  	}

  	function login(googleUser) {
  		var profile = googleUser.getBasicProfile();
  		console.log("Name: " + profile.getName());
  		console.log("Image URL: " + profile.getImageUrl());
  		console.log("Email: " + profile.getEmail());

  		// The ID token you need to pass to your backend:
  		var id_token = googleUser.getAuthResponse().id_token;
  		console.log("ID Token: " + id_token);
      $cookieStore.put('token', id_token)

  		$window.gapi.auth2.getAuthInstance().then(function () {
				$rootScope.isLoggedIn = true;
				$rootScope.$apply();
				$state.go("events");
			})
  	}

  	function logout() {
  		var auth2 = $window.gapi.auth2.getAuthInstance();
      auth2.signOut().then(function () {
				$rootScope.isLoggedIn = false;
				$rootScope.$apply();
        console.log('User signed out.');
        $state.go("events");
      });
  	}
  }
})();
