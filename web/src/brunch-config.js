var path = require('path');

exports.config = {
	// See docs at https://github.com/brunch/brunch/blob/master/docs/config.md
	modules: {
		definition: false,
		wrapper: false
	},

	paths: {
		"public": path.resolve('../public'),
		"watched": ['app', 'vendor']
	},

	files: {
		javascripts: {
			joinTo: {
				'js/app.js': /^app/,
				'js/vendor.js': [
					/^vendor/,

					// external libs
					'bower_components/modernizr/modernizr.js',
					'bower_components/jquery/dist/jquery.js',
					'bower_components/lodash/dist/lodash.js',
					'bower_components/bootstrap/dist/js/bootstrap.js',
					'bower_components/webcomponentsjs/webcomponents-lite.js',
					'bower_components/web-animations-js/web-animations.min.js',
					'bower_components/ng-polymer-elements/ng-polymer-elements.js',
					'bower_components/angular-cookies/angular-cookies.js',

					// angular
					'bower_components/angular/angular.js',
					'bower_components/angular-resource/angular-resource.js',
					'bower_components/angular-sanitize/angular-sanitize.js',
					'bower_components/angular-ui-router/release/angular-ui-router.js',
					'bower_components/ocModal/dist/ocModal.min.js'
				],
				'test/scenarios.js': /^test(\/|\\)e2e/
			},
			order: {
				before: [
					// jquery
					'bower_components/jquery/dist/jquery.js',

					// angular
					'bower_components/angular/angular.js',

					// bootstrap
					'bower_components/bootstrap/dist/js/bootstrap.js',

					// module ordering
					'app/modules/core/core.module.js',
					'app/modules/events/events.module.js',
					'app/modules/layout/layout.module.js',
					'app/modules/profile/profile.module.js',
					'app/app.module.js',
				]
			}
		},
		stylesheets: {
			joinTo: {
				'css/app.css': [/^app/, 'bower_components/bootstrap/scss/bootstrap.scss']
			}
		}
	},

	plugins: {
		ng_annotate: {
			pattern: /^app/
		},
		traceur: {
			paths: /^app/,
			options: {
				experimental: true
			}
		},
		autoprefixer: {
			browsers: [
				"last 2 version",
				"> 1%", // browsers with > 1% usage
				"ie >= 9"
			],
			cascade: false
		},
		afterBrunch: [
			"call postbuild.bat",
			"./postbuild.sh"
		]
	},

	server: {
		port: 3333
	},

	conventions: {
		assets: /app(\\|\/)assets(\\|\/)/
	},

	sourceMaps: true
};
