import { styled } from "styled-components";
import avatar from "../../img/avatar.png";
import { menuItems } from "../../utils/menuItems";
import { signout } from "../../utils/icons";

const NavStyle = styled.nav`
	padding: 2rem 1.5rem;
	width: 374px;
	height: 100%;
	background: rgba(252, 246, 249, 0.78);
	border: 3px solid #fff;
	backdrop-filter: blur(4.5px);
	border-radius: 32px;
	display: flex;
	flex-direction: column;
	justify-content: space-between;
	gap: 2rem;
	.user-container {
		height: 100px;
		display: flex;
		align-items: center;
		gap: 1rem;
		img {
			width: 80px;
			height: 80px;
			border-radius: 50%;
			object-fit: cover;
			background: #fcf6f9;
			box-shadow: 0px 1px 17px rgba(0, 0, 0, 0.06);
		}
		h2 {
			color: var(--primary-color);
		}
		p {
			color: var(--primary-color2);
		}
	}
	.menu-items {
		flex: 1;
		display: flex;
		flex-direction: column;
		li {
			display: grid;
			grid-template-columns: 40px auto;
			align-items: center;
			margin: 0.6rem 0;
			font-weight: 500;
			cursor: pointer;
			transition: all 0.4s ease-in-out;
			color: var(--primary-color2);
			padding-left: 1rem;
			position: relative;
			i {
				font-size: 1.4rem;
				transition: all 0.4s ease-in-out;
			}
		}
		li.active {
			color: var(--primary-color);
			&::before {
				content: "";
				position: absolute;
				left: 0;
				top: 0;
				width: 4px;
				height: 100%;
				background: #222260;
				border-radius: 0px 10px 10px 0px;
			}
		}
	}
`;

function Navigation({ active, setActive }) {
	return (
		<NavStyle>
			<div className="user-container">
				<img
					src={avatar}
					alt="user-avatar"
				/>

				<div className="text">
					<h2>
                Tushar
					</h2>

					<p>
    Your Money
					</p>
				</div>
			</div>

			<ul className="menu-items">
				{menuItems.map((mi) => {
					return (
						<li
							key={mi.id}
							onClick={() => {
								setActive(mi.id);
							}}
							className={active === mi.id ? "active" : ""}
						>
							{mi.icon}

							<span> 
								{" "}

								{mi.title}
							</span>
						</li>
					);
				})}
			</ul>

			<div className="bottom-nav">
				<li>
					{signout}

					{" "}
                Sign Out
				</li>
			</div>
		</NavStyle>
	);
}

export default Navigation;
