import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
	ArcElement,
} from "chart.js";

import { Line } from "react-chartjs-2";
import { styled } from "styled-components";
import { useGlobalContext } from "../../context/GlobalContext";
import moment from "moment";

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

	const process_data = () => {
		const data = {};
		incomes.forEach((income) => {
			data[income?.date] = { income: 0, expense: 0 };
		});
		expenses.forEach((expense) => {
			data[expense?.date] = { income: 0, expense: 0 };
		});
		incomes.forEach((income) => {
			data[income.date].income += income.amount;
		});
		expenses.forEach((expense) => {
			data[expense.date].expense += expense.amount;
		});
		const dates = Object.keys(data).sort((a, b) => {
			const aDate = moment(a, "dd-MM-yyyy").toDate().getTime();
			const bDate = moment(b, "dd-MM-yyyy").toDate().getTime();
			return aDate - bDate;
		});
		const res = {};
		res.labels = dates;
		res.datasets = [
			{
				label: "Income",
				data: [],
				backgroundColor: "green",
				tension: 0.2,
			},
			{
				label: "Expense",
				data: [],
				backgroundColor: "red",
				tension: 0.2,
			},
		];
		dates.forEach((date) => {
			res.datasets[0].data.push(data[date].income);
			res.datasets[1].data.push(data[date].expense);
		});
		return res;
	};

	return (
		<ChartStyled>
			<Line data={process_data()} />
		</ChartStyled>
	);
}

export default Chart;
