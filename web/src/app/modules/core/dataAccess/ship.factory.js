(function() {
	'use strict';

	angular.module('app.core').factory('Ships', ['$resource', shipFactory]);

	function shipFactory ($resource) {
		return $resource('/api/ships/:id', null,
			{

			}
		);
	}
})();
