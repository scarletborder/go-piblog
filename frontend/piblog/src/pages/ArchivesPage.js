// src/pages/HomePage.js
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import NavBar from '../components/NavBar';
import PostList from "../components/PostList";
import Sidebar from "../components/indexSidebar";

import { SidebarContext, SidebarProvider } from '../context/SidebarContext';

function ArchivesPage() {
    return (
        <SidebarProvider>
            <div className='container'>
                <NavBar />

                <div className='content'>
                    <SidebarContext.Consumer>
                        {({ sidebarVisible }) => (
                            <div className={`main ${sidebarVisible ? '' : 'full-width'}`}>
                                <h1>归档</h1>
                            </div>
                        )}
                    </SidebarContext.Consumer>
                    <Sidebar />

                </div>


            </div>
        </SidebarProvider>
    )
}

export default ArchivesPage;