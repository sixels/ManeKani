import { Config } from "jest";
import { pathsToModuleNameMapper } from "ts-jest";
const { compilerOptions } = require("./tsconfig");

const paths_with_prefix: Record<string, string[]> = compilerOptions.paths;

export const paths = Object.entries(paths_with_prefix).reduce(
	(prev, [pattern, path]) => {
		prev[pattern] = path.map((p) => p.replace(/^src\//, ""));
		return prev;
	},
	{},
);

const config: Config = {
	moduleNameMapper: pathsToModuleNameMapper(paths, {
		prefix: "<rootDir>",
	}),
	moduleFileExtensions: ["js", "json", "ts"],
	rootDir: "src",

	testRegex: [".*\\.spec\\.ts$", ".*\\.e2e-spec\\.ts$"],
	transform: {
		"^.+\\.(t|j)s$": "ts-jest",
	},
	collectCoverageFrom: ["**/*.(t|j)s"],
	coverageDirectory: "../coverage",
	testEnvironment: "node",
};

export default config;
