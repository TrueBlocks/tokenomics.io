import React, { useEffect } from 'react';
import { Layout } from 'antd';
import { Affix } from 'antd';

import './App.css';
import 'antd/dist/antd.css';
import { Head } from './Header';

import { RightSider } from './RightSider';
import { LeftSider } from './LeftSider';
import { Foot } from './Footer';
import { HomePage } from './HomePage';
import { useGlobalState } from './GlobalState';
import { useNavigate, useParams } from 'react-router-dom';

const { Content } = Layout;

function App() {
  const navigate = useNavigate();
  const { setProject, project } = useGlobalState();
  const { project: projectParam } = useParams();

  // Set the project from URL parameter or, if missing, change
  // the URL so that it has the default project
  useEffect(() => {
    if (!projectParam) {
      navigate(`/${project}`, { replace: true });
    }
    setProject(projectParam || project);
  }, [navigate, project, projectParam, setProject]);

  return (
      <Layout className='App' style={{ backgroundColor: '#e7dddc' }}>
        <Affix offsetTop={0}>
          <Head />
        </Affix>
        <Layout>
          <LeftSider />
          <Content style={{ padding: '10px', paddingLeft: '10px', backgroundColor: 'lightblue' }}>
            <Layout>
              <HomePage />
              <RightSider />
            </Layout>
          </Content>
        </Layout>
        <Affix offsetBottom={0}>
          <Foot />
        </Affix>
      </Layout>
  );
}
export default App;
