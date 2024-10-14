import React, { useContext, useEffect, useState } from 'react';
import { SidebarContext } from '../context/SidebarContext';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';  // 支持 GitHub Flavored Markdown (GFM)
import rehypeRaw from 'rehype-raw';  // 允许解析 Markdown 中的 HTML
import rehypeHighlight from 'rehype-highlight';  // 支持代码高亮
import { readme } from '../static/about'


function Sidebar() {
    const { sidebarVisible } = useContext(SidebarContext);


    let cont = (
        <ReactMarkdown
            remarkPlugins={[remarkGfm]}    // 使用 GitHub Flavored Markdown 插件
            rehypePlugins={[rehypeRaw, rehypeHighlight]}  // 支持嵌入 HTML 和代码高亮
        >
            {readme}
        </ReactMarkdown>
    );

    return (
        sidebarVisible && (
            <div className="sidebar">
                {cont}
            </div>
        )
    );
}

export default Sidebar;
