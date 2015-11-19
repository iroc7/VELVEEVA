#!/usr/bin/env node


(function init() {
	var path = require("path");
	var cliArgs = require("command-line-args");

	var Velveeva = require(path.join(__dirname, "lib/velveeva.js"));

	var configFile = require(path.join(process.cwd(),'VELVEEVA-config.json'));

	var cli = cliArgs([
	  { name: "clean", type: Boolean, alias: "c", description: "Clean up the mess (mom would be proud!) [Selected when no options are given]"},
	  { name: "dev", type: Boolean, alias: "dev", description: "Use the quick-bake test kitchen environment (no screenshots, no packaging). This is a shortcut to using velveeva --clean --watch"},
	  { name: "help", type: Boolean, alias: "h", description: "Display this message"},
	  { name: "package", type: Boolean, alias: "p", description: "Wrap it up [Selected when no options are given]"},
	  { name: "relink", type: Boolean, alias: "r", description: "Make some href saussage (replace relative links with global and convert to veeva: protocol)"},
	  { name: "screenshots", type: Boolean, alias: "s", description: "Include Screenshots [Selected when no options are given]"},
	  { name: "veev2rel", type: Boolean, alias: "2", description: "Convert veeva: hrefs to relative links"},
	  { name: "verbose", type: Boolean, alias: "v", description: "Chatty Cathy"},
	  { name: "watch", type: Boolean, alias: "w", description: "Watch for changes and re-bake on change" } 
	]);

	var options = cli.parse();
	var V = new Velveeva(configFile);

	if (options.clean) V.config.FLAGS.CLEAN = true;
	if (options.dev) {
		V.config.FLAGS.PACKAGE = false;
		V.config.FLAGS.SCREENSHOTS = false;
		V.config.FLAGS.DEV = true;
		V.config.FLAGS.WATCH = true;
	}
	if (options.package) V.config.FLAGS.PACKAGE = true;
	if (options.relink) V.config.FLAGS.RELINK = true;
	if (options.screenshots) V.config.FLAGS.SCREENSHOTS = true;
	if (options.verbose) V.config.FLAGS.VERBOSE = true;
	if (options.veev2rel) V.config.FLAGS.VEEV2REL = true;
	if (options.watch) V.config.FLAGS.WATCH = true;

	if (Object.keys(options).length === 0 || (Object.keys(options).length === 1) && options.verbose) {
	  // default case, also allows for default options with the verbose flag
	  V.config.FLAGS.PACKAGE = true;
	  V.config.FLAGS.RELINK = false;
	  V.config.FLAGS.SCREENSHOTS = true;
	  V.config.FLAGS.CLEAN = true;

	}

	if (options.help) {
		console.log(cli.getUsage());
	} else {
	    V.run();
	}
})();