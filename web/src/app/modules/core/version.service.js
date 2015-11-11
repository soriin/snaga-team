(function() {
	'use strict';

	angular.module('app').factory('$version', version);

	function version(VERSION) {
		return VERSION;
	}
})();
