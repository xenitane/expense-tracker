import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	root: process.cwd(),
	server: {
		port: 3000,
		strictPort: true,
		cors: true,
		proxy: {
			"/api": {
				target: "http://localhost:5000",
				changeOrigin: true
			}
		}
	}
});
