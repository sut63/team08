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
import moment from 'moment';
//entity
import {
  EntCoveredPerson
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

    const [coveredperson, getCoveredperson] = useState<EntCoveredPerson>();
    const [name, setName] = React.useState(String);
    const NamehandleChange = (event: any) => {
      setName(event.target.value as string);
    };
    const [loading, setLoading] = useState(true);
    const [search, setSearch] = useState(false);
    const [coveredpersons, setCoveredpersons] = React.useState<EntCoveredPerson[]>([]);

    const getCoveredpersons = async () => {
        const res = await http.listCoveredperson({ limit: 10, offset: 0 });
        setCoveredpersons(res);
    };
    const searchCovered = async () => {
        const res = await http.getCoveredperson({ id: covered_id })
        if (res != undefined) {
          getCoveredperson(res);
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
      getCoveredpersons();
    }, [loading]);

  var covered_id = 0
  var status = false
  const checkresearch = async () => {
    console.log(coveredpersons);
    coveredpersons.map(item => {
        if (status === false) {
            if (item.edges?.patient?.patientName == name) {
                status = true
                if (item.id != undefined) {
                  covered_id = item.id;
                    console.log(covered_id);
                    searchCovered();
                    console.log(coveredperson);
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
      <Header style={HeaderCustom} title={`ระบบค้นหาสิทธิการรักษาพยาบาล`}>
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
            href="/covered"
            variant="contained"
            color="primary"
            startIcon={<AddCircleOutlineTwoToneIcon />}
          >
            เพิ่มข้อมูล
          </Button>

          <div>&nbsp;&nbsp;&nbsp;</div>
          <Button
            href="/homemedical"
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
                            <TableCell align="center">ลำดับที่</TableCell>
                            <TableCell align="center">รหัสสิทธิการรักษา</TableCell>
                            <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                            <TableCell align="center">ประเภทสิทธิการรักษาพยาบาล</TableCell>
                            <TableCell align="center">กองทุน</TableCell>
                            <TableCell align="center">คำร้องกองทุน</TableCell>
                            <TableCell align="center">ใบรับรองแพทย์</TableCell>
                            <TableCell align="center">หมายเหตุ</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow >
                        <TableCell align="center">{coveredperson?.id}</TableCell>
                        <TableCell align="center">{coveredperson?.coveredPersonNumber}</TableCell>
                      <TableCell align="center">{coveredperson?.edges?.patient?.patientName}</TableCell>
                      <TableCell align="center">{coveredperson?.edges?.schemeType?.schemeTypeName}</TableCell>
                      <TableCell align="center">{coveredperson?.edges?.fund?.fundName}</TableCell>
                      <TableCell align="center">{coveredperson?.fundTitle}</TableCell>
                      <TableCell align="center">{coveredperson?.edges?.certificate?.certificateName}</TableCell>
                      <TableCell align="center">{coveredperson?.coveredPersonNote}</TableCell>
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
