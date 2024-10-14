// src/components/NavBar.js
import React, { useContext, useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { SidebarContext } from '../context/SidebarContext'

const NavBar = () => {
    const { toggleSidebar, sidebarVisible } = useContext(SidebarContext);

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
            <div className='left_list'>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li style={{
                        color: { statusColor }
                    }}>{apiStatus ? 'online' : 'offline'}</li>
                </ul>
            </div>
            <div className='right_list'>
                <button onClick={toggleSidebar}>
                    {sidebarVisible ? '隐藏侧边栏' : '显示侧边栏'}
                </button>
            </div>

        </nav >
    );
};



export default NavBar;
