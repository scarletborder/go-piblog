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
        const totalVisiblePages = Math.min(9, maxPage); // 至少显示9个页码，最多显示 maxPage 个

        let leftLimit = currentPage - Math.floor(totalVisiblePages / 2); // 计算左边限制
        let rightLimit = currentPage + Math.floor(totalVisiblePages / 2); // 计算右边限制

        // 调整当左边超出界限时
        if (leftLimit < 0) {
            rightLimit += Math.abs(leftLimit);
            leftLimit = 0;
        }

        // 调整当右边超出界限时
        if (rightLimit >= maxPage) {
            leftLimit -= (rightLimit - maxPage + 1);
            rightLimit = maxPage - 1;
        }

        leftLimit = Math.max(leftLimit, 0); // 确保左边界不小于0

        // 生成页码按钮
        for (let i = leftLimit; i <= rightLimit; i++) {
            pages.push(
                <button
                    key={i}
                    onClick={() => handlePageClick(i)}
                    style={{
                        margin: "0 5px",
                        backgroundColor: i === currentPage ? 'rgb(101,101,101)' : '#444'
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