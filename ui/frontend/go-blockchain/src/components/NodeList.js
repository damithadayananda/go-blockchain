import React, { useState } from 'react';
import {
    Box,
    Typography,
    List,
    ListItem,
    ListItemText,
    Collapse,
    IconButton,
    Tooltip
} from '@mui/material';
import { ExpandLess, ExpandMore, ContentCopy } from '@mui/icons-material';

const NodeList = ({ nodes }) => {
    const [expandedIndex, setExpandedIndex] = useState(null);

    const handleToggle = (index) => {
        setExpandedIndex(expandedIndex === index ? null : index);
    };

    const decodeBase64 = (data) => {
        try {
            return atob(data);
        } catch (e) {
            return "Invalid Base64 data";
        }
    };

    const handleCopy = (text) => {
        navigator.clipboard.writeText(text)
            .then(() => {
                alert("copied to clipboard!");
            })
            .catch(() => {
                alert("Failed to copy");
            });
    };

    return (
        <Box sx={{ flex: 0, flexDirection: 'column', alignItems: 'flex-start'}}>
            <Typography variant="h5" sx={{ color: '#1976d2', fontWeight: 'bold', marginBottom: 2 }}>
                Available Nodes
            </Typography>
            <List sx={{ width: '100%', backgroundColor: 'rgba(255, 255, 255, 0.5)', borderRadius: 2, boxShadow: 3 }}>
                {nodes.map((node, index) => {
                    const decodedCert = decodeBase64(node.Certificate);
                    return (
                        <React.Fragment key={index}>
                            <ListItem
                                button
                                onClick={() => handleToggle(index)}
                                sx={{ borderBottom: '1px solid #e0e0e0' }}
                            >
                                <ListItemText
                                    primary={node.Ip}
                                    sx={{ color: '#424242' }}
                                />
                                {expandedIndex === index ? <ExpandLess /> : <ExpandMore />}
                            </ListItem>
                            <Collapse in={expandedIndex === index} timeout="auto" unmountOnExit>
                                <Box sx={{ padding: 2, backgroundColor: '#f9f9f9', borderTop: '1px solid #ddd' }}>
                                    <Typography variant="subtitle2" sx={{ fontWeight: 'bold' }}>STATUS:</Typography>
                                    <Typography variant="body2" sx={{ mb: 1 }}>{node.Status}</Typography>

                                    <Typography variant="subtitle2" sx={{ fontWeight: 'bold' }}>IP:</Typography>
                                    <Typography variant="body2" sx={{ mb: 1 }}>{node.Ip}</Typography>

                                    <Typography variant="subtitle2" sx={{ fontWeight: 'bold', display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                                        Certificate:
                                        <Tooltip title="Copy to clipboard">
                                            <IconButton
                                                size="small"
                                                onClick={() => handleCopy(decodedCert)}
                                            >
                                                <ContentCopy fontSize="small" />
                                            </IconButton>
                                        </Tooltip>
                                    </Typography>

                                    <Box
                                        component="pre"
                                        sx={{
                                            backgroundColor: '#272822',
                                            color: '#f8f8f2',
                                            fontSize: '0.85rem',
                                            padding: '10px',
                                            borderRadius: '6px',
                                            overflowX: 'auto',
                                            whiteSpace: 'pre-wrap',
                                            wordBreak: 'break-word',
                                            mb: 1
                                        }}
                                    >
                                        {decodedCert}
                                    </Box>

                                    <Typography variant="subtitle2" sx={{ fontWeight: 'bold' }}>
                                        Address:
                                        <Tooltip title="Copy to clipboard">
                                            <IconButton
                                                size="small"
                                                onClick={() => handleCopy(node.Address)}>
                                                <ContentCopy fontSize="small" />
                                            </IconButton>
                                        </Tooltip>
                                    </Typography>
                                    <Box
                                        component="pre"
                                        sx={{
                                            backgroundColor: '#272822',
                                            color: '#f8f8f2',
                                            fontSize: '0.85rem',
                                            padding: '10px',
                                            borderRadius: '6px',
                                            overflowX: 'auto',
                                            whiteSpace: 'pre-wrap',
                                            wordBreak: 'break-word',
                                            mb: 1
                                        }}
                                    >
                                        {node.Address}
                                    </Box>
                                </Box>
                            </Collapse>
                        </React.Fragment>
                    );
                })}
            </List>
        </Box>
    );
};

export default NodeList;
