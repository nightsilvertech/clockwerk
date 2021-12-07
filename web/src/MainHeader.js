import * as React from 'react';
import { styled, alpha } from '@mui/material/styles';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import InputBase from '@mui/material/InputBase';
import SearchIcon from '@mui/icons-material/Search';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import Tooltip from '@mui/material/Tooltip';
import Modal from '@mui/material/Modal';
import FilledInput from '@mui/material/FilledInput';
import FormControl from '@mui/material/FormControl';
import Button from '@mui/material/Button';
import InputLabel from '@mui/material/InputLabel';
import TextareaAutosize from '@mui/material/TextareaAutosize';
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import Link from '@mui/material/Link';
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import axios from "axios";

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 1000,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

const Search = styled('div')(({ theme }) => ({
  position: 'relative',
  borderRadius: theme.shape.borderRadius,
  backgroundColor: alpha(theme.palette.common.white, 0.15),
  '&:hover': {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  marginRight: theme.spacing(2),
  marginLeft: 0,
  width: '100%',
  [theme.breakpoints.up('sm')]: {
    marginLeft: theme.spacing(3),
    width: 'auto',
  },
}));

const SearchIconWrapper = styled('div')(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: '100%',
  position: 'absolute',
  pointerEvents: 'none',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
  color: 'inherit',
  '& .MuiInputBase-input': {
    padding: theme.spacing(1, 1, 1, 0),
    // vertical padding + font size from searchIcon
    paddingLeft: `calc(1em + ${theme.spacing(4)})`,
    transition: theme.transitions.create('width'),
    width: '100%',
    [theme.breakpoints.up('md')]: {
      width: '20ch',
    },
  },
}));

export default function MainHeader() {
  // modal state
  const [openModal, setOpenModal] = React.useState(false);
  const handleClickOpenModal = () => setOpenModal(true);
  const handleClickCloseModal = () => setOpenModal(false);

  // text inputs
  const [name, setName] = React.useState("");
  const [url, setUrl] = React.useState("");
  const [method, setMethod] = React.useState("get");
  const [retry, setRetry] = React.useState("");
  const [retryThreshold, setRetryThreshold] = React.useState("");
  const [cronSpec, setCronSpec] = React.useState("");
  const [body, setBody] = React.useState("");
  const [persist, setPersist] = React.useState(false);
  const [disabled, setDisabled] = React.useState(false);

  function uuidv4() {
    return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
      (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
    );
  }

  function createScheduler(){
    require('dotenv').config()
    let username = process.env.REACT_APP_SCHEDULER_USERNAME
    let password = process.env.REACT_APP_SCHEDULER_PASSWORD
    let host = process.env.REACT_APP_SCHEDULER_HOST
    let port = process.env.REACT_APP_SCHEDULER_PORT

    if (name == "" || url == "" || method == "" || retry == "" || retryThreshold == "" || cronSpec == "") {
      alert("Please fill the form")
    }else{
      let req = {
        'name': name,
        'url': url,
        'referenceId': uuidv4(),
        'executor': 'http',
        'method': method,
        'body': body,
        'retry': parseInt(retry),
        'retryThreshold': parseInt(retryThreshold),
        'persist': persist,
        'disabled': disabled,
        'spec': cronSpec,
        'headers': ["Content-Type|application/json"],
        'username': username,
        'password': password
      }
      axios({
        baseURL: `http://${host}:${port}/v1/scheduler`,
        headers: {
          'Content-Type': "application/json",
        },
        method: 'post',
        data: req
      }).then((response) => {
        handleClickCloseModal()
      }).then((error) => {
        handleClickCloseModal()
      });
    }
  }

  return (
    <Box sx={{ flexGrow: 1 }}>
      <div>
        <Modal
          open={openModal}
          onClose={handleClickCloseModal}
          aria-labelledby="modal-modal-title"
          aria-describedby="modal-modal-description"
        >
          <Box sx={style}>
            <Typography id="modal-modal-title" variant="h6" component="h2" sx={{pb: 2}}>
              Create New Scheduler
            </Typography>

            <FormControlLabel control={<Switch defaultChecked={false} />} label="Persist" fullWidth sx={{pb: 3}} onChange={(event) => {
              setPersist(event.target.value == "on" ? true : false);
            }}/>
            <FormControlLabel control={<Switch defaultChecked={false} />} label="Disable" fullWidth sx={{pb: 3}} onChange={(event) => {
              setDisabled(event.target.value == "on" ? false : true);
            }}/>

            <FormControl variant="filled" fullWidth sx={{pb: 3}}>
              <InputLabel htmlFor="component-filled">Name</InputLabel>
              <FilledInput id="component-filled" value={name} onChange={(event) => {
                setName(event.target.value);
              }} />
            </FormControl>

            <FormControl variant="filled" fullWidth sx={{pb: 3}}>
              <InputLabel htmlFor="component-filled">Cron Spec</InputLabel>
              <FilledInput id="component-filled" value={cronSpec} onChange={(event) => {
                setCronSpec(event.target.value);
              }} />
              <Link href="https://crontab.guru/" underline="always" target="_blank">
                Cron spec reference
              </Link>
            </FormControl>

            <FormControl variant="filled" fullWidth sx={{pb: 3}}>
              <InputLabel htmlFor="component-filled">Retry</InputLabel>
              <FilledInput id="component-filled" value={retry} onChange={(event) => {
                setRetry(event.target.value);
              }}/>
            </FormControl>

            <FormControl variant="filled" fullWidth sx={{pb: 3}}>
              <InputLabel htmlFor="component-filled">Retry Threshold (In second unit)</InputLabel>
              <FilledInput id="component-filled" value={retryThreshold} onChange={(event) => {
                setRetryThreshold(event.target.value);
              }} />
            </FormControl>

            <FormControl variant="filled" fullWidth sx={{pb: 3}}>
              <InputLabel htmlFor="component-filled">URL</InputLabel>
              <FilledInput id="component-filled" value={url} onChange={(event) => {
                setUrl(event.target.value);
              }} />
            </FormControl>

            <RadioGroup
              aria-label="method"
              defaultValue="get"
              name="radio-buttons-group"
              row
              sx={{pb: 3}}
            >
              <FormControlLabel value="get" control={<Radio />} label="GET" onChange={(event) => {
                setMethod(event.target.value);
              }} />
              <FormControlLabel value="post" control={<Radio />} label="POST" onChange={(event) => {
                setMethod(event.target.value);
              }}/>
              <FormControlLabel value="put" control={<Radio />} label="PUT" onChange={(event) => {
                setMethod(event.target.value);
              }} />
              <FormControlLabel value="delete" control={<Radio />} label="DELETE" onChange={(event) => {
                setMethod(event.target.value);
              }}/>
            </RadioGroup>

            <FormControl variant="filled" fullWidth sx={{pb: 3}}>
              <TextareaAutosize
                aria-label="empty textarea"
                placeholder="Body Request"
                style={{ width: "100%" }}
                value={body} 
                onChange={(event) => {
                  setBody(event.target.value);
                }}
              />
            </FormControl>

            <Button variant="contained" fullWidth onClick={() => createScheduler()}>CREATE</Button>

            <Typography id="modal-modal-description" sx={{ mt: 2 }} color="error">
              Press ESC on keyboard to cancel
            </Typography>
          </Box>
        </Modal>
      </div>
      <AppBar positionFixed>
        <Toolbar>
          <Typography
            variant="h6"
            noWrap
            component="div"
            sx={{ display: { xs: 'none', sm: 'block' } }}
          >
            Clockwerk Scheduler Manager
          </Typography>
          <Search>
            <SearchIconWrapper>
              <SearchIcon />
            </SearchIconWrapper>
            <StyledInputBase
              placeholder="Search…"
              inputProps={{ 'aria-label': 'search' }}
            />
          </Search>
          <Box sx={{ flexGrow: 1 }} />
          <Box sx={{ display: { xs: 'none', md: 'flex' } }}>
            <Tooltip title="Add new scheduler">
              <IconButton 
                onClick={handleClickOpenModal}
                size="large" 
                color="inherit">
                <AddCircleIcon />
              </IconButton>
            </Tooltip>
          </Box>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
