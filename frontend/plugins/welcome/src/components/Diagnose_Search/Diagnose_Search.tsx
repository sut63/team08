//style
import React, {  useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button, Grid, Link } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
//api
import { DefaultApi } from '../../api/apis';
//table
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
//entity
import {
  EntDiagnose,
  EntPatient
} from '../../api/models/';//alert
import Swal from 'sweetalert2'
//icon
import AddCircleOutlineTwoToneIcon from '@material-ui/icons/AddCircleOutlineTwoTone';
import BackspaceTwoToneIcon from '@material-ui/icons/BackspaceTwoTone';
import CancelTwoToneIcon from '@material-ui/icons/CancelTwoTone';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import DeleteTwoToneIcon from '@material-ui/icons/DeleteTwoTone';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie
import { Link as RouterLink } from 'react-router-dom';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

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
   width: 500 ,
   marginLeft:7,
   marginRight:-7,
  },
  select: {
    width: 500 ,
    marginLeft:7,
    marginRight:-7,
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

  const [diagnose, getDiagnose] = useState<EntDiagnose>();
  const [name, setName] = React.useState(String);
  const NamehandleChange = (event: any) => {
    setName(event.target.value as string);
  };
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);
  const [diagnoses, setDiagnoses] = React.useState<EntDiagnose[]>(Array);

  const getDiagnoses = async () => {
      const res = await http.listDiagnose({ limit: 10, offset: 0 });
      setDiagnoses(res);
  };
  const searchDiagnose = async () => {
      const res = await http.getDiagnose({ id: diagnose_id })
      if (res != undefined) {
        getDiagnose(res);
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
    getDiagnoses();
  }, [loading]);

var diagnose_id = 0
var status = false
const checkresearch = async () => {
  console.log(diagnoses);
  diagnoses.map(item => {
      if (status === false) {
          if (item.edges?.patient?.patientName == name) {
              status = true
              if (item.id != undefined) {
                diagnose_id = item.id;
                  console.log(diagnose_id);
                  searchDiagnose();
                  console.log(diagnose);
              }
          }
      }

  })
  if (status === false) {
    setSearch(false);
    diagnose_id = 0;
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
      <Header style={HeaderCustom} title={`ระบบค้นหาประวัติการวินิจฉัยโรคผู้ป่วยใน`} >
      <AccountCircleIcon aria-controls="fade-menu" aria-haspopup="true"  fontSize="large" />
        <div style={{ marginLeft: 10 }}> </div>
        <div style={{ marginLeft: 1 }}>{Name}</div>
        <div style={{ marginLeft: 10 }}>
        <Link component={RouterLink} to="/">
             Logout
         </Link>
         </div>

      </Header>
      <Content>
        <ContentHeader title=" ">

        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/dia" 
          variant="contained"
          color="primary"
          startIcon={<AddCircleOutlineTwoToneIcon/>}
          > 
          เพิ่มข้อมูล 
          </Button>

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/homedoctor" 
          variant="contained"
          color="default"
          startIcon={<BackspaceTwoToneIcon/>}
          > 
          ย้อนกลับ 
          </Button>
        </ContentHeader>

        <div className={classes.root}>
        <form noValidate autoComplete="off">
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
            id="number"
            name="number"
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
              startIcon={<SearchTwoToneIcon/>}
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
                                <TableCell align="center">รหัสการวินิจฉัย</TableCell>
                                <TableCell align="center">แพทย์ที่รักษา</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                <TableCell align="center">โรค</TableCell>
                                <TableCell align="center">อาการเพิ่มเติม</TableCell>
                                <TableCell align="center">แผนกการเข้ารักษา</TableCell>
                                <TableCell align="center">หมายเหตุ</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>
                            <TableRow > 
                            <TableCell align="center">{diagnose?.diagnoseID}</TableCell>                          
                            <TableCell align="center">{diagnose?.edges?.doctor?.doctorName}</TableCell>
                            <TableCell align="center">{diagnose?.edges?.patient?.patientName}</TableCell>
                            <TableCell align="center">{diagnose?.edges?.disease?.diseaseName}</TableCell>
                            <TableCell align="center">{diagnose?.diagnoseSymptoms}</TableCell>
                            <TableCell align="center">{diagnose?.edges?.department?.departmentName}</TableCell>
                            <TableCell align="center">{diagnose?.diagnoseNote}</TableCell>
                            <TableCell align="center"/>
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