(function() {
	'use strict';

	angular.module('app.core').factory('EventAccess', ['Events', EventAccess]);

	function EventAccess(Events, logger) {
		var svc = {
			createEvent: createEvent,
			updateEvent: updateEvent
		};
		return svc;

		////////////////////////////////////////
		function createEvent() {
			return Events.save({}).$promise.then(onSuccess).catch(onError);
		}

		function updateEvent(id, data) {
			return Events.update({id: id}, data).$promise.then(onSuccess).catch(onError);
		}

		function onSuccess(data) {
			return data;
		}

		function onError(error) {
			//logger.error(error);
		}
	}

	})();
