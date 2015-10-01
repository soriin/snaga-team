angular.module('app', ['ngSanitize', 'ngResource', 'ui.router', 'ngCookies', 'ng-polymer-elements'])
	.constant('VERSION', '0.1.0')
	.config(function appConfig($stateProvider, $locationProvider, $urlRouterProvider) {
		$locationProvider.hashPrefix('!');
		$urlRouterProvider.otherwise("/login");

		$stateProvider.state('login', {
			url: "/login", // root route
			views: {
				"mainView": {
					templateUrl: "partials/login.html",
					controller: 'LoginCtrl'
				}
			}
		})
		.state('profile', {
			url: "/profile",
			views: {
				"mainView": {
					templateUrl: "partials/profile.html",
					controller: 'ProfileCtrl'
				}
			}
		});

		return $locationProvider.html5Mode(false);
	})
	.config(['$resourceProvider', function($resourceProvider) {
	 // Don't strip trailing slashes from calculated URLs
	 $resourceProvider.defaults.stripTrailingSlashes = false;
 }])
 .config([
	 "$httpProvider", function ($httpProvider) {
	     $httpProvider.interceptors.push('authHttpRequestInterceptor');
 }]);
