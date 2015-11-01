'use strict';
describe("login controller", function() {
	beforeEach(module("app"));

	describe("LoginCtrl", function() {
		return it("should make scope testable", inject(function($rootScope, $controller) {
			var view = $controller("ViewCtrl");
			return expect(view.content).toEqual("This is the partial for view.");
		}));
	});
});
