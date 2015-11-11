(function() {
	'use strict';

	angular.module('app').directive('appVersion', appVersion);

	function appVersion() {
		return {
			restrict: 'E',
			template: '<span>v{{ "%VERSION%" | interpolate  }}</span>',
			link: link
		};

		function link($scope, $element, $attrs) {
			// you can do things here if you want!
		}
	}
})();
