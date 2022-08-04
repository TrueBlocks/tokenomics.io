import { config } from './Config';
let cache = null;
let cachedProject = '';

export async function loadData(project, chain) {
    if (!config.projects.includes(project)) {
        throw new Error(`Unsupported project: ${project}`);
    }

    if (cachedProject === project && cache) return cache;

    const dataDirectory = config.buildPath(project, 'data');
    const response = await fetch(
        new URL(`${dataDirectory}${chain}/theData.json`, window.location)
    );
    const data = await response.json();

    cache = data;
    cachedProject = project;

    return data;
}