/* eslint-disable no-mixed-spaces-and-tabs */
import { styled } from "styled-components";
import { InnerLayout } from "../../styles/Layouts";
import Chart from "../Chart/Chart";
import { useGlobalContext } from "../../context/GlobalContext";
import { useEffect } from "react";
import { dollar } from "../../utils/icons";
import History from "../History/History";

const DashboardStyled = styled.div`
	.stats-container {
		display: grid;
		grid-template-columns: repeat(5, 1fr);
		gap: 2rem;
		.chart-container {
			grid-column: 1 / 4;
			height: 400px;
			.amount-container {
				display: grid;
				grid-template-columns: repeat(4, 1fr);
				gap: 2rem;
				margin-top: 2rem;
				.income,
				.expense {
					grid-column: span 2;
				}
				.income,
				.expense,
				.balance {
					background: #fcf6f9;
					border: 2px solid #fff;
					box-shadow: 0px 1px 15px #0000000f;
					border-radius: 20px;
					padding: 1rem;
					p {
						font-weight: 700;
						font-size: 3.5rem;
					}
				}
				.balance {
					grid-column: 2/4;
					display: flex;
					flex-direction: column;
					justify-content: center;
					align-items: center;
					p {
						color: var(--color-green);
						opacity: 0.6;
						font-size: 4.5rem;
					}
				}
			}
		}
		.history-container {
			grid-column: 4/-1;
			h2 {
				margin: 1rem 0;
				display: flex;
				align-items: center;
				justify-content: space-between;
			}
			.income-title,
			.expense-title {
				font-size: 1.2rem;
				span {
					font-size: 1.8rem;
				}
			}
			.income-item,
			.expense-item {
				background: #fcf6f9;
				border: 2px solid #fff;
				box-shadow: 0px 1px 15px #0000000f;
				padding: 1rem;
				border-radius: 20px;
				display: flex;
				justify-content: space-between;
				align-items: center;
				p {
					font-weight: 600;
					font-size: 1.6rem;
				}
			}
		}
	}
`;

function Dashboard() {
	const {
		getIncomes,
		getExpenses,
		totalIncome,
		totalExpense,
		incomes,
		expenses,
	} = useGlobalContext();
	useEffect(() => {
		getIncomes();
		getExpenses();
	}, []);
	return (
		<DashboardStyled>
			<InnerLayout>
				<h1>All Transactions</h1>

				<div className="stats-container">
					<div className="chart-container">
						<Chart />

						<div className="amount-container">
							<div className="income">
								<h2>Total Income</h2>

								<p>
									{dollar}

									{totalIncome()}
								</p>
							</div>

							<div className="expense">
								<h2>Total Expense</h2>

								<p>
									{dollar}

									{totalExpense()}
								</p>
							</div>

							<div className="balance">
								<h2>Total Balance</h2>

								<p>
									{dollar}

									{totalIncome() - totalExpense()}
								</p>
							</div>
						</div>
					</div>

					<div className="history-container">
						<History />
						<h2 className="income-title">
							Min <span>Income</span> Max
						</h2>
						<div className="income-item">
							<p>
								{dollar}
								{incomes.length
									? Math.min(
										...incomes.map(
											(income) => income.amount
										)
									  )
									: 0}
							</p>
							<p>
								{dollar}
								{incomes.length
									? Math.max(
										...incomes.map(
											(income) => income.amount
										)
									  )
									: 0}
							</p>
						</div>
						<h2 className="expense-title">
							Min <span>Expense</span> Max
						</h2>
						<div className="expense-item">
							<p>
								{dollar}
								{expenses.length
									? Math.min(
										...expenses.map(
											(expense) => expense.amount
										)
									  )
									: 0}
							</p>
							<p>
								{dollar}
								{expenses.length
									? Math.max(
										...expenses.map(
											(expense) => expense.amount
										)
									  )
									: 0}
							</p>
						</div>
					</div>
				</div>
			</InnerLayout>
		</DashboardStyled>
	);
}

export default Dashboard;
