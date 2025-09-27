import React, { useState } from 'react';
import { Box, Card, CardContent, Typography, IconButton } from "@mui/material";
import { ExpandLess, ExpandMore } from "@mui/icons-material";

const Block = ({ block }) => {
    const [expanded, setExpanded] = useState(true);

    const toggleExpand = () => {
        setExpanded(!expanded);
    };

    const formatData = (data) => (
        <pre style={{ whiteSpace: "pre-wrap", wordWrap: "break-word" }}>
      {JSON.stringify(data, null, 2)}
    </pre>
    );

    return (
        <Card
            className="block"
            sx={{ margin: 2, padding: 2, backgroundColor: 'rgba(255, 255, 255, 0.5)' }}
        >
            <CardContent>
                <Box sx={{ display: "flex", alignItems: "center", justifyContent: "space-between" }}>
                    <Typography variant="h6" component="div">
                        {expanded ? `Block ${block.Index}` : `Block ${block.Index}`}
                    </Typography>
                    <IconButton onClick={toggleExpand}>
                        {expanded ? <ExpandLess /> : <ExpandMore />}
                    </IconButton>
                </Box>

                {expanded && (
                    <Box sx={{ marginTop: 2 }}>
                        <Typography variant="body2" component="div">
                            <strong>Data:</strong>{" "}
                            {Array.isArray(block.Data) ? formatData(block.Data) : block.Data}
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
                )}
            </CardContent>
        </Card>
    );
};

export default Block;
