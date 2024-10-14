import React, { createContext, useState } from 'react';

// 创建上下文
export const SidebarContext = createContext();

export const SidebarProvider = ({ children }) => {
    const [sidebarVisible, setSidebarVisible] = useState(true);

    const toggleSidebar = () => {
        setSidebarVisible(!sidebarVisible);
    };

    return (
        <SidebarContext.Provider value={{ sidebarVisible, toggleSidebar }}>
            {children}
        </SidebarContext.Provider>
    );
};
