// src/pages/HomePage.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import NavBar from '../components/NavBar';
import PostList from "../components/PostList";
import Sidebar from "../components/indexSidebar";

import { SidebarContext, SidebarProvider } from '../context/SidebarContext';

const HomePage = () => {
    const [post_ids, setPostIds] = useState([]);

    // 主要用于处理组件渲染之外的副作用，例如数据获取、事件监听、订阅等。
    useEffect(() => {
        // 模拟请求 API
        fetch('/api/v1/recommend/blog/latest')
            .then(response => response.json())
            .then(data => {
                let ids = data['ids'];
                if (ids !== undefined && ids !== null) {
                    setPostIds(ids);
                } else {
                    console.log("No data get");
                }

            });
    }, []);

    return (
        <SidebarProvider>
            <div className='container'>
                <NavBar />

                <div className='content'>
                    <SidebarContext.Consumer>
                        {({ sidebarVisible }) => (
                            <div className={`main ${sidebarVisible ? '' : 'full-width'}`}>
                                <h1>绯境之外 第3代 站点</h1>

                                <h2>Recommend Posts List</h2>

                                <h2>Latest Posts List</h2>
                                <PostList
                                    ids={post_ids}
                                />
                            </div>
                        )}
                    </SidebarContext.Consumer>
                    <Sidebar />

                </div>


            </div>
        </SidebarProvider>
    );
};

export default HomePage;
