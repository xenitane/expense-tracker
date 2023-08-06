import { dashboard, transactions, trend, expenses } from "./icons";

export const menuItems = [
	{
		id: 0,
		title: "Dashboard",
		icon: dashboard,
		link: "/dashboard"
	},
	{
		id: 1,
		title: "View Transactions",
		icon: transactions,
		link: "/dashboard"
	},
	{
		id: 2,
		title: "Incomes",
		icon: trend,
		link: "/dashboard"
	},
	{
		id: 3,
		title: "Expenses",
		icon: expenses,
		link: "/dashboard"
	}
];
