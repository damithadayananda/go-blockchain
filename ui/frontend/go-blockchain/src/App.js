import React, { useEffect, useState } from 'react';
import Block from './components/Block';
import TransactionForm from './components/TransactionForm';
import NodeList from "./components/NodeList";
import BlockchainBackground from "./components/Background";
import Transaction from "./components/Transaction";
import {fetchChain, fetchNodes, submitTransaction, fetchTransactions} from './services/api';
import './App.css';
import {Box, Button, Container, List, ListItem, ListItemText, Typography} from "@mui/material";

const App = () => {
    const [blocks, setBlocks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [nodes, setNodes] = useState([]);
    const [transactions, setTransactions] = useState([]);


    const getChain = async () => {
        try {
            const data = await fetchChain();
            setBlocks(data);
        } catch (error) {
            setError("Error fetching the blockchain data.");
        } finally {
            setLoading(false);
        }
    };

    const getNodes = async () => {
        try {
            const data = await fetchNodes();
            setNodes(data);
        } catch (error) {
            setError("Error fetching the blockchain data.");
        } finally {
            setLoading(false);
        }
    };

    const getTransactions = async () => {
        try {
            const data = await fetchTransactions();
            setTransactions(data);
        } catch (error) {
            setError("Error fetching the transactions data.");
        }
    };

    useEffect(() => {
        getChain();
    }, []);

    useEffect(() => {
        getNodes();
    }, []);

    useEffect(() => {
        getTransactions();
    }, [])

    const handleTransactionSubmit = async (transaction) => {
        try {
            await submitTransaction(transaction);
            await Promise.all([getChain(), getTransactions()]);
        } catch (error) {
            setError("Error submitting the transaction.");
        }
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <Container className="app" sx={{ display: 'flex', height: '100vh', position: "relative" }}>
            <BlockchainBackground/>
            <Box className="column column-3-4" sx={{ flexGrow: 1, flexShrink: 1, marginRight: 2 , position: "relative", zIndex: 1}}>
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
                    <Typography variant="h4" sx={{ fontWeight: 'bold', color: '#1976d2' }}>
                        Blocks
                    </Typography>
                    <Button variant="contained" color="primary" onClick={getChain}>
                        RE SYNC
                    </Button>
                </Box>
                {blocks.map((block, index) => (
                    <Block key={index} block={block} />
                ))}
                <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2, mt: 4 }}>
                    <Typography variant="h4" sx={{ fontWeight: 'bold', color: '#1976d2' }}>
                        Mempool Transactions
                    </Typography>
                    <Button variant="contained" color="primary" onClick={getTransactions}>
                        REFRESH
                    </Button>
                </Box>
                <Transaction transactions={transactions} />
            </Box>
            <Box className="column column-1-4" sx={{ width: '25%', minWidth: 250, display: 'flex', flexDirection: 'column', gap: 2, position: "relative", zIndex: 1  }}>
                <Box className="row row-1-2" sx={{ flex: 0, position: "relative", zIndex: 2 }}>
                    <TransactionForm onSubmit={handleTransactionSubmit} />
                </Box>
                <NodeList nodes={nodes}/>
            </Box>
        </Container>
    );
};

export default App;
