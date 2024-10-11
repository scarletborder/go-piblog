import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';  // 支持 GitHub Flavored Markdown (GFM)
import rehypeRaw from 'rehype-raw';  // 允许解析 Markdown 中的 HTML
import rehypeHighlight from 'rehype-highlight';  // 支持代码高亮

function PostDetail(props) {
    const tag_str = `Tag: ${props.tags.join(" ")}`;

    return (
        <>
            <h1>{props.title}</h1>
            <p>{tag_str}</p>
            <ReactMarkdown
                remarkPlugins={[remarkGfm]}    // 使用 GitHub Flavored Markdown 插件
                rehypePlugins={[rehypeRaw, rehypeHighlight]}  // 支持嵌入 HTML 和代码高亮
            >
                {props.content}
            </ReactMarkdown>
        </>
    );
}

export default PostDetail;