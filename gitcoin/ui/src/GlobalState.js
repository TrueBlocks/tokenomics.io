import React, { useReducer, useContext, useMemo, createContext, useEffect } from 'react';

export const actionSetLocalExplorer = 'SET_LOCAL_EXPLORER';

const initialState = {
    localExplorer: false,
};

const GlobalStateContext = createContext([initialState, () => { }]);

const GlobalStateReducer = (state, action) => {
    switch (action.type) {
        case actionSetLocalExplorer:
            return {
                ...state,
                localExplorer: action.value
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

    const setLocalExplorer = (value) => dispatch({ type: actionSetLocalExplorer, value });

    return {
        localExplorer: state.localExplorer,
        setLocalExplorer,
    };
};
