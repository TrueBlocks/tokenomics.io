import React, { useReducer, useContext, useMemo, createContext, useEffect, useCallback } from 'react';
import { loadData } from './DataSource'
import { config } from './Config';

export const actionSetProject = 'SET_PROJECT';
export const actionSetLocalExplorer = 'SET_LOCAL_EXPLORER';
export const actionSetShowZero = 'SET_SHOW_ZERO';
export const actionSetChain = 'SET_CHAIN';
export const actionSelectGrant = 'SELECT_GRANT';
export const actionSidebarVisible = 'SIDEBAR_VISIBLE';
export const actionSidebarEnabled = 'SIDEBAR_ENABLED';

const initialState = {
    project: config.defaultProject(),
    localExplorer: false,
    showZero: false,
    chain: "mainnet",
    selectedGrant: {},
    sidebarVisible: false,
    sidebarEnabled: true,
};

const GlobalStateContext = createContext([initialState, () => { }]);

const GlobalStateReducer = (state, action) => {
    switch (action.type) {
        case actionSetProject:
            const project = action.value
            if (!config.projects.includes(project)) {
                throw new Error(`Wrong value for 'project': ${project}`);
            }
            return {
                ...state,
                project,
            };
        case actionSetLocalExplorer:
            return {
                ...state,
                localExplorer: action.value
            };
        case actionSetShowZero:
            return {
                ...state,
                showZero: action.value
            };
        case actionSetChain:
            return {
                ...state,
                chain: action.value
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
    const [state, dispatch] = useReducer(GlobalStateReducer, { ...initialState, ...savedState });
    const value = useMemo(() => [state, dispatch], [state]);

    useEffect(() => {
        localStorage.setItem('settings', JSON.stringify(state));
    }, [state]);

    return <GlobalStateContext.Provider value={value}>{children}</GlobalStateContext.Provider>;
};

export const useGlobalState = () => {
    const [state, dispatch] = useContext(GlobalStateContext);

    const setProject = useCallback((value) => dispatch({ type: actionSetProject, value }), [dispatch]);
    const setLocalExplorer = useCallback((value) => dispatch({ type: actionSetLocalExplorer, value }), [dispatch]);
    const setShowZero = useCallback((value) => dispatch({ type: actionSetShowZero, value }), [dispatch]);
    const setChain = useCallback((value) => dispatch({ type: actionSetChain, value }), [dispatch]);
    const selectGrant = useCallback((value) => dispatch({ type: actionSelectGrant, value }), [dispatch]);
    const setSidebarVisible = useCallback((value) => dispatch({ type: actionSidebarVisible, value }), [dispatch]);
    const setSidebarEnabled = useCallback((value) => dispatch({ type: actionSidebarEnabled, value }), [dispatch]);

    return {
        project: state.project,
        setProject,
        localExplorer: state.localExplorer,
        setLocalExplorer,
        showZero: state.showZero,
        setShowZero,
        chain: state.chain,
        setChain,
        selectedGrant: state.selectedGrant,
        selectGrant,
        sidebarVisible: state.sidebarVisible,
        setSidebarVisible,
        sidebarEnabled: state.sidebarEnabled,
        setSidebarEnabled,
    };
};

export const useGlobalGrantsData = (project, chain) => {
    return loadData(project, chain);
}

export const getChainData = (grantData) => {
    var chain = grantData.curChain;
    if (!grantData.chainData) return grantData.chainData[0];
    for (var i = 0; i < grantData.chainData.length; i++) {
        if (grantData.chainData[i].chainName === chain) {
            return grantData.chainData[i]
        }
    }
    return grantData.chainData[0]
}
