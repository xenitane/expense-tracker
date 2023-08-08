import { styled, keyframes } from "styled-components";
import useWindowSize from "../../utils/useWindowSize";

const moveOrb = (data) => keyframes`
	0%{
		transform: translate(0,0);
	}
	50%{
		transform: translate(${data.width}px,${data.height / 2}px);
	}
	100%{
		transform :translate(0,0);
	}`;
const OrbStyled = (data) => styled.div`
	width: 70vh;
	height: 70vh;
	position: absolute;
	border-radius: 50%;
	margin-top: -37vh;
	margin-left: -37vh;
	background: linear-gradient(180deg, #f56692 0%, #f2994a 100%);
	filter: blur(400px);
	animation: ${moveOrb(data)} 15s alternate linear infinite;
`;

function Orb() {
	const data = useWindowSize();
	const OrbTag = OrbStyled(data);
	return <OrbTag />;
}

export default Orb;
