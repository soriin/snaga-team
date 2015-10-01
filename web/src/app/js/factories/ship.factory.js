angular.module('app').factory('Ships', ['$resource', function ($resource) {
	return $resource('/api/ships/:id', null,
		{

		}
	);
}]);
