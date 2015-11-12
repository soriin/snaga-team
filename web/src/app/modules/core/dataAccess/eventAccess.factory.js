(function() {
	'use strict';

	angular.module('app.core').factory('EventAccess', ['$http', EventAccess]);

	function EventAccess($http, logger) {
		var svc = {
			createEvent: createEvent,
			updateEvent: updateEvent
		};
		return svc;

		////////////////////////////////////////
		function createEvent(body) {
			return sendReq({
				method: "POST",
				url: "/api/events/",
				data: body
			});
		}

		function updateEvent(id, body) {
			return sendReq({
				method: "PUT",
				url: "/api/events/" + id,
				data: body
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
