import { Link } from "react-router-dom";

function PostListItem({ id, title, tags, brief, c_time }) {
    var tag_str = "This blog has no tags";
    brief = brief.substring(0, 200);
    if (tags !== undefined && tags !== null) {
        tag_str = `Tag: ${tags.join(" | ")}`;
    }
    const date = new Date(c_time);
    let localDateString = date.toLocaleString(undefined, {
        hour12: false, // 24小时制
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        timeZoneName: 'short' // 显示时区
    });

    return (
        <li className="PostListItem">
            <Link to={`/post/${id}`}>
                <h4>{title}</h4>
            </Link>
            <p style={{ color: 'red' }}>{tag_str} {localDateString}</p>
            <p>{brief}</p>
        </li >
    )
}

export default PostListItem;