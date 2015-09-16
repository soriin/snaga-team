// Declare app level module which depends on filters, and services
angular.module('app', ['ngSanitize', 'ngResource', 'ui.router'])
	.constant('VERSION', '0.1.0')
	.config(function appConfig($stateProvider, $locationProvider, $urlRouterProvider) {
		// $locationProvider.hashPrefix('!');
		$urlRouterProvider.otherwise("/login");

		$stateProvider.state('login', {
			url: "/login", // root route
			views: {
				"mainView": {
					templateUrl: "partials/login.html",
					controller: 'LoginCtrl'
				}
			}
		});

		// /!\ Without server side support html5 must be disabled.
		return $locationProvider.html5Mode(false);
	});
