export default {
	roots: ["<rootDir>/src"],
	collectCoverageFrom: [
		"src/**/*.{js,jsx,ts,tsx}",
		"!src/**/*.d.ts",
		"!src/mocks/**"
	],
	coveragePathIgnorePatterns: [],
	setupFilesAfterEnv: ["./config/jest/setupTests.js"],
	testEnvironment: "jsdom",
	modulePaths: ["<rootDir>/src"],
	transform: {
		"^.+\\.(js|jsx|ts|tsx)$": "@swc/jest",
		"^.+\\.css$": "<rootDir>/config/jest/cssTransform.cjs",
		"^(?!.*\\.(js|jsx|mjs|cjs|ts|tsx|css|json)$)":
			"<rootDir>/config/jest/fileTransform.cjs"
	},
	transformIgnorePatterns: [
		"[/\\\\]node_modules[/\\\\].+\\.(js|jsx|mjs|cjs|ts|tsx)$",
		"^.+\\.module\\.(css|sass|scss)$"
	],
	moduleNameMapper: {
		"^react-native$": "react-native-web",
		"^.+\\.module\\.(css|sass|scss)$": "identity-obj-proxy"
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
		"node"
	],
	watchPlugins: [
		"jest-watch-typeahead/filename",
		"jest-watch-typeahead/testname"
	],
	resetMocks: true
};
