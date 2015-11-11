(function() {
	'use strict';

	angular.module('app').config(['$stateProvider', '$locationProvider', '$urlRouterProvider',
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
})();
