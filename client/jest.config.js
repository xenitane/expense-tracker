export default {
	roots: ["<rootDir>/src"],
	collectCoverageFrom: [
		"src/**/*.{js,jsx,ts,tsx}",
		"!src/**/*.d.ts",
		"!src/mocks/**",
	],
	coveragePathIgnorePatterns: [],
	setupFilesAfterEnv: ["./config/jest/setupTests.js"],
	testEnvironment: "jsdom",
	modulePaths: ["<rootDir>/src"],
	transform: {
		"^.+\\.(ts|js|tsx|jsx)$": "@swc/jest",
		"^.+\\.css$": "<rootDir>/config/jest/cssTransform.cjs",
		"^.+\\.scss$": "<rootDir>/config/jest/scssTransform.cjs",
		"^(?!.*\\.(js|jsx|mjs|cjs|ts|tsx|css|scss|json)$)":
			"<rootDir>/config/jest/fileTransform.cjs",
	},
	transformIgnorePatterns: [
		"[/\\\\]node_modules[/\\\\].+\\.(js|jsx|mjs|cjs|ts|tsx)$",
		"^.+\\.module\\.(css|sass|scss)$",
	],
	moduleNameMapper: {
		"^react-native$": "react-native-web",
		"^.+\\.module\\.(css|sass|scss)$": "identity-obj-proxy",
	},
	moduleFileExtensions: [
		"tsx",
		"ts",
		"web.js",
		"js",
		"web.ts",
		"web.tsx",
		"json",
		"web.jsx",
		"jsx",
		"node",
	],
	watchPlugins: [
		"jest-watch-typeahead/filename",
		"jest-watch-typeahead/testname",
	],
	resetMocks: true,
};
