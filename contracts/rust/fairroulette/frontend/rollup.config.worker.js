import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import typescript from '@rollup/plugin-typescript';
import { terser } from 'rollup-plugin-terser';

const production = !process.env.ROLLUP_WATCH;
console.log("prod: " + production);
function serve() {
	let server;

	function toExit() {
		if (server) server.kill(0);
	}

	return {
		writeBundle() {
			if (server) return;
			server = require('child_process').spawn('npm', ['run', 'start', '--', '--dev'], {
				stdio: ['ignore', 'inherit', 'inherit'],
				shell: true
			});

			process.on('SIGTERM', toExit);
			process.on('exit', toExit);
		}
	};
}

export default {
	input: 'src/wasp_client/web_worker/pow.worker.ts',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'app',
		file: 'public/build/pow.worker.js'
	},
	plugins: [
		typescript({
			sourceMap: !production,
			inlineSources: !production,
		}),
		resolve({
			browser: true,
			preferBuiltins: true,
		}),
		commonjs(),

		!production && serve(),
		production && terser(),
	],
	watch: {
		clearScreen: false
	}
};
