import { styled } from "styled-components";
import { InnerLayout } from "../../styles/Layouts";
import { useGlobalContext } from "../../context/GlobalContext";
import IncomeForm from "../Form/IncomeForm";
import { useEffect } from "react";
import IncomeItem from "../IncomeItem/IncomeItem";
import { dollar } from "../../utils/icons";

const IncomesStyled = styled.div`
	display: flex;
	overflow: auto;
	.income-content {
		display: flex;
		gap: 2rem;
		.incomes {
			flex: 1;
		}
	}
	.total-income {
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

function Incomes() {
	const { incomes, getIncomes, deleteIncome, totalIncome } =
		useGlobalContext();
	useEffect(() => {
		getIncomes();
	}, []);
	return (
		<IncomesStyled>
			<InnerLayout>
				<h1>Incomes</h1>
				<h2 className="total-income">
					Total Income:{" "}
					<span>
						{dollar}
						{totalIncome()}
					</span>
				</h2>
				<div className="income-content">
					<div className="form-container">
						<IncomeForm />
					</div>
					<div className="incomes">
						{incomes.map((income) => {
							return (
								<IncomeItem
									key={income.id}
									id={income.id}
									title={income.title}
									amount={income.amount}
									date={income.date}
									category={income.category}
									description={income.description}
									indicator={"var(--color-green)"}
									deleteItem={deleteIncome}
								/>
							);
						})}
					</div>
				</div>
			</InnerLayout>
		</IncomesStyled>
	);
}

export default Incomes;
