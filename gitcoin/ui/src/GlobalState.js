import React, { useReducer, useContext, useMemo, createContext, useEffect, useCallback } from 'react';

export const actionSetLocalExplorer = 'SET_LOCAL_EXPLORER';
export const actionSelectGrant = 'SELECT_GRANT';
export const actionSidebarVisible = 'SIDEBAR_VISIBLE';
export const actionSidebarEnabled = 'SIDEBAR_ENABLED';

const initialState = {
    localExplorer: false,
    selectedGrant: {},
    sidebarVisible: false,
    sidebarEnabled: true,
};

const GlobalStateContext = createContext([initialState, () => { }]);

const GlobalStateReducer = (state, action) => {
    switch (action.type) {
        case actionSetLocalExplorer:
            return {
                ...state,
                localExplorer: action.value
            };
        case actionSelectGrant:
            return {
                ...state,
                selectedGrant: action.value
            };
        case actionSidebarVisible:
            return {
                ...state,
                sidebarVisible: action.value
            };
        case actionSidebarEnabled:
            return {
                ...state,
                sidebarEnabled: action.value
            };
        default:
            return state;
    }
};

export const GlobalStateProvider = ({ children }) => {
    const savedState = JSON.parse(localStorage.getItem('settings') || null);
    const [state, dispatch] = useReducer(GlobalStateReducer, savedState || initialState);
    const value = useMemo(() => [state, dispatch], [state]);

    useEffect(() => {
        localStorage.setItem('settings', JSON.stringify(state));
    }, [state]);

    return <GlobalStateContext.Provider value={value}>{children}</GlobalStateContext.Provider>;
};

export const useGlobalState = () => {
    const [state, dispatch] = useContext(GlobalStateContext);

    const setLocalExplorer = useCallback((value) => dispatch({ type: actionSetLocalExplorer, value }), [dispatch]);
    const selectGrant = useCallback((value) => dispatch({ type: actionSelectGrant, value }), [dispatch]);
    const setSidebarVisible = useCallback((value) => dispatch({ type: actionSidebarVisible, value }), [dispatch]);
    const setSidebarEnabled = useCallback((value) => dispatch({ type: actionSidebarEnabled, value }), [dispatch]);

    return {
        localExplorer: state.localExplorer,
        setLocalExplorer,
        selectedGrant: state.selectedGrant,
        selectGrant,
        sidebarVisible: state.sidebarVisible,
        setSidebarVisible,
        sidebarEnabled: state.sidebarEnabled,
        setSidebarEnabled,
    };
};
