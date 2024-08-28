import React, { useEffect, useState } from 'react';
import Block from './components/Block';
import TransactionForm from './components/TransactionForm';
import {fetchChain, fetchNodes, submitTransaction} from './services/api';
import './App.css';
import {Box, Button, Container, List, ListItem, ListItemText, Typography} from "@mui/material";

const App = () => {
    const [blocks, setBlocks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [nodes, setNodes] = useState([]);


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

    useEffect(() => {
        getChain();
    }, []);

    useEffect(() => {
        getNodes();
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
                <Typography variant="h4" sx={{ fontWeight: 'bold', color: '#1976d2' }}>Blockchain Data</Typography>
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
                <Box className="row row-1-2" sx={{ flex: 0, flexDirection: 'column', alignItems: 'flex-start' }}>
                    <Typography variant="h5" sx={{ color: '#1976d2', fontWeight: 'bold', marginBottom: 2 }}>
                        Available Nodes
                    </Typography>
                    <List sx={{ width: '100%', bgcolor: 'background.paper', borderRadius: 2, boxShadow: 3 }}>
                        {nodes.map((node, index) => (
                            <ListItem key={index} sx={{ borderBottom: '1px solid #e0e0e0' }}>
                                <ListItemText primary={node} sx={{ color: '#424242' }} />
                            </ListItem>
                        ))}
                    </List>
                </Box>
            </Box>
        </Container>
    );
};

export default App;
