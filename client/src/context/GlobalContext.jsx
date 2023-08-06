import axios from "axios";
import { createContext, useContext, useState } from "react";

const GLobalContext = createContext();

const GlobalProvider = ({ children }) => {
	const [incomes, setIncomes] = useState([]);
	const [expenses, setExpenses] = useState([]);
	const [error, setError] = useState(null);

	const addIncome = async (income) => {
		const response = await axios.post("/income", income).catch((err) => {
			setError(err);
		});
	};

	return (
		<GLobalContext.Provider
			value={{
				addIncome
			}}
		>
			{children}
		</GLobalContext.Provider>
	);
};

export default GlobalProvider;

export const useGlobalContext = () => {
	return useContext(GLobalContext);
};
