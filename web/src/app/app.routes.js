(function() {
	'use strict';

	angular.module('app').config(['$stateProvider', '$locationProvider', '$urlRouterProvider',
	 function appConfig($stateProvider, $locationProvider, $urlRouterProvider) {
		$locationProvider.hashPrefix('!');
		$urlRouterProvider.otherwise("/events");

		var getProfileRes = ['UserAccess', '$stateParams', '$q', 'currentUser', getProfileData];
		var getCurrentUserRes = ['UserAccess', '$window', '$rootScope', '$q', '$state', fetchCurrentUser]

		$stateProvider.state('wrapper', {
			templateUrl: "partials/layout/default.layout.html"
		})
		.state('login', {
			url: "/login",
			views: {
				"mainView": {
					templateUrl: "partials/core/login.html",
					controller: 'LoginCtrl as login'
				}
			}
		})
		.state('events', {
			parent: 'wrapper',
			url: "/events", // Main display for active events
			views: {
				"mainView": {
					templateUrl: "partials/events/events.list.html",
					controller: 'EventsListCtrl as events'
				}
			}
		})
		.state('events.new', {
			parent: 'wrapper',
			url: "/events/new", // Main display for active events
			views: {
				"mainView": {
					templateUrl: "partials/events/events.new.html",
					controller: 'EventsNewCtrl as event'
				}
			}
		})
		.state('profile', {
			parent: 'wrapper',
			url: "/profile",
			views: {
				"mainView": {
					templateUrl: "partials/profile/profile.html",
					controller: 'ProfileCtrl as profile'
				}
			},
			resolve: {
				currentUser: getCurrentUserRes,
				profileData: getProfileRes
			}
		})
		.state('profile-other', {
			url: "/profile/:userId",
			views: {
				"mainView": {
					templateUrl: "partials/profile/profile.html",
					controller: 'ProfileCtrl as profile'
				}
			},
			resolve: {
				currentUser: getCurrentUserRes,
				profileData: getProfileRes
			}
		});

		function fetchCurrentUser(UserAccess, $window, $rootScope, $q, $state) {
			console.log("entering fetchCurrentUser");
			if (!!$rootScope.currentUser) {
				return $rootScope.currentUser;
			}
			var defer = $q.defer();
			if (!!$window.gapi.auth2 && $window.gapi.auth2.getAuthInstance().isSignedIn.get() == true) {
				console.log("fetching current user");
				UserAccess.createUser().then(function(data) {
					if (!!data) {
						console.log("setting current user");
						$rootScope.currentUser = data;
						$rootScope.isLoggedIn = true;
						defer.resolve(data);
					}
					else {
						console.error("error fetching current user");
						defer.reject();
						$state.go("events");
					}
				});
			}
			else {
				defer.reject();
				$state.go("login");
			}

			return defer.promise;

			console.log("not fetching current user");
		}

		function getProfileData(UserAccess, $stateParams, $q, currentUser) {
			if ($stateParams.userId === undefined) {
				console.log("getProfileData returning current user");
				return currentUser;
			}
			console.log("getProfileData fetching userId "+ $stateParams.userId);
			return UserAccess.getUser($stateParams.userId);
		}

		return $locationProvider.html5Mode(false);
	}]);
})();
