import axios from "axios";
import { useState } from "react";
import { GlobalContext } from "./GlobalContext";

function GlobalProvider({ children }) {
	const [incomes, setIncomes] = useState([]);
	const [expenses, setExpenses] = useState([]);
	const [error, setError] = useState(null);

	const addIncome = async (income) => {
		await axios
			.post("api/v1/income", income)
			.then(async () => await getIncomes())
			.catch((err) => setError(err));
	};

	const getIncomes = async () => {
		await axios
			.get("api/v1/income")
			.then((res) => {
				setIncomes(res.data.data.incomes);
			})
			.catch((err) => setError(err));
	};

	const deleteIncome = async (id) => {
		await axios
			.delete(`api/v1/income/${id}`)
			.then(async () => {
				await getIncomes();
			})
			.catch((err) => setError(err));
	};

	const totalIncome = () =>
		incomes.reduce((pv, cv) => {
			return pv + cv.amount;
		}, 0);

	const addExpense = async (expense) => {
		await axios
			.post("api/v1/expense", expense)
			.then(async () => await getExpenses())
			.catch((err) => setError(err));
	};

	const getExpenses = async () => {
		await axios
			.get("api/v1/expense")
			.then((res) => {
				setExpenses(res.data.data.expenses);
			})
			.catch((err) => setError(err));
	};

	const deleteExpense = async (id) => {
		await axios
			.delete(`api/v1/expense/${id}`)
			.then(async () => {
				await getExpenses();
			})
			.catch((err) => setError(err));
	};

	const totalExpense = () =>
		expenses.reduce((pv, cv) => {
			return pv + cv.amount;
		}, 0);

	const transactionHistory = () => {
		const history = [
			...incomes.map((income) => ({ ...income, type: true })),
			...expenses.map((expense) => ({ ...expense, type: false })),
		].sort((a, b) => {
			return new Date(b.createdAt) - new Date(a.createdAt);
		});
		return history.slice(0, 3);
	};

	return (
		<GlobalContext.Provider
			value={{
				addExpense,
				getExpenses,
				deleteExpense,
				totalExpense,
				addIncome,
				getIncomes,
				deleteIncome,
				totalIncome,
				expenses,
				incomes,
				transactionHistory,
				error,
				setError,
			}}
		>
			{children}
		</GlobalContext.Provider>
	);
}

export default GlobalProvider;
