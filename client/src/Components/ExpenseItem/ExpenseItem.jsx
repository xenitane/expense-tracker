import { styled } from "styled-components";
import {
	book,
	circle,
	clothing,
	food,
	freelance,
	medical,
	takeaway,
	tv,
	dollar,
	calender,
	comment,
	trash
} from "../../utils/icons";
import Button from "../Button/Button";

const IncomeItemStyled = styled.div`
	background: #fcf6f9;
	border: 2px solid #fff;
	box-shadow: 0px 1px 15px rgba(0, 0, 0, 0.06);
	border-radius: 20px;
	padding: 1rem;
	margin-bottom: 1rem;
	display: flex;
	align-items: center;
	gap: 1rem;
	width: 100%;
	color: #222260;
	.icon {
		width: 80px;
		height: 80px;
		border-radius: 20px;
		background: #f5f5f5;
		display: flex;
		align-items: center;
		justify-content: center;
		border: 2px solid #fff;
		i {
			font-size: 2.6rem;
		}
	}
	.content {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: 0.2rem;
		h5 {
			font-size: 1.3rem;
			padding-left: 2rem;
			position: relative;
			&::before {
				position: absolute;
				content: "";
				left: 0;
				top: 50%;
				transform: translateY(-50%);
				width: 0.8rem;
				height: 0.8rem;
				border-radius: 50%;
				background: ${(props) => props.$indicator};
			}
		}
		.inner-content {
			display: flex;
			justify-content: space-between;
			align-items: center;
			.text {
				display: flex;
				justify-content: center;
				gap: 1.5rem;
				p {
					display: flex;
					align-items: center;
					gap: 0.5rem;
					color: var(--primary-color);
					opacity: 0.8;
				}
			}
		}
	}
`;

function ExpenseItem({
	id,
	title,
	amount,
	date,
	category,
	description,
	deleteItem,
	indicator
}) {
	const categoryIcon = () => {
		switch (category) {
		case "education":
			return book;
		case "groceries":
			return food;
		case "health":
			return medical;
		case "subscriptions":
			return tv;
		case "takeaways":
			return takeaway;
		case "clothing":
			return clothing;
		case "travelling":
			return freelance;
		case "other":
			return circle;
		default:
			return "";
		}
	};

	const handleDelete = async (e) => {
		e.preventDefault();
		await deleteItem(id);
	};

	return (
		<IncomeItemStyled $indicator={indicator}>
			<div className="icon">{categoryIcon()}</div>
			<div className="content">
				<h5>{title}</h5>
				<div className="inner-content">
					<div className="text">
						<p>
							{dollar}
							{amount}
						</p>
						<p>
							{calender}
							{date}
						</p>
						<p>
							{comment}
							{description}
						</p>
					</div>
					<div className="btn-container">
						<Button
							icon={trash}
							bPad={"1rem"}
							bRad={"50%"}
							bg={"var(--primary-color)"}
							color={"#fff"}
							iColor={"#fff"}
							hColor={"var(--color-green)"}
							onClick={handleDelete}
						/>
					</div>
				</div>
			</div>
		</IncomeItemStyled>
	);
}

export default ExpenseItem;
