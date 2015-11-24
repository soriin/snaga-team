(function() {
	'use strict';

	angular.module('app.events').controller('EventsNewCtrl', ['$state', EventsNewCtrl]);

	function EventsNewCtrl($state) {
		// this is the controller for the whole page
		var eventsVm = this;
		eventsVm.Create = create;
		eventsVm.Title = "Test Title";

		///////////////// Functions ////////////////////////

		function create() {
			$state.go("events");
		}
	}
})();
