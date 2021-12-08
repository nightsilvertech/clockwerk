import * as React from 'react';
import ListItem from '@mui/material/ListItem';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemText from '@mui/material/ListItemText';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import DeleteIcon from '@mui/icons-material/Delete';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
import ScheduleSendIcon from '@mui/icons-material/ScheduleSend';
import Switch from '@mui/material/Switch';
import Chip from '@mui/material/Chip';
import Tooltip from '@mui/material/Tooltip';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import axios from "axios";

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

export default function SchedulerListItem(props) {  

  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  // modal state
  const [openModal, setOpenModal] = React.useState(false);
  const handleClickOpenModal = () => setOpenModal(true);
  const handleClickCloseModal = () => setOpenModal(false);


  // delete dialog state
  const [openDialog, setOpenDialog] = React.useState(false);  
  const handleClickOpenDialog = () => setOpenDialog(true);
  const handleClickCloseDialog = () => setOpenDialog(false);

  const [disabled, setDisabled] = React.useState(!props.scheduler.disabled);

  function deleteScheduler(){
    require('dotenv').config()
    let username = process.env.REACT_APP_SCHEDULER_USERNAME
    let password = process.env.REACT_APP_SCHEDULER_PASSWORD
    let host = process.env.REACT_APP_SCHEDULER_HOST
    let port = process.env.REACT_APP_SCHEDULER_PORT

    let req = {
      'id': props.scheduler.id,
      'referenceId': props.scheduler.referenceId,
      'username': username,
      'password': password
    }
    axios({
      baseURL: `http://${host}:${port}/v1/scheduler/${props.scheduler.id}`,
      headers: {
        'Content-Type': "application/json",
      },
      method: 'post',
      data: req
    }).then((response) => {
      handleClickCloseDialog()
    }).then((error) => {
      handleClickCloseDialog()
    });
  }

  function switchScheduler(){
    require('dotenv').config()
    let username = process.env.REACT_APP_SCHEDULER_USERNAME
    let password = process.env.REACT_APP_SCHEDULER_PASSWORD
    let host = process.env.REACT_APP_SCHEDULER_HOST
    let port = process.env.REACT_APP_SCHEDULER_PORT

    let req = {
      'id': props.scheduler.id,
      'referenceId': props.scheduler.referenceId,
      'disabled': disabled,
      'username': username,
      'password': password
    }
    axios({
      baseURL: `http://${host}:${port}/v1/scheduler/toggle/${props.scheduler.id}`,
      headers: {
        'Content-Type': "application/json",
      },
      method: 'post',
      data: req
    }).then((response) => {
      handleClickCloseDialog()
    }).then((error) => {
      handleClickCloseDialog()
    });
  }

  return (
    <ListItem
      secondaryAction={
        <Box component="span">
          <Tooltip title="Turn on/off sheduler">
            <Switch defaultChecked={!props.scheduler.disabled} onChange={(event) => {
              setDisabled(!disabled)
              switchScheduler()
            }}/>
          </Tooltip>

          <Tooltip title="Delete scheduler">
            <IconButton edge="end" aria-label="delete" onClick={handleClickOpenDialog}>
              <DeleteIcon/>
            </IconButton>  
          </Tooltip>
          
          <Tooltip title="Scheduler retry attemps 15">
            <Chip label={props.scheduler.retry} sx={{ml:2}} color="error"/>
          </Tooltip>
        </Box>
      }
    >
      <div>
        <Dialog
          fullScreen={fullScreen}
          open={openDialog}
          onClose={handleClickCloseDialog}
          aria-labelledby="responsive-dialog-title"
        >
          <DialogTitle id="responsive-dialog-title">
            Confirmation
          </DialogTitle>
          <DialogContent>
            <DialogContentText>
              Are you sure want to delete this scheduler?
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button autoFocus onClick={handleClickCloseDialog}>
              NO
            </Button>
            <Button onClick={() => deleteScheduler()} autoFocus>
              YES
            </Button>
          </DialogActions>
        </Dialog>
      </div>

      <div>
        <Modal
          open={openModal}
          onClose={handleClickCloseModal}
          aria-labelledby="modal-modal-title"
          aria-describedby="modal-modal-description"
        >
          <Box sx={style}>
            <Typography id="modal-modal-title" variant="h6" component="h2">
              Scheduler Detail 
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              NAME : {props.scheduler.name}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              ENTRY ID : {props.scheduler.entryId}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              SCH ID : {props.scheduler.id}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              REF ID : {props.scheduler.referenceId}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              EXECUTOR : {props.scheduler.executor}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              SPEC : {props.scheduler.spec}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              ACTIVE : {props.scheduler.disabled == true ? "NO" : "YES"}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              PERSISTED : {props.scheduler.persist == true ? "YES" : "NO"}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              URL : {props.scheduler.url}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              METHOD : {props.scheduler.method}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              RETRY : {props.scheduler.retry}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              RETRY THRESHOLD : {props.scheduler.retryThreshold}
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }} color="error">
              Press ESC on keyboard to close modal
            </Typography>
          </Box>
        </Modal>
      </div>
      
      <ListItemAvatar onClick={handleClickOpenModal}>
        <Avatar>
          <ScheduleSendIcon />
        </Avatar>
      </ListItemAvatar>
      
      <ListItemText
        onClick={handleClickOpenModal}
        primary={props.scheduler.name}
        secondary={props.scheduler.desc}
      />
    </ListItem>
  );
}