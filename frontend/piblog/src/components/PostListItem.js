import { Link } from "react-router-dom";

function PostListItem({ id, title, tags, brief }) {
    return (
        <li key={id}>
            <Link to={`/post/${id}`}>
                <h4>{title}</h4>
            </Link>
            <p style={{ color: 'red' }}>{tags.join(' ')}</p>
            <p>{brief}</p>
        </li >
    )
}

export default PostListItem;