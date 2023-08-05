import React from "react";
import { styled } from "styled-components";
import bg from "./img/bg.png";

const AppStyled = styled.div`
	height: 100vh;
	background-image: url(${bg});
	position:relative;
`;
function App(): React.JSX.Element {
	return (
		<AppStyled className="App">
			<main></main>
		</AppStyled>
	);
}

export default App;
