import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
	ArcElement
} from "chart.js";

import { Line } from "react-chartjs-2";
import { styled } from "styled-components";
import { useGlobalContext } from "../../context/GlobalContext";

ChartJS.register(
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
	ArcElement
);

const ChartStyled = styled.div`
	background: #fcf6f9;
	border: 2px solid #fff;
	box-shadow: 0px 1px 15px #0000000f;
	padding: 1rem;
	border-radius: 20px;
	height: 100%;
`;

function Chart() {
	const { incomes, expenses } = useGlobalContext();

	const data = {
		labels: incomes.map((income) => {
			return income.date;
		}),
		datasets: [
			{
				label: "Income",
				data: incomes.map((income) => income.amount),
				backgroundColor: "green",
				tension: 0.2
			},
			{
				label: "Expenses",
				data: expenses.map((expense) => expense.amount),
				backgroundColor: "red",
				tension: 0.2
			}
		]
	};

	return (
		<ChartStyled>
			<Line data={data} />
		</ChartStyled>
	);
}

export default Chart;
