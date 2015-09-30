angular.module('app').factory('Users', ['$resourceProvider', function ($resource) {
	return $resource('/api/users/:id', null,
		{
			
		}
	);
}]);
