angular.module("app").factory("authHttpRequestInterceptor", ['$cookieStore',
	function ($cookieStore) {
     return {
       request: function (config) {
				 var token = $cookieStore.get('token');
         if (token != undefined && token.length > 0) {
           config.headers["Auth-Token"] = token;
         }
         return config;
			 }
		 };
}]);
