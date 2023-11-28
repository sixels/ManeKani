import { resolve } from "path";
import { unstable_vitePlugin as remix } from "@remix-run/dev";
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin";
import { defineConfig } from "vite";

import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
	server: {
		host: true,
		port: 11011,
	},

	// ssr: {
	//   external: ['nock', 'aws-sdk', 'mock-aws-s3'],
	// },
	// build: {
	//   rollupOptions: {
	//     external: ['nock', 'aws-sdk', 'mock-aws-s3'],
	//   },
	// },
	resolve: {
		preserveSymlinks: true,
		alias: {
			nock: resolve(__dirname, "app/empty.ts"),
			"mock-aws-s3": resolve(__dirname, "app/empty.ts"),
			"aws-sdk": resolve(__dirname, "app/empty.ts"),
		},
	},
	plugins: [
		remix({
			ignoredRouteFiles: ["**/.*"],
			browserNodeBuiltinsPolyfill: {
				modules: {
					assert: true,
					crypto: true,
					path: true,
					util: true,
					fs: true,
					constants: true,
					buffer: true,
					punycode: true,
					querystring: true,
					timers: true,
					https: true,
					url: true,
					stream: true,
					events: true,
					http: true,
					child_process: true,
					zlib: true,
					os: true,
				},
			},
		}),
		tsconfigPaths(),
		vanillaExtractPlugin(),
	],
});
