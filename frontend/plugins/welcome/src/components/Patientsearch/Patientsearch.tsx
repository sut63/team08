import React, { useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { TableContainer, Table, FormControl, TextField, Button, Grid, Link } from '@material-ui/core';
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
  EntPatient
} from '../../api/models/';
//alert
import Swal from 'sweetalert2'
//icon
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import BackspaceTwoToneIcon from '@material-ui/icons/BackspaceTwoTone';
import AddCircleOutlineTwoToneIcon from '@material-ui/icons/AddCircleOutlineTwoTone';
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
    const auth = React.useState(true);
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

    const [patient, getPatient] = useState<EntPatient>();
    const [name, setName] = React.useState(String);
    const NamehandleChange = (event: any) => {
        setName(event.target.value as string);
    };
    const [loading, setLoading] = useState(true);
    const [search, setSearch] = useState(false);
    const [patients, setPatients] = React.useState<EntPatient[]>([]);

    const getPatients = async () => {
        const res = await http.listPatient({ limit: 10, offset: 0 });
        setPatients(res);
    };
    const searchPatient = async () => {
        const res = await http.getPatient({ id: patient_id })
        if (res != undefined) {
            getPatient(res);
        }
    };
    const open = Boolean(anchorEl);
    const handleClose = () => {
        setAnchorEl(null);
      };
    const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
      };
    useEffect(() => {
        getPatients();
    }, [loading]);

  


  var patient_id = 0
  var status = false
  const checkresearch = async () => {
    console.log(patients);
    patients.map(item => {
        if (status === false) {
            if (item.patientName === name) {
                status = true
                if (item.id != undefined) {
                    patient_id = item.id;
                    console.log(patient_id);
                    searchPatient();
                    console.log(patient);
                }
            }
        }

    })
    if (status === false) {
      setSearch(false);
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
      <Header style={HeaderCustom} title={`ระบบค้นหาข้อมูลการรักษาผู้ป่วยใน`}>
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
            href="/reg"
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
          <FormControl
            fullWidth
            className={classes.margin}
            variant="outlined"
            size="small"
          >
           
            <TextField
            style={{ width: 250 ,marginLeft:7,marginRight:-7,marginTop:5}}
            className={classes.textField}
            value={name}
            id="name"
            name="name"
            variant="outlined"
            onChange={NamehandleChange}
            InputLabelProps={{
                shrink: true,
            }}
            />
             </FormControl>
             &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
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
                        <TableCell align="center">คำนำหน้าชื่อ</TableCell>
                        <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                        <TableCell align="center">เพศ</TableCell>
                        <TableCell align="center">น้ำหนัก</TableCell>
                        <TableCell align="center">ส่วนสูง</TableCell>
                        <TableCell align="center">อายุ</TableCell>
                        <TableCell align="center">กรุ๊ปเลือด</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow >
                        <TableCell align="center">{patient?.edges?.prefix?.pname}</TableCell>
                        <TableCell align="center">{patient?.patientName}</TableCell>
                        <TableCell align="center">{patient?.edges?.gender?.gname}</TableCell>
                        <TableCell align="center">{patient?.patientWeight}</TableCell>
                        <TableCell align="center">{patient?.patientHeight}</TableCell>
                        <TableCell align="center">{patient?.patientAge}</TableCell>
                        <TableCell align="center">{patient?.edges?.bloodtype?.bTname}</TableCell>
                        
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