// src/pages/HomePage.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import NavBar from '../components/NavBar';
import PostList from "../components/PostList";

const HomePage = () => {
    const [post_ids, setPostIds] = useState([]);

    // 主要用于处理组件渲染之外的副作用，例如数据获取、事件监听、订阅等。
    useEffect(() => {
        // 模拟请求 API
        fetch('/api/v1/recommend/blog/latest')
            .then(response => response.json())
            .then(data => {
                let ids = data['ids'];
                if (ids !== undefined) {
                    setPostIds(ids);
                } else {
                    console.log("No data get");
                }

            });
    }, []);

    return (
        <div className='container'>
            <NavBar />

            <div className='content'>
                <div className='main'>
                    <h1>test site</h1>

                    <h2>Recommend Posts List</h2>

                    <h2>Latest Posts List</h2>
                    <PostList
                        ids={post_ids}
                    />
                </div>

                <div className='sidebar'>
                    aaa
                </div>

            </div>


        </div>
    );
};

export default HomePage;
