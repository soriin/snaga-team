(function() {
	'use strict';

	angular.module('app').factory('Events', ['$resource', eventFactory]);

	function userFactory ($resource) {
		return $resource('/api/events/:id', null,
			{
				'update' : {method: 'PUT'}
			}
		);
	}
})();
