import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.jsx";
import GlobalProvider from "./context/GlobalProvider";
import GlobalLayout from "./styles/GlobalLayout";

ReactDOM.createRoot(document.getElementById("root")).render(
	<React.StrictMode>
		<GlobalLayout />

		<GlobalProvider>
			<App />
		</GlobalProvider>
	</React.StrictMode>
);
