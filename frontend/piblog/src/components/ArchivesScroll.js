import React, { useState, useEffect } from "react";

// setContentList() : 直接设置 PostListItem
function ArchivesScroll({ setIds }) {
    // 当前页数，初始值为0（从第0页开始）
    const [currentPage, setCurrentPage] = useState(0);

    // 最大页数，初始值为0
    const [maxPage, setMaxPage] = useState(0);

    useEffect(() => {
        fetch(`/api/v1/archives/${currentPage}`).then(
            response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            }
        ).then(
            data => {
                setMaxPage(data.maxpages);
                setIds(data.ids);
            }
        ).catch(err => {
            console.error(err);
        });
    }, [currentPage, setIds]);

    // 点击页码的处理函数
    const handlePageClick = (page) => {
        setCurrentPage(page);
    };

    // 渲染页码按钮
    const renderPagination = () => {
        const pages = [];
        const leftLimit = Math.max(currentPage - 4, 0); // 向左最多显示 4 页
        const rightLimit = Math.min(currentPage + 4, maxPage - 1); // 向右最多显示 4 页

        // 生成页码按钮
        for (let i = leftLimit; i <= rightLimit; i++) {
            pages.push(
                <button
                    key={i}
                    onClick={() => handlePageClick(i)}
                    style={{
                        margin: "0 5px",
                        backgroundColor: i === currentPage ? "lightblue" : "white"
                    }}
                >
                    {i + 1}
                </button>
            );
        }

        return pages;
    };

    return (
        <div pages>{renderPagination()}</div>
    );
}

export default ArchivesScroll;