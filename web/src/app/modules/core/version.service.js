(function() {
	'use strict';

	angular.module('app.core').factory('$version', version);

	function version(VERSION) {
		return VERSION;
	}
})();
