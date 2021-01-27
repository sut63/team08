import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircle from '@material-ui/icons/AccountCircle';
import ApartmentIcon from '@material-ui/icons/Apartment';
import LoupeIcon from '@material-ui/icons/Loupe';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import IconButton from '@material-ui/core/IconButton';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import HowToRegIcon from '@material-ui/icons/HowToReg';
import Typography from '@material-ui/core/Typography';
import {
  Container,
  Grid,
  Table,
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { Cookies } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie


const cookies = new Cookies();
const Name = cookies.get('Name');

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
    width: 1000,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },

}));


function a11yProps(index: any) {
  return {
    id: `scrollable-force-tab-${index}`,
    'aria-controls': `scrollable-force-tabpanel-${index}`,
  };
}

export default function ScrollableTabsButtonForce() {
  const classes = useStyles();
  const [value, setValue] = React.useState(0);
  const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  function Logout() {
    removeCookie('ID', { path: '/' })
    removeCookie('Name', { path: '/' })
    removeCookie('Email', { path: '/' })
    window.location.href = "http://localhost:3000/";
  }
  
  const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
    setValue(newValue);
  };

  const [auth, setAuth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (

    <Page theme={pageTheme.service}>
      <Header style={HeaderCustom} title={`เจ้าหน้าที่เวชระเบียน`}>
       
         
          <div style={{ marginLeft: 1 }}>{Name}</div>
     

          {auth && (
            <div>
              <IconButton
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
                color="inherit"
              >
                <AccountCircle fontSize="large"/>
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={open}
                onClose={handleClose}
              >
                <MenuItem component={RouterLink} to="/homemedical"> Home </MenuItem>
                <MenuItem onClick={Logout}> Logout </MenuItem>
              </Menu>
            </div>
          )}
      </Header>



      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={1}>
            <Grid item xs={12}>
              <Table >
                <tr><td></td>
                  <td align="center">
                    <br></br>
                    <ApartmentIcon color="primary" style={{ fontSize: 150 }} />
                    <Typography component="h1" variant="h5">
                      SE63 Hospital
     </Typography><br></br>
                  </td>
                </tr>

                <tr><td></td><td>
                  <hr ></hr></td>
                </tr>
                <Grid item xs={5}>
                  <div className={classes.paper}> </div>
                </Grid>
                <tr><br></br></tr>
                <tr><td></td>
                  <td>
                    <AppBar position="static" color="default">
                      <Tabs
                        value={value}
                        onChange={handleChange}
                        variant="scrollable"
                        scrollButtons="on"
                        indicatorColor="primary"
                        textColor="primary"
                        aria-label="scrollable force tabs example"
                      >
                        <Tab label="ระบบบันทึกสิทธิการรักษาพยาบาล" icon={<LoupeIcon />} {...a11yProps(0)} component={RouterLink} to="/covered" />
                        <Tab label="ระบบค้นหาสิทธิการรักษาพยาบาล" icon={<HowToRegIcon />} {...a11yProps(4)} component={RouterLink} to="/coveredsearch" />
                      </Tabs>
                    </AppBar>
                  </td>
                </tr>

              </Table>
            </Grid>


          </Grid>
        </Container>
      </Content>

    </Page>
  );
}