import React, {  useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { FormControl, TextField, Button,  Grid,Link } from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
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
import { EntOperativerecord } from '../../api/models/EntOperativerecord';
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
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  const [checkoperativename, setOperativenames] = useState(false);
  const [operativerecord, setOperativerecord] = useState<EntOperativerecord[]>([])

  const [operativename, setOperativename] = useState(String);
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getOperativerecords = async () => {
      const res = await http.listOperativerecord({ offset: 0 });
      setLoading(false);
      setOperativerecord(res);
    };
    getOperativerecords();
  }, [loading]);

  const operativenamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setOperativenames(false);
    setOperativename(event.target.value as string);

  };

  const deleteOperativerecord = async (id: number) => {
    const res = await http.deleteOperativerecord({ id: id });
    setLoading(true);
  };

  const checkresearch = async () => {
    var check = false;
    operativerecord.map(item => {
      if (operativename != "") {
        if (item.edges?.operative?.operativeName?.includes(operativename)) {
          setOperativenames(true);
          alertMessage("success", "ค้นหาข้อมูลหัตถการสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ค้นหาข้อมูลหัตถการไม่สำเร็จ");
    }
    console.log(checkoperativename)
    if (operativename == "") {
      alertMessage("info", "แสดงข้อมูลหัตถการทั้งหมดในระบบ");
    }
  };

  // name
function a11yProps(index: any) {
  return {
    id: `scrollable-force-tab-${index}`,
    'aria-controls': `scrollable-force-tabpanel-${index}`,
  };
}
const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  function Logout() {
    removeCookie('ID', { path: '/' })
    removeCookie('Name', { path: '/' })
    removeCookie('Email', { path: '/' })
    window.location.href = "http://localhost:3000/";
  }


  return (
    <Page theme={pageTheme.service}>
       <Header style={HeaderCustom} title={`ระบบค้นหาการทำหัตถการให้ผู้ป่วยใน`}>
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
        <ContentHeader title="">

         
          
      

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/opera" 
          variant="contained"
          color="primary"
          startIcon={<AddCircleOutlineTwoToneIcon/>}
          > 
          เพิ่มข้อมูล 
          </Button>

        <div>&nbsp;&nbsp;&nbsp;</div>
          <Button 
          href="/homenurse" 
          variant="contained"
          color="default"
          startIcon={<BackspaceTwoToneIcon/>}
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
            <div className={classes.paper}><strong>กรอก " รายการหัตถการ " เพื่อทำการค้นหา</strong></div>
            <TextField
            style={{ width: 250 ,marginLeft:7,marginRight:-7,marginTop:5}}
            className={classes.textField}
              id="operative"
              variant="outlined"
              color="primary"
              type="string"
              size="small"
              value={operativename}
              onChange={operativenamehandlehange}
            />
            </FormControl>
           
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <Button  
          onClick={() => {
            checkresearch();
            setSearch(true);
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
                  {  checkoperativename ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                          <TableCell align="center">ลำดับที่</TableCell>
                          <TableCell align="center">ชื่อพยาบาล</TableCell>
                          <TableCell align="center">รหัสพยาบาล</TableCell>
                          <TableCell align="center">รายการหัตถการ</TableCell>
                          <TableCell align="center">เครื่องมือที่ใช้</TableCell>
                          <TableCell align="center">ห้องตรวจที่ใช้</TableCell>
                          <TableCell align="center">วันที่ทำหัตถการ</TableCell>
                        
                          </TableRow>
                        </TableHead>
                        <TableBody>

                        {operativerecord.filter((filter: any) => filter.edges?.operative?.operativeName.includes(operativename)).map((item: any) => (
                            <TableRow key={item.id}>
                             <TableCell align="center">{item.id}</TableCell>
                             <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
                             <TableCell align="center">{item.number}</TableCell>
                              <TableCell align="center">{item.edges?.operative?.operativeName}</TableCell>
                              <TableCell align="center">{item.edges?.tool?.toolName}</TableCell>
                              <TableCell align="center">{item.edges?.examinationroom?.examinationroomName}</TableCell>
                              <TableCell align="center">{moment(item.operativeDateTime).format('DD/MM/YYYY hh:mm')}</TableCell>
                        
                            
                              <Button
                                onClick={() => {
                                   if (item.id === undefined) {
                                    return;
                                    }
                                   deleteOperativerecord(item.id);
                                    }}
                                   style={{ marginLeft: 10 }}
                                    variant="contained"
                                  color="secondary"
                                 >
                                  Delete
                                  </Button>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : operativename == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                              <TableCell align="center">ลำดับที่</TableCell>
                          <TableCell align="center">ชื่อพยาบาล</TableCell>
                          <TableCell align="center">รหัสพยาบาล</TableCell>
                          <TableCell align="center">รายการหัตถการ</TableCell>
                          <TableCell align="center">เครื่องมือที่ใช้</TableCell>
                          <TableCell align="center">ห้องตรวจที่ใช้</TableCell>
                          <TableCell align="center">วันที่ทำหัตถการ</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {operativerecord.map((item: any) => (
                                <TableRow key={item.id}>
                         
                             <TableCell align="center">{item.id}</TableCell>
                             <TableCell align="center">{item.edges?.nurse?.nurseName}</TableCell>
                             <TableCell align="center">{item.number}</TableCell>
                              <TableCell align="center">{item.edges?.operative?.operativeName}</TableCell>
                              <TableCell align="center">{item.edges?.tool?.toolName}</TableCell>
                              <TableCell align="center">{item.edges?.examinationroom?.examinationroomName}</TableCell>
                              <TableCell align="center">{moment(item.operativeDateTime).format('DD/MM/YYYY hh:mm')}</TableCell>
                              <Button
                                onClick={() => {
                                   if (item.id === undefined) {
                                    return;
                                    }
                                   deleteOperativerecord(item.id);
                                    }}
                                   style={{ marginLeft: 10 }}
                                    variant="contained"
                                  color="secondary"
                                 >
                                  Delete
                                </Button>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
}