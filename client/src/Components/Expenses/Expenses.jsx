import { styled } from "styled-components";
import { InnerLayout } from "../../styles/Layouts";
import { useGlobalContext } from "../../context/GlobalContext";
import ExpenseForm from "../Form/ExpenseForm";
import { useEffect } from "react";
import ExpenseItem from "../ExpenseItem/ExpenseItem";
import { dollar } from "../../utils/icons";

const ExpensesStyled = styled.div`
	display: flex;
	overflow: auto;
	.expense-content {
		display: flex;
		gap: 2rem;
		.expenses {
			flex: 1;
		}
	}
	.total-expense {
		display: flex;
		justify-content: center;
		align-items: center;
		background: #fcf6f9;
		border: 2px solid #fff;
		box-shadow: 0px 1px 15px rgba(0, 0, 0, 0.06);
		border-radius: 20px;
		padding: 1rem;
		margin: 1rem 0;
		font-size: 2rem;
		gap: 0.5rem;
		span {
			font-size: 2.5rem;
			color: var(--color-green);
			font-weight: 800;
		}
	}
`;

function Expenses() {
	const { expenses, getExpenses, deleteExpense, totalExpense } =
		useGlobalContext();
	useEffect(() => {
		getExpenses();
	}, []);
	return (
		<ExpensesStyled>
			<InnerLayout>
				<h1>Expenses</h1>
				<h2 className="total-expense">
					Total Expense:{" "}
					<span>
						{dollar}
						{totalExpense()}
					</span>
				</h2>
				<div className="expense-content">
					<div className="form-container">
						<ExpenseForm />
					</div>
					<div className="expenses">
						{expenses.map((expense) => {
							return (
								<ExpenseItem
									key={expense.id}
									id={expense.id}
									title={expense.title}
									amount={expense.amount}
									date={expense.date}
									category={expense.category}
									description={expense.description}
									indicator={"var(--color-green)"}
									deleteItem={deleteExpense}
								/>
							);
						})}
					</div>
				</div>
			</InnerLayout>
		</ExpensesStyled>
	);
}

export default Expenses;
