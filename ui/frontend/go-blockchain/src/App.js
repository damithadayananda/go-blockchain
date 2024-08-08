import React, { useEffect, useState } from 'react';
import Block from './components/Block';
import TransactionForm from './components/TransactionForm';
import { fetchChain, submitTransaction } from './services/api';
import './App.css';
import {Box, Button, Container, Typography} from "@mui/material";

const App = () => {
    const [blocks, setBlocks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

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

    useEffect(() => {
        getChain();
    }, []);

    const handleTransactionSubmit = async (transaction) => {
        try {
            await submitTransaction(transaction);
            await getChain();
        } catch (error) {
            setError("Error submitting the transaction.");
        }
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <Container className="app" sx={{ display: 'flex', height: '100vh' }}>
            <Box className="column column-3-4" sx={{ flexGrow: 1, flexShrink: 1, marginRight: 2 }}>
                <Typography variant="h4">Blockchain Data</Typography>
                {blocks.map((block, index) => (
                    <Block key={index} block={block} />
                ))}
            </Box>
            <Box className="column column-1-4" sx={{ width: '25%', minWidth: 250, display: 'flex', flexDirection: 'column', gap: 2 }}>
                <Box className="row row-1-2" sx={{ flex: 0 }}>
                    <Button variant="contained" color="primary" onClick={getChain}>
                        GET CHAIN
                    </Button>
                </Box>
                <Box className="row row-1-2" sx={{ flex: 0 }}>
                    <TransactionForm onSubmit={handleTransactionSubmit} />
                </Box>
            </Box>
        </Container>
    );
};

export default App;
