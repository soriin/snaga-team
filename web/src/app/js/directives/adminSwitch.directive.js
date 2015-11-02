angular.module('app').directive('snagaAdminSwitch', function adminSwitch() {
	return {
		restrict: 'E',
		template: '<div>',
		link: link
	};

	function link($scope, $element, $attrs) {
		// you can do things here if you want!
	}
});
