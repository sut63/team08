import React, { useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { MenuItem, Select, InputLabel, Button,TableContainer, Table, Grid, Link } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
//api
import { DefaultApi } from '../../api/apis';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
//entity
import {
  EntRent,
  EntRoom,
  EntPatient
} from '../../api/models/';
//alert
import Swal from 'sweetalert2'
//icon
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import BackspaceTwoToneIcon from '@material-ui/icons/BackspaceTwoTone';
import AddCircleOutlineTwoToneIcon from '@material-ui/icons/AddCircleOutlineTwoTone';
import moment from 'moment';
// name
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie
const cookies = new Cookies();
const Name = cookies.get('Name');
const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(2),
    },
    insideLabel: {
      margin: theme.spacing(1),
    },
    button: {
      marginLeft: '40px',
    },
    textField: {
      width: 500,
      marginLeft: 7,
      marginRight: -7,
    },
    select: {
      width: 500,
      marginLeft: 7,
      marginRight: -7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    pa: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
  }),
);

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


export default function ComponentsTable() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  const [rents, setRents] = useState<EntRent[]>(Array);
  const [rooms, setRooms] = useState<EntRoom[]>(Array);
  const [room, getRoom] = useState<EntRoom>();
  const [rent, getRent] = useState<EntRent>();
  const [patientID, setPatientID] = React.useState(Number);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const PatienthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPatientID(event.target.value as number);
  };
  const getPatient = async () => {
    const res = await http.listPatient({ limit: 10, offset: 0 });
    setPatients(res);
  };
  const getRents = async () => {
    const res = await http.listRent({ limit: 20, offset: 0 });
    setLoading(false);
    setRents(res);
  };
  const getRooms = async () => {
    const res = await http.listRoom({ limit: 20, offset: 0 });
    setLoading(false);
    setRooms(res);
  };
  const searchRoom = async () => {
    const res = await http.getRoom({ id: room_id })
    if (res != undefined) {
      getRoom(res);
    }
  };
  const searchRent = async () => {
    const res = await http.getRent({ id: rent_id })
    if (res != undefined) {
      getRent(res);
    }
  };
  useEffect(() => {
    getRents();
    getRooms();
    getPatient();

  }, [loading]);

  


  var rent_id = 0
  var room_id = 0
  var status = false
  const checkresearch = async () => {
    console.log(rooms);
    console.log(patients);
    console.log(rents);
    if (status === false) {
      rents.map(item => {
        if (item.edges?.patient?.id === patientID) {
          status = true
          if (item.edges.room?.id != undefined) {
            room_id = item.edges.room?.id
            console.log(room_id);
            searchRoom();
          }
          if (item.id != undefined) {
            rent_id = item.id;
            console.log(rent_id);
            searchRent();
            console.log(rent);
          }
        }
      })
    }
    if (status === false) {
      setSearch(false);
      rent_id = 0;
      room_id = 0;
      Toast.fire({
        icon: 'error',
        title: 'ค้นหาข้อมูลไม่สำเร็จ',
      });
    } else {
      setSearch(true);
      Toast.fire({
        icon: 'success',
        title: 'ค้นหาข้อมูลสำเร็จ',

      });
    }
  };

  
  const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  function Logout() {
    removeCookie('ID', { path: '/' })
    removeCookie('Name', { path: '/' })
    removeCookie('Email', { path: '/' })
    window.location.href = "http://localhost:3000/";
  }

  return (
    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`ระบบค้นหาห้องพักผู้ป่วย`}>
        <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true" fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <div style={{ marginLeft: 1 }}>{Name}</div>
        <div style={{ marginLeft: 10 }}>
          <Link component={RouterLink} to="/">
            Logout
         </Link>
        </div>
      </Header>
      <Content>
        <ContentHeader title="">

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            href="/rent"
            variant="contained"
            color="primary"
            startIcon={<AddCircleOutlineTwoToneIcon />}
          >
            เพิ่มข้อมูล
          </Button>

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            href="/homenurse"
            variant="contained"
            color="default"
            startIcon={<BackspaceTwoToneIcon />}
          >
            ย้อนกลับ
          </Button>
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="on">
          <InputLabel htmlFor="demo-simple-select-readonly-label">Patient's Name</InputLabel>
            <Select
              id="patient_id"
              name="patient_id"
              labelId="demo-simple-select-readonly-label"
              value={patientID}
              onChange={PatienthandleChange}
              style={{ width: 150 }}
            > 
              {patients.map(item => {
                return (
                  <MenuItem key={item.id} value={item.id}>
                    {item.patientName}
                  </MenuItem>
                );
              })}
              

            </Select>
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <Button
              onClick={() => {
                checkresearch();
              }}
              variant="contained"
              color="secondary"
              startIcon={<SearchTwoToneIcon />}

            >
              ค้นหาข้อมูล
              </Button>
          </form>
        </div>

        <Grid container justify="center">
          <Grid item xs={12} md={12}>
            <Paper>
              {search ? (
                <div>
                  <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      <TableCell align="center">รหัสการจองห้องพัก</TableCell>
                      <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                      <TableCell align="center">ชื่อญาติผู้ป่วย</TableCell>
                      <TableCell align="center">เบอร์โทรญาติผู้ป่วย</TableCell>
                      <TableCell align="center">วันเวลาที่บันทึก</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>

                    <TableRow >
                      <TableCell align="center">{rent?.rentId}</TableCell>
                      <TableCell align="center">{rent?.edges?.patient?.patientName}</TableCell>
                      <TableCell align="center">{rent?.kinName}</TableCell>
                      <TableCell align="center">{rent?.kinTel}</TableCell>
                      <TableCell align="center">{moment(rent?.addedTime).format('YYYY/MM/DD HH:mm:ss')}</TableCell>
                      <TableCell align="center" />
                    </TableRow>

                  </TableBody>
                  <TableHead>
                    <TableRow>
                      <TableCell align="center">รหัสห้องพัก</TableCell>
                      <TableCell align="center">ตึก</TableCell>
                      <TableCell align="center">ชั้น</TableCell>
                      <TableCell align="center">พยาลาลประจำห้องพัก</TableCell>
                      <TableCell align="center">เบอร์ติดต่อพยาบาล</TableCell>
                      <TableCell align="center" />
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    <TableRow >
                      <TableCell align="center">{rent?.edges?.room?.name}</TableCell>
                      <TableCell align="center">{rent?.edges?.room?.building}</TableCell>
                      <TableCell align="center">{rent?.edges?.room?.floor}</TableCell>
                      <TableCell align="center">{rent?.edges?.nurse?.nurseName}</TableCell>
                      <TableCell align="center">{rent?.edges?.nurse?.nurseTel}</TableCell>
                      <TableCell align="center" />
                    </TableRow>
                  </TableBody>
                  <TableHead>
                    <TableRow>
                      <TableCell align="center">ประเภทห้อง</TableCell>
                      <TableCell align="center">ห้องน้ำ</TableCell>
                      <TableCell align="center">ห้องสุขา</TableCell>
                      <TableCell align="center">ขนาดห้องพัก</TableCell>
                      <TableCell align="center">อื่นๆ</TableCell>
                      <TableCell align="center" />
                    </TableRow>
                  </TableHead>
                  <TableBody >
                    <TableRow >
                      <TableCell align="center">{room?.edges?.roomtype?.name}</TableCell>
                      <TableCell align="center">{room?.edges?.roomtype?.bathroom}</TableCell>
                      <TableCell align="center">{room?.edges?.roomtype?.toilets}</TableCell>
                      <TableCell align="center">{room?.edges?.roomtype?.areasize}</TableCell>
                      <TableCell align="center">{room?.edges?.roomtype?.etc}</TableCell>
                      <TableCell align="center" />
                    </TableRow>
                  </TableBody>
                </Table>
                </TableContainer>
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}