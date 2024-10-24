import React, { useState, useEffect, useRef, useContext } from "react";
import remarkParse from 'remark-parse';
import { unified } from 'unified';
import { SidebarContext } from '../context/SidebarContext';
import { visit } from "unist-util-visit";

function PostDetailSideBar({ content }) {
    const { sidebarVisible } = useContext(SidebarContext);

    const [toc, setToc] = useState([]);
    const [copySuccess, setCopySuccess] = useState(""); // 控制提示框的状态
    const [showTooltip, setShowTooltip] = useState(false); // 控制是否显示提示框
    const tooltipRef = useRef(null); // 引用提示框元素
    const buttonRef = useRef(null); // 引用按钮元素

    // 提取标题
    useEffect(() => {
        const extractHeadings = () => {
            const tree = unified().use(remarkParse).parse(content);
            const headings = [];
            const parentStack = []; // 栈，存储当前层级中的父标题

            visit(tree, "heading", (node) => {
                if (node.children[0] === undefined) {
                    return;
                }
                const text = node.children[0].value; // 获取标题的文本内容
                const level = node.depth; // 获取标题的层级

                // 栈处理：弹出比当前层级大的标题（即低层级的标题）
                while (parentStack.length > 0 && parentStack[parentStack.length - 1].level >= level) {
                    parentStack.pop();
                }

                // 拼接ID，如果有父标题，则将父标题的ID与当前标题拼接
                const parent = parentStack.length > 0 ? parentStack[parentStack.length - 1].id : '';
                const currentId = `${parent}${parent ? '-' : ''}${text.toLowerCase().replace(/\s+/g, '-')}`;

                // 将当前标题加入队列
                headings.push({ text, level, id: currentId });

                // 将当前标题压入栈，作为下一个可能的父标题
                parentStack.push({ id: currentId, level });
            });

            setToc(headings); // 设置最终生成的目录
        };

        extractHeadings();
    }, [content]);

    // 复制当前页面的链接
    const handleShareClick = (e) => {
        e.stopPropagation(); // 防止事件传播
        const pageUrl = window.location.href; // 获取当前页面的 URL

        if (navigator.clipboard && navigator.clipboard.writeText) {
            // 如果 navigator.clipboard 可用，直接复制链接
            navigator.clipboard.writeText(pageUrl).then(() => {
                setCopySuccess("Link copied successfully!"); // 设置成功提示
                setShowTooltip(true); // 显示提示框
            }).catch(err => {
                setCopySuccess("Failed to copy the link"); // 设置失败提示
                setShowTooltip(true); // 显示提示框
            });
        } else {
            // navigator.clipboard 不可用时的回退方案
            const tempInput = document.createElement("input");
            tempInput.style.position = "fixed";
            tempInput.style.left = "-9999px"; // 隐藏输入框
            tempInput.value = pageUrl;
            document.body.appendChild(tempInput);
            tempInput.select(); // 自动选中输入框内容

            try {
                document.execCommand("copy"); // 尝试复制
                setCopySuccess("Link copied successfully!"); // 设置成功提示
            } catch (err) {
                setCopySuccess("Failed to copy. Please copy manually."); // 设置失败提示
            }

            document.body.removeChild(tempInput); // 移除临时输入框
            setShowTooltip(true); // 显示提示框
        }
    };

    // 调整提示框位置，确保提示框不会超出页面右边界
    useEffect(() => {
        if (showTooltip && tooltipRef.current && buttonRef.current) {
            const tooltip = tooltipRef.current;
            const button = buttonRef.current;
            const tooltipRect = tooltip.getBoundingClientRect();
            const buttonRect = button.getBoundingClientRect();
            const pageWidth = window.innerWidth;

            // 默认情况下 tooltip left 对齐
            tooltip.style.left = "0";

            // 如果提示框右侧超出页面边界
            if (buttonRect.left + tooltipRect.width > pageWidth) {
                tooltip.style.left = `-${tooltipRect.width - buttonRect.width}px`; // 向左调整
            }
        }
    }, [showTooltip]);

    // 点击页面的其他地方时，关闭提示框
    useEffect(() => {
        const handleClickOutside = (event) => {
            if (tooltipRef.current && !tooltipRef.current.contains(event.target)) {
                setShowTooltip(false); // 点击外部时关闭提示框
            }
        };

        if (showTooltip) {
            document.addEventListener("click", handleClickOutside);
        }

        // 移除事件监听器
        return () => {
            document.removeEventListener("click", handleClickOutside);
        };
    }, [showTooltip]);

    return (sidebarVisible && (
        <div className="sidebar" style={{ position: 'relative' }}>
            <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                <h2 style={{ margin: 0 }}>Table of Contents</h2>
                <div style={{ position: 'relative' }}>
                    <button
                        onClick={handleShareClick}
                        ref={buttonRef} // 为按钮添加引用
                        style={{
                            background: 'none',
                            border: 'none',
                            cursor: 'pointer',
                            padding: '5px'
                        }}
                    >
                        📋 {/* 这个是分享图标，可以根据需要替换成其他图标 */}
                    </button>
                    {/* 提示框：根据状态显示或隐藏 */}
                    {showTooltip && (
                        <div
                            ref={tooltipRef} // 为提示框添加引用
                            style={{
                                position: 'absolute',
                                top: '100%',
                                left: '0',
                                backgroundColor: '#333',
                                color: '#fff',
                                padding: '5px 10px',
                                borderRadius: '4px',
                                whiteSpace: 'nowrap',
                                zIndex: '1000' // 确保提示框在最上层
                            }}
                        >
                            {copySuccess}
                        </div>
                    )}
                </div>
            </div>
            <ul>
                {toc.map((heading) => (
                    <li key={heading.id} style={{ marginLeft: (heading.level - 1) * 20 }}>
                        <a href={`#${heading.id}`}>{heading.text}</a>
                    </li>
                ))}
            </ul>
        </div>
    ));
}

export default PostDetailSideBar;
