// src/components/NavBar.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

const NavBar = () => {
    const [apiStatus, setApiStatus] = useState(false);
    const [statusColor, setStatusColor] = useState('red');

    useEffect(() => {
        // 模拟请求 API
        fetch('/api/v1/ping')
            .then(response => response.status
            ).then(code => {
                if (code === 200) {
                    setApiStatus(true);
                    setStatusColor('green');
                } else {
                    setApiStatus(false);
                    setStatusColor('red');
                }
            }).catch(err => {
                setApiStatus(false);
                setStatusColor('red');
                console.error(`Fail to connect with api server ${err}`);
            })
    }, []);



    return (
        <nav className='nav'>
            <ul style={styles.ulStyle}>
                <li style={styles.liStyle}><Link to="/">Home</Link></li>
                <li style={{
                    ...styles.liStyle,
                    color: { statusColor }
                }}>{apiStatus ? 'online' : 'offline'}</li>
            </ul>
        </nav >
    );
};

const styles = {
    ulStyle: {
        display: 'flex',          // 横向排列
        listStyleType: 'none',    // 移除默认列表样式
        padding: 0,               // 移除内边距
    },
    liStyle: {
        marginRight: '20px',
    }
};


export default NavBar;
