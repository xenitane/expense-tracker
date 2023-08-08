import { styled } from "styled-components";
import bg from "./img/bg.png";
import { useMemo, useState } from "react";
import { MainLayout } from "./styles/Layouts";
import Navigation from "./Components/Navigation/Navigation";
import Orb from "./Components/Orb/Orb";
import Dashboard from "./Components/Dashboard/Dashboard";
import Incomes from "./Components/Incomes/Incomes";
import Expenses from "./Components/Expenses/Expenses";

const AppStyled = styled.div`
	background: url(${(props) => props.theme.bg});
	height: 100vh;
	position: relative;
	main {
		flex: 1;
		background: rgba(252, 246, 249, 0.78);
		border: 3px solid #fff;
		backdrop-filter: blur(4.5px);
		border-radius: 32px;
		overflow-x: hidden;
		&::-webkit-scrollbar {
			width: 0;
		}
	}
`;

function App() {
	const [active, setActive] = useState(0);

	const orbMemo = useMemo(() => {
		return <Orb />;
	}, []);

	const displayData = () => {
		switch (active) {
		case 0:
			return <Dashboard />;
		case 1:
			return <Dashboard />;
		case 2:
			return <Incomes />;
		case 3:
			return <Expenses />;
		default:
			return <Dashboard />;
		}
	};

	return (
		<AppStyled
			theme={{ bg: bg }}
			className="App"
		>
			{orbMemo}

			<MainLayout>
				<Navigation
					active={active}
					setActive={setActive}
				/>

				<main>
					{displayData()}
				</main>
			</MainLayout>
		</AppStyled>
	);
}

export default App;
