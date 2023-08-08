import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import * as config from "./.env.json";

// https://vitejs.dev/config/
export default defineConfig({
	...config,
	plugins: [react()],
	root: process.cwd()
});
