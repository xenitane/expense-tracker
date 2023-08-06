import { JSX, MouseEventHandler } from "react";
import { styled } from "styled-components";

const ButtonStyled = styled.button`
	outline: none;
	border: none;
	font-family: inherit;
	font-size: inherit;
	display: flex;
	align-items: center;
	gap: 0.5rem;
	cursor: pointer;
	transition: all 0.4s ease-in-out;
`;

function Button({ name, icon, bg, bPad, color, bRad, onClick }) {
	return (
		<ButtonStyled
			style={{
				background: bg,
				padding: bPad,
				borderRadius: bRad,
				color: color
			}}
			onClick={onClick}
		>
			{icon}
			{name}
		</ButtonStyled>
	);
}

export default Button;
