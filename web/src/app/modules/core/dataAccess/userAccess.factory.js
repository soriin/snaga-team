(function() {
	'use strict';

	angular.module('app.core').factory('UserAccess', ['Users', UserAccess]);

	function UserAccess(Users, logger) {
		var svc = {
			createUser: createUser,
			updateUser: updateUser
		};
		return svc;

		////////////////////////////////////////
		function createUser() {
			return Users.save({}).$promise.then(onSuccess).catch(onError);
		}

		function updateUser(id, data) {
			return Users.update({id: id}, data).$promise.then(onSuccess).catch(onError);
		}

		function onSuccess(data) {
			return data;
		}

		function onError(error) {
			//logger.error(error);
		}
	}

	})();
