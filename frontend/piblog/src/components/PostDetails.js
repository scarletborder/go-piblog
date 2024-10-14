import React, { ReactElement } from 'react';

import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';  // 支持 GitHub Flavored Markdown (GFM)
import rehypeRaw from 'rehype-raw';  // 允许解析 Markdown 中的 HTML
import rehypeHighlight from 'rehype-highlight';  // 支持代码高亮
import './PostDetail.css';
import ReactDOMServer from 'react-dom/server';


function PostDetail(props) {
    var tag_str = "This blog has no tags";

    if (props.tags !== undefined && props.tags !== null) {
        tag_str = `Tag: ${props.tags.join(" ")}`;
    }

    let cont = (
        <ReactMarkdown
            remarkPlugins={[remarkGfm]}    // 使用 GitHub Flavored Markdown 插件
            rehypePlugins={[rehypeRaw, rehypeHighlight]}  // 支持嵌入 HTML 和代码高亮
        >
            {props.content}
        </ReactMarkdown>
    );

    const htmlString = ReactDOMServer.renderToStaticMarkup(cont);

    // 使用 DOMParser 将 HTML 字符串解析为 DOM 树
    const parser = new DOMParser();
    const doc = parser.parseFromString(htmlString, 'text/html');

    // 生成目录算法，并为 <h1> 到 <h6> 添加 id
    const headings = [];
    const parentStack = [];

    const processHeadings = (node) => {
        const level = parseInt(node.tagName.substring(1), 10); // 从标签名中获取标题层级
        const text = node.textContent || ''; // 获取标题的文本内容

        // 栈处理：弹出比当前层级大的标题（即低层级的标题）
        while (parentStack.length > 0 && parentStack[parentStack.length - 1].level >= level) {
            parentStack.pop();
        }

        // 拼接ID，如果有父标题，则将父标题的ID与当前标题拼接
        const parent = parentStack.length > 0 ? parentStack[parentStack.length - 1].id : '';
        const currentId = `${parent}${parent ? '-' : ''}${text.toLowerCase().replace(/\s+/g, '-')}`;

        // 将当前标题加入数组，并设置 id 属性
        headings.push({ text, level, id: currentId });
        node.setAttribute('id', currentId); // 为当前 heading 添加 id

        // 将当前标题压入栈，作为下一个可能的父标题
        parentStack.push({ id: currentId, level });
    };

    // 查找所有 <h1> 到 <h6> 标签并处理它们
    const headingTags = doc.querySelectorAll('h1, h2, h3, h4, h5, h6');
    headingTags.forEach((node) => processHeadings(node));

    // 输出处理后的 HTML
    const processedHtml = doc.documentElement.outerHTML;

    return (
        <div className='PostDetail'>
            <h1>{props.title}</h1>
            <p className='tags'>{tag_str}</p>
            <div dangerouslySetInnerHTML={{ __html: processedHtml }} />
        </div>
    );
}

export default PostDetail;
