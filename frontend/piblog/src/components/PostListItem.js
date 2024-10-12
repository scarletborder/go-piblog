import { Link } from "react-router-dom";

function PostListItem({ id, title, tags, brief }) {
    var tag_str = "This blog has no tags";
    brief = brief.substring(0, 200);
    if (tags !== undefined && tags !== null) {
        tag_str = `Tag: ${tags.join(" | ")}`;
    }

    return (
        <li className="PostListItem">
            <Link to={`/post/${id}`}>
                <h4>{title}</h4>
            </Link>
            <p style={{ color: 'red' }}>{tag_str}</p>
            <p>{brief}</p>
        </li >
    )
}

export default PostListItem;