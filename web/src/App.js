import * as React from 'react';
import { Container } from '@mui/material';
import MainHeader from './components/MainHeader';
import List from '@mui/material/List';
import SchedulerListItem from './components/SchedulerListItem';
import axios from "axios"

function App() {

  const [schedulers, setSchedulers] = React.useState([]);

  React.useEffect(() => {
    require('dotenv').config()
    let username = process.env.REACT_APP_SCHEDULER_USERNAME
    let password = process.env.REACT_APP_SCHEDULER_PASSWORD

    axios({
      baseURL: 'http://localhost:1929/v1/schedulers',
      headers: {
        'Content-Type': "application/json",
        'Authorization': `Basic ${username}:${password}`
      },
      method: 'get',
    }).then((response) => {
      setSchedulers(response.data.schedulers);
    });
  }, []);

  return (
    <Container fixed>
      <MainHeader/>
      <List sx={{pt:"7%"}}>
        {schedulers.map((scheduler) => <SchedulerListItem scheduler={scheduler}/>)}
      </List>
    </Container>
  );
}

export default App;
