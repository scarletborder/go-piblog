// src/pages/HomePage.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import NavBar from '../components/NavBar';
import PostList from "../components/PostList";
import Sidebar from "../components/indexSidebar";
import ArchivesScroll from "../components/ArchivesScroll";
import { SidebarContext, SidebarProvider } from '../context/SidebarContext';

function ArchivesPage() {
    const [ids, setIds] = useState([]);
    return (
        <SidebarProvider>
            <div className='container'>
                <NavBar />

                <div className='content'>
                    <SidebarContext.Consumer>
                        {({ sidebarVisible }) => (
                            <div className={`main ${sidebarVisible ? '' : 'full-width'}`}>
                                <h1>归档</h1>
                                <PostList
                                    ids={ids}
                                />
                            </div>
                        )}
                    </SidebarContext.Consumer>
                    <Sidebar />

                </div>
                <ArchivesScroll
                    setIds={setIds}
                />
            </div>
        </SidebarProvider>
    )
}

export default ArchivesPage;