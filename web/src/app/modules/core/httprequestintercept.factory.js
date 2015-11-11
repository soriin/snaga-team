(function() {
	'use strict';

	angular.module("app").factory("authHttpRequestInterceptor", ['$cookieStore', httpRequestIntercept]);

	function httpRequestIntercept ($cookieStore) {
		var svc = {
			request: request
		};

		return svc;

		///////////////////////////////////

    function request (config) {
			var token = $cookieStore.get('token');
        if (token != undefined && token.length > 0) {
          config.headers["Auth-Token"] = token;
        }
      return config;
	  }
  }
})();
