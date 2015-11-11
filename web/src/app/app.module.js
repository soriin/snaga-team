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
