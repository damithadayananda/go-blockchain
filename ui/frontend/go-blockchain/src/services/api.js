import axios from 'axios';

const API_URL = 'https://localhost:8080/chain';
const TRANSACTION_URL = 'https://localhost:8080/transaction/add';
const NODE_URL = 'https://localhost:8080/node/get';
const MEMPOOL_URL = 'https://localhost:8080/transaction/get/mempool';


export const fetchChain = async () => {
    const response = await axios.get(API_URL);
    return response.data.result;
};

export const fetchNodes = async () => {
    const response = await axios.get(NODE_URL)
    return response.data.result;
}

export const submitTransaction = async (transaction) => {
    try {
        const response = await axios.post(TRANSACTION_URL, transaction, {
            headers: {
                'Content-Type': 'application/json'
            }
        });
        return response.data;
    } catch (error) {
        console.error("Error submitting the transaction:", error);
        throw error;
    }
};

export const fetchTransactions = async () => {
    const response = await axios.get(MEMPOOL_URL);
    return response.data.Result;
};
