import React from 'react';
import {Box, Card, CardContent, Typography} from "@mui/material";

const Block = ({ block }) => {
    const formatData = (data) => {
        return (
            <pre>{JSON.stringify(data, null, 2)}</pre>
        );
    };

    return (
        <Card className="block" sx={{ margin: 2, padding: 2 }}>
            <CardContent>
                <Typography variant="h5" component="div">
                    Block
                </Typography>
                <Box sx={{ marginTop: 2 }}>
                    <Typography variant="body2" component="div">
                        <strong>Data:</strong> {Array.isArray(block.Data) ? formatData(block.Data) : block.Data}
                    </Typography>
                    <Typography variant="body2" component="div">
                        <strong>Hash:</strong> {block.Hash}
                    </Typography>
                    <Typography variant="body2" component="div">
                        <strong>Previous Hash:</strong> {block.PreviousHash}
                    </Typography>
                    <Typography variant="body2" component="div">
                        <strong>Merkle Root:</strong> {block.MerkleRoot}
                    </Typography>
                    <Typography variant="body2" component="div">
                        <strong>Timestamp:</strong> {block.Timestamp}
                    </Typography>
                    <Typography variant="body2" component="div">
                        <strong>Nonce:</strong> {block.Nonce}
                    </Typography>
                </Box>
            </CardContent>
        </Card>
    );
};

export default Block;
