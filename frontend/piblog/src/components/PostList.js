import { useEffect, useState } from "react"
import PostListItem from './PostListItem'

function PostList({ ids }) {
    const [liList, setLiList] = useState([]); // 使用 useState 存储 liList

    useEffect(() => {
        const payload = {
            ids: ids
        }
        // 请求获得briefs
        fetch('/api/v1/blog/info', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        }).then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        }).then(data => {
            data.sort((a, b) => b.c_time - a.c_time);
            const tmpLiList = data.map(item => {
                return (
                    <PostListItem
                        key={item['id']}
                        id={item['id']}
                        title={item['title']}
                        tags={item['tags']}
                        brief={item['brief']}
                        c_time={item['c_time']}
                    />
                );
            });
            setLiList(tmpLiList);
        }).catch(err => {
            console.error(err);
        })
    }, [ids]);

    return (
        <div className="PostList"><ul>
            {liList}
        </ul></div>
    );
}


export default PostList;