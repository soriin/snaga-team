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
			}
		});

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
})();
