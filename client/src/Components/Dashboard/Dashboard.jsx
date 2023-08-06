import { styled } from "styled-components";
import { InnerLayout } from "../../styles/Layouts";
import Chart from "../Chart/Chart";

const DashboardStyled = styled.div``;

function Dashboard() {
	return (
		<DashboardStyled>
			<InnerLayout>
				<h1>All Transactions</h1>
				<div className="stats-container">
					<div className="chart-container">
						<Chart />
					</div>
				</div>
			</InnerLayout>
		</DashboardStyled>
	);
}

export default Dashboard;
