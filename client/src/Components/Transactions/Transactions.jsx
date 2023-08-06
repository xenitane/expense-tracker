import { styled } from "styled-components";
import { InnerLayout } from "../../styles/Layouts";

const TransactionsStyled = styled.div``;

function Transactions() {
	return (
		<TransactionsStyled>
			<InnerLayout>
				<h1>Transactions</h1>
				<div className="income-content">
					<div className="form-container"></div>
					<div className="incomes"></div>
				</div>
			</InnerLayout>
		</TransactionsStyled>
	);
}

export default Transactions;
