import React from 'react';
import { Box, Table, TableBody, TableCell, TableHead, TableRow, Typography } from '@mui/material';
import 'tailwindcss/tailwind.css';

const Transaction = ({ transactions }) => {
    return (
        <Box
            className="transactions"
            sx={{
                margin: 2,
                padding: 2,
                backgroundColor: 'rgba(255, 255, 255, 0.5)',
                borderRadius: '8px',
                boxShadow: 3,
            }}
        >
            <Table>
                <TableHead>
                    <TableRow sx={{ backgroundColor: 'rgba(25, 118, 210, 0.1)' }}>
                        <TableCell sx={{ fontWeight: 'bold' }}>ID</TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }}>Amount</TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }}>Sender</TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }}>Receiver</TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }}>Fee</TableCell>
                        <TableCell sx={{ fontWeight: 'bold' }}>Status</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {transactions.map((tx) => (
                        <TableRow key={tx.id} sx={{ '&:hover': { backgroundColor: 'rgba(0, 0, 0, 0.05)' } }}>
                            <TableCell>{tx.id.slice(0, 8)}...</TableCell>
                            <TableCell>{tx.amount}</TableCell>
                            <TableCell>{tx.sender}</TableCell>
                            <TableCell>{tx.receiver}</TableCell>
                            <TableCell>{tx.fee}</TableCell>
                            <TableCell>
                <span
                    className={`px-2 py-1 rounded ${
                        tx.miningStatus === 'READY_FOR_MINING' ? 'bg-yellow-500' : 'bg-green-500'
                    } text-white`}
                >
                  {tx.miningStatus}
                </span>
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </Box>
    );
};

export default Transaction;