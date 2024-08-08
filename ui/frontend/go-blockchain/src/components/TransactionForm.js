import React, { useState } from 'react';
import {Box, Button, TextField} from "@mui/material";

const TransactionForm = ({ onSubmit }) => {
    const [amount, setAmount] = useState('');
    const [receiver, setReceiver] = useState('');
    const [sender, setSender] = useState('');
    const [fee, setFee] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit({ amount: parseFloat(amount), receiver, sender, fee: parseFloat(fee)});
        setAmount('');
        setReceiver('');
        setSender('');
        setFee('');
    };

    return (
        <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}>
            <TextField
                label="Amount"
                variant="outlined"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                type="number"
                required
            />
            <TextField
                label="Receiver"
                variant="outlined"
                value={receiver}
                onChange={(e) => setReceiver(e.target.value)}
                required
            />
            <TextField
                label="Sender"
                variant="outlined"
                value={sender}
                onChange={(e) => setSender(e.target.value)}
                required
            />
            <TextField
                label="Fee"
                variant="outlined"
                value={fee}
                onChange={(e) => setFee(e.target.value)}
                type="number"
                required
            />
            <Button variant="contained" color="primary" type="submit">
                Submit Transaction
            </Button>
        </Box>
    );
};

export default TransactionForm;
