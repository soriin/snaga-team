(function() {
	'use strict';

	var app = angular.module('app', ['ngSanitize', 'ngResource', 'ui.router', 'ngCookies', 'ng-polymer-elements']);

	app.constant('VERSION', '0.1.0');

	app.config(['$stateProvider', '$locationProvider', '$urlRouterProvider',
	 function appConfig($stateProvider, $locationProvider, $urlRouterProvider) {
		$locationProvider.hashPrefix('!');
		$urlRouterProvider.otherwise("/login");

		$stateProvider.state('login', {
			url: "/login", // root route
			views: {
				"mainView": {
					templateUrl: "partials/login.html",
					controller: 'LoginCtrl as login'
				}
			}
		})
		.state('events', {
			url: "/events", // Main display for active events
			views: {
				"mainView": {
					templateUrl: "partials/events.html",
					controller: 'EventsCtrl as events'
				}
			}
		})
		.state('profile', {
			url: "/profile",
			views: {
				"mainView": {
					templateUrl: "partials/profile.html",
					controller: 'ProfileCtrl as profile'
				}
			},
			data: {
				rule: ensureLoggedIn
			},
			resolve: {
				currentData: getCurrentUserData
			}
		});

		function ensureLoggedIn(user, $window) {
			if ($window.gapi.auth2 == undefined || $window.gapi.auth2.getAuthInstance().isSignedIn.get() == false) {
				return { to: "login"};
			}
		}

		function getCurrentUserData(UserAccess) {
			return UserAccess.createUser();
		}

		return $locationProvider.html5Mode(false);
	}]);

	app.config(['$resourceProvider', function($resourceProvider) {
		// Don't strip trailing slashes from calculated URLs
		$resourceProvider.defaults.stripTrailingSlashes = false;
	}]);

	app.config([
	 "$httpProvider", function ($httpProvider) {
	     $httpProvider.interceptors.push('authHttpRequestInterceptor');
	}]);

	app.run(function($rootScope, $state, $currentUser, $window) {
	  $rootScope.$on('$stateChangeStart', function(e, to) {
	    if (to.data == undefined || !angular.isFunction(to.data.rule)) return;
	    var result = to.data.rule($currentUser, $window);

	    if (result && result.to) {
	      e.preventDefault();
	      // Optionally set option.notify to false if you don't want
	      // to retrigger another $stateChangeStart event
	      $state.go(result.to, result.params, {notify: false});
	    }
	  });
	});
})();
