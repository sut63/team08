import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page,pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntTool} from '../../api/models/EntTool'; // import interface User
import { EntOperative } from '../../api/models/EntOperative'; // import interface Car
import { EntExaminationroom } from '../../api/models/EntExaminationroom'; // import interface job
import { EntNurse } from '../../api/models/EntNurse'; // import interface job
// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));


const Operativerecord: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

 
  const [examinationrooms, setExaminationrooms] = React.useState<EntExaminationroom[]>([]);
  const [operatives, setOperatives] = React.useState<EntOperative[]>([]);
  const [tools, setTools] = React.useState<EntTool[]>([]);
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);
    
  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getExaminationrooms = async () => {
    const res = await http.listExaminationroom({ limit: 2, offset: 0 });
    setExaminationrooms(res);
  };

  const getTools = async () => {
    const res = await http.listTool({ limit: 3, offset: 0 });
    setTools(res);
  };

  const getOperatives = async () => {
    const res = await http.listOperative({ limit: 7, offset: 0 });
    setOperatives(res);
  };

  const getNurse = async () => {
    const res = await http.listNurse({ limit: 2, offset: 0 });
    setNurses(res);
  };

  const OperativerecordDateTimehandleChange = (event: any) => {
    setOperativerecordDateTime(event.target.value as string);
  };
  const ExaminationroomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setExaminationroomID(event.target.value as number);
  };
  const ToolhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setToolID(event.target.value as number);
  };
  const OperativehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setOperativeID(event.target.value as number);
  };
  const NursehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setNurseID(event.target.value as number);
  };

  const [added, setOperativerecordDateTime] = React.useState(String);
  const [examinationroomID, setExaminationroomID] = React.useState(Number);
  const [toolID, setToolID] = React.useState(Number);
  const [operativeID, setOperativeID] = React.useState(Number);
  const [nurseID, setNurseID] = React.useState(Number);



  // Lifecycle Hooks
  useEffect(() => {
    getExaminationrooms();
    getOperatives();
    getTools();
    getNurse();
  }, []);


  // function save data
  function save() {
    const operativerecords = {
      added: added + ":00+00:00",
      tool: toolID,
      examinationroom: examinationroomID,
      operative: operativeID,
      nurse: nurseID,
    }
    const apiUrl = 'http://localhost:8080/api/v1/operativerecords';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(operativerecords),
    };
    console.log(operativerecords); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
        
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }

      });
}
  
  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`Operativerecord system`}>
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>ชื่อพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="nurse"
                  value={nurseID}
                  onChange={NursehandleChange}
                >
                  {nurses.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nurseName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>รายการหัตถการ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="operative"
                  value={operativeID}
                  onChange={OperativehandleChange}
                >
                  {operatives.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.operativeName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เครื่องมือ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="tool"
                  value={toolID}
                  onChange={ToolhandleChange}
                >
                  {tools.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.toolName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>ห้องตรวจ</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel></InputLabel>
                <Select
                  name="examinationroom"
                  value={examinationroomID}
                  onChange={ExaminationroomhandleChange}
                >
                  {examinationrooms.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.examinationroomName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>วันที่และเวลา</div>
            </Grid>
            <Grid item xs={8}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่และเวลา"
                  name="added"
                  type="datetime-local"
                  value={added}
                  onChange={OperativerecordDateTimehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </form>
            </Grid>
            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
            <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                SAVE
              </Button>
          </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Operativerecord;