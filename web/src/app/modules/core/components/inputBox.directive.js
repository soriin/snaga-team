(function() {
	'use strict';

	angular.module('app.core').directive('snagaInputBox', inputBox);

	function inputBox() {

		return {
			restrict: 'E',
			templateUrl: 'partials/core/components/input_box.html',
			controller: 'InputBoxController as ib',
			bindToController: true,
			scope: {
				label: '@',
				model: '='
			}
		};
	}
})();
