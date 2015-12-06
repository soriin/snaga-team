(function() {
	'use strict';

	angular.module('app.events').controller('EventsListCtrl', ['$state', 'EventAccess', EventsListCtrl]);

	function EventsListCtrl($state, EventAccess) {
		// this is the controller for the whole page
		var eventsVm = this;
		eventsVm.Create = create;
		eventsVm.Events = [];
		activate();

		///////////////// Functions ////////////////////////
		function activate() {
			EventAccess.getEvents().then(refreshEvents);
		}

		function create() {
			$state.go("events.new");
		}

		function refreshEvents(data) {
			if (Object.prototype.toString.call( data ) === '[object Array]') {
				eventsVm.Events = data;
			}
		}
	}
})();
