// src/pages/PostPage.js
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import NavBar from '../components/NavBar';
import PostDetail from "../components/PostDetails";
import PostDetailSideBar from "../components/PostDetailSidebar";
import { SidebarContext, SidebarProvider } from '../context/SidebarContext';

const PostPage = () => {
    const { id } = useParams(); // 获取 URL 中的 id 参数
    const defaultPost = {
        title: "Now Loading",
        tags: ["Now Loading"],
        content: "Requesting from remote"
    }
    const [post, setPost] = useState(defaultPost);

    useEffect(() => {
        // 模拟请求 API 获取具体博文
        fetch(`/api/v1/blog/content/${id}`)
            .then(response => response.json())
            .then(data => setPost(data));
    }, [id]);

    return (
        <SidebarProvider>
            <div className='container'>
                <NavBar />
                <div className='content'>
                    <SidebarContext.Consumer>
                        {({ sidebarVisible }) => (
                            <div className={`main ${sidebarVisible ? '' : 'full-width'}`}>
                                <PostDetail
                                    title={post['title']}
                                    tags={post['tags']}
                                    content={post['content']}
                                />
                            </div>
                        )}
                    </SidebarContext.Consumer>
                    <PostDetailSideBar
                        content={post['content']}
                    />
                </div>
            </div>
        </SidebarProvider>
    );
};

export default PostPage;
