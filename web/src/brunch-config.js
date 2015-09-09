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
					'bower_components/bootstrap/dist/js/bootstrap.js'
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
			'if not exist "../public/bower_components" mkdir "../public/bower_components"',
			'cp -R  bower_components/*paper* ../public/bower_components/',
			'cp -R  bower_components/*iron* ../public/bower_components/',
			'cp -R  bower_components/*polymer* ../public/bower_components/'
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
