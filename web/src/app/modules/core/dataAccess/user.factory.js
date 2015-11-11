(function() {
	'use strict';

	angular.module('app.core').factory('Users', ['$resource', userFactory]);

	function userFactory ($resource) {
		return $resource('/api/users/:id', null,
			{
				'update' : {method: 'PUT'}
			}
		);
	}
})();
