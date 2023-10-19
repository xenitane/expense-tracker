import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import { server } from "./.env.json";

// https://vitejs.dev/config/
export default defineConfig({
	server,
	plugins: [react()],
	root: process.cwd(),
});
