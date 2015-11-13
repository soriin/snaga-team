(function() {
	'use strict';

	var app = angular.module('app', ['app.core', 'app.events', 'app.layout', 'app.profile']);

	app.constant('VERSION', '0.1.0');

	app.config(['$resourceProvider', function($resourceProvider) {
		// Don't strip trailing slashes from calculated URLs
		$resourceProvider.defaults.stripTrailingSlashes = false;
	}]);

	app.config([
	 "$httpProvider", function ($httpProvider) {
	     $httpProvider.interceptors.push('authHttpRequestInterceptor');
	}]);

	app.run(function($rootScope) {
	  $rootScope.$on('$stateChangeStart', function(e, to) {
			console.log("routing to:");
			console.log(to);
	  });
	});

	app.run(function($rootScope) {
	  $rootScope.$on('$stateChangeSuccess', function(e, to) {
			console.log("routed to:");
			console.log(to);
	  });
	});

	app.run(function($rootScope) {
	  $rootScope.$on('$stateChangeError', function(e, toState, toParams, fromState, fromParams, error) {
			console.log("error routing to:");
			console.log(toState);
			console.error(error);
	  });
	});
})();
