import styled from "styled-components";
import { useGlobalContext } from "../../context/GlobalContext";

const HistoryStyled = styled.div`
	display: flex;
	flex-direction: column;
	gap: 1rem;
	.history-item {
		background: #fcf6f9;
		border: 2px solid #fff;
		box-shadow: 0px 1px 15px #0000000f;
		padding: 1rem;
		border-radius: 20px;
		display: flex;
		justify-content: space-between;
		align-items: center;
		.expense {
			color: red;
		}
		.income {
			color: var(--color-green);
		}
	}
`;

function History() {
	const { transactionHistory } = useGlobalContext();
	const history = transactionHistory();

	return (
		<HistoryStyled>
			<h2>Recent History</h2>

			{history.map((item) => {
				return (
					<div className="history-item" key={item.id}>
						<p className={item.type ? "income" : "expense"}>
							{item.title}
						</p>
						<p className={item.type ? "income" : "expense"}>
							{(item.type ? "+" : "-") + `${item.amount}`}
						</p>
					</div>
				);
			})}
		</HistoryStyled>
	);
}

export default History;
