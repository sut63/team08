import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import ApartmentIcon from '@material-ui/icons/Apartment';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import LocalLibraryIcon from '@material-ui/icons/LocalLibrary';
import HealingIcon from '@material-ui/icons/Healing';
import LocalPharmacyIcon from '@material-ui/icons/LocalPharmacy';
import ColorizeIcon from '@material-ui/icons/Colorize';
import AccessibilityIcon from '@material-ui/icons/Accessibility';
import AirlineSeatFlatAngledIcon from '@material-ui/icons/AirlineSeatFlatAngled';
import EmojiPeopleIcon from '@material-ui/icons/EmojiPeople';
import Typography from '@material-ui/core/Typography';
import AirlineSeatIndividualSuiteIcon from '@material-ui/icons/AirlineSeatIndividualSuite';
import AccountCircle from '@material-ui/icons/AccountCircle';
import IconButton from '@material-ui/core/IconButton';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
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
      <Header style={HeaderCustom} title={`พยาบาล`}>
       
         
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
                <MenuItem component={RouterLink} to="/homemenurse"> Home </MenuItem>
                <MenuItem onClick={Logout}> Logout </MenuItem>
              </Menu>
            </div>
          )}
      </Header>




      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={10}>
            <Grid item xs={1}>
              <Table >
                <tr>
                  <td align="center">
                    <br></br>
                    <ApartmentIcon color="primary" style={{ fontSize: 150 }} />
                    <Typography component="h1" variant="h5">
                      SE63 Hospital
     </Typography><br></br>
                  </td>
                </tr>

                <tr>
                  <hr ></hr>
                </tr>
                <Grid item xs={5}>
                  <div className={classes.paper}> </div>
                </Grid>
                <tr><br></br></tr>
                <tr>
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
                        <Tab label="ระบบลงทะเบียนผู้ป่วยใน" icon={<EmojiPeopleIcon />} {...a11yProps(0)} component={RouterLink} to="/reg" />
                        <Tab label="ระบบจองห้องพักผู้ป่วย" icon={<AirlineSeatIndividualSuiteIcon />} {...a11yProps(1)} component={RouterLink} to="/rent" />
                        <Tab label="ระบบบันทึกข้อมูลการทำหัตถการให้ผู้ป่วย" icon={<HealingIcon />} {...a11yProps(2)} component={RouterLink} to="/opera" />
                        <Tab label="ระบบการสั่งยา" icon={<LocalPharmacyIcon />} {...a11yProps(3)} component={RouterLink} to="/pre" />
                      </Tabs>
                    </AppBar>
                  </td>
                </tr>
                <br></br><br></br><br></br><br></br>
                <tr><Typography component="h1" variant="h5">
                  History
     </Typography><br></br></tr>
                <br></br>
                <tr>
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
                        <Tab label="ระบบค้นหาข้อมูลการเข้ารักษาผู้ป่วยใน" icon={<AccessibilityIcon />} {...a11yProps(4)} component={RouterLink} to="/patientsearch" />
                        <Tab label="ระบบค้นหาห้องพักผู้ป่วย" icon={<AirlineSeatFlatAngledIcon />} {...a11yProps(5)} component={RouterLink} to="/rentsearch"/>
                        <Tab label="ระบบค้นหาข้อมูลการทำหัตถการให้ผู้ป่วย" icon={<LocalLibraryIcon />} {...a11yProps(6)} component={RouterLink} to="/operasearch"/>
                        <Tab label="ระบบค้นหาข้อมูลการสั่งยาห้ผู้ป่วยใน" icon={<ColorizeIcon />} {...a11yProps(7)} component={RouterLink} to="/searchpre" />
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