(function() {
	'use strict';

	angular.module('app.core').factory('UserAccess', ['$http', UserAccess]);

	function UserAccess($http, logger) {
		var svc = {
			createUser: createUser,
			updateUser: updateUser,
			addGroup: addGroup,
			removeGroup: removeGroup
		};
		return svc;

		////////////////////////////////////////
		function createUser() {
			return sendReq({
				method: "POST",
				url: "/api/users/"
			});
		}

		function updateUser(id, body) {
			return sendReq({
				method: "PUT",
				url: "/api/users/" + id,
				data: body
			});
		}

		function addGroup(id, groupName) {
			return sendReq({
				method: "PUT",
				url: "/api/users/" + id + "/groups",
				data: {
					Action: "add",
					GroupName: groupName
				}
			});
		}

		function removeGroup(id, groupName) {
			return sendReq({
				method: "PUT",
				url: "/api/users/" + id + "/groups",
				data: {
					Action: "remove",
					GroupName: groupName
				}
			});
		}

		function sendReq(req) {
			return $http(req).then(onSuccess, onError);
		}

		function onSuccess(response) {
			return response.data;
		}

		function onError(error) {
			console.log(error);
		}
	}

})();
