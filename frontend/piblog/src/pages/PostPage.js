// src/pages/PostPage.js
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import NavBar from '../components/NavBar';
import PostDetail from "../components/PostDetails";

const PostPage = () => {
    const { id } = useParams(); // 获取 URL 中的 id 参数
    const [post, setPost] = useState(null);

    useEffect(() => {
        // 模拟请求 API 获取具体博文
        fetch(`/api/v1/blog/content/${id}`)
            .then(response => response.json())
            .then(data => setPost(data));
    }, [id]);

    if (!post) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <NavBar />
            <PostDetail
                title={post['title']}
                tags={post['tags']}
                content={post['content']}
            />

        </div>
    );
};

export default PostPage;
