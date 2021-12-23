import React, {useEffect, useState } from "react";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import { TenantsInterface } from "../models/ITenant";
import { RentalsInterface } from "../models/IRental";
import { RepairInterface } from "../models/IRepair";
import { RoomsInterface } from "../models/IRoom";
import { RoomEquipmentsInterface } from "../models/IRoomEquipment";
import { TypeEquipmentsInterface } from "../models/ITypeEquipment";
import { Select, Snackbar, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TableSortLabel } from '@material-ui/core';
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import {
    MuiPickersUtilsProvider,
    KeyboardDateTimePicker,
  } from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import { format } from 'date-fns'


const Alert = (props: AlertProps) => {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1,
          },
          container: {
            marginTop: theme.spacing(2),
          },
          paper: {
            padding: theme.spacing(2),
            color: theme.palette.text.secondary,
          },
          table: {
            minWidth: 650,
          },
          tableSpace: {
            marginTop: 20,
          },
    }),
);

function CreateRepair() {
    const classes = useStyles();
    const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
    const [tenants, setTenants] = useState<TenantsInterface>();
    const [rooms, setRooms] = useState<RoomsInterface>();
    const [typeEquipments, setTypeEquipments] = useState<TypeEquipmentsInterface[]>([]);
    const [roomEquipments, setRoomEquipments] = useState<RoomEquipmentsInterface[]>([]);
    const [repairs, setRepairs] = useState<RepairInterface[]>([]);
    const [repair, setRepair] = useState<Partial<RepairInterface>>(
      {}
    );  
    const [rental, setRental] = useState<Partial<RentalsInterface>>(
      {}
    );
    const [typeEquipment, setTypeEquipment] = useState<Partial<TypeEquipmentsInterface>>(
      {}
    );
   
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
    };
    
    const handleChange = (
      event: React.ChangeEvent<{ name?: string; value: any }>
    ) => {
      const name = event.target.name as keyof typeof repair;
      setRepair({
        ...repair,
        [name]: event.target.value,
      });
    };

    const handleDateChange = (date: Date | null) => {
        console.log(date);
        setSelectedDate(date);
      };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
        ) => {
        const id = event.target.id as keyof typeof repair;
        const { value } = event.target;
        setRepair({ ...repair, [id]: value });
    };

    const handleType = (
      event: React.ChangeEvent<{name?: string; value: unknown}>
    ) => {
      getRoomEquipments(Number(event.target.value));
      
    };

    const getRoomEquipments = async (id: number) => {
      let number = localStorage.getItem("room_number");
      fetch(`${apiUrl}/RoomEquipments/room/${number}/type/${id}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
          if(res.data){
            setRoomEquipments(res.data);
          } else {
            console.log("else");
          }
        });
    };

    const getRental = async () => {
        let uid = localStorage.getItem("uid");
        fetch(`${apiUrl}/rentals/${uid}`, requestOptions)
          .then((response) => response.json())
          .then((res) => { 
            if (res.data) {
              localStorage.setItem("room_number", res.data.Room.ROOM_NUMBER);
              setRooms(res.data.Room);
              console.log(res.data.Room);
              
            } else {
              console.log("else");
            }
          });
    };

    const getTenant = async () => {
        let uid = localStorage.getItem("uid");
        fetch(`${apiUrl}/tenant/${uid}`, requestOptions)
          .then((response) => response.json())
          .then((res) => {
            repair.TenantID = res.data.ID;
            if (res.data) {
              setTenants(res.data);
            } else {
              console.log("else");
            }
          });
    };

    const getRepairs = async () => {
      let uid = localStorage.getItem("uid");
      fetch(`${apiUrl}/Repair/${uid}`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
          if (res.data) {
            setRepairs(res.data);
            console.log(res);
          } else {
            console.log("else");
          }
        });
  };

    const getTypeEquipments = async () => {
      fetch(`${apiUrl}/TypeEquipments`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
        if (res.data) {
          setTypeEquipments(res.data);
        } else {
          console.log("else");
        }
        });
    };

    useEffect(() => {
        getTenant();
        getTypeEquipments();
        getRepairs();
        getRental();
    }, []);

  
    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };
    
    function submit() {
        let data = {
            AddedTime: selectedDate,
            Note: repair.Note ?? "",
            TenantID: convertType(tenants?.ID),
            RoomEquipmentID: convertType(repair.RoomEquipmentID),         
        };
        console.log(data)

        const requestOptionsPost = {
          method: "POST",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        };
    
        fetch(`${apiUrl}/Repairs`, requestOptionsPost)
          .then((response) => response.json())
          .then((res) => {
            if (res.data) {
              setSuccess(true);
              window.location.href = "/create_repair";
            } else {
              setError(true);
            }
          });
    }

    return (
        <Container className={classes.container} maxWidth="md">
          <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              บันทึกข้อมูลสำเร็จ
            </Alert>
          </Snackbar>
          <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              บันทึกข้อมูลไม่สำเร็จ
            </Alert>
          </Snackbar>
          <Paper className={classes.paper}>
            <Box display="flex">
              <Box flexGrow={1}>
                <Typography
                  component="h2"
                  variant="h6"
                  color="primary"
                  gutterBottom
                >
                  ระบบแจ้งซ่อม
                </Typography>
              </Box>
            </Box>
            <Divider />
            <Grid container spacing={3} className={classes.root}>
              <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                  <p>เลขที่ห้องพัก</p>
                  <Select
                    native
                    value={rental.ID}
                    onChange={handleChange}
                    disabled
                    inputProps={{
                      name: "RentalID",
                    }}
                  >
                    <option value={rooms?.ID} key={rooms?.ID}>
                      {rooms?.ROOM_NUMBER}
                    </option>
                    
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                  <p>ชื่อผู้แจ้งซ่อม</p>
                  <Select
                    native
                    value={repair.TenantID}
                    onChange={handleChange}
                    disabled
                    inputProps={{
                      name: "TenantID",
                    }}
                  >
                    <option aria-label="None" value=""> 
                      กรุณาเลือกผู้แจ้งซ่อม
                    </option>
              
                    <option value={tenants?.ID} key={tenants?.ID}>
                        {tenants?.Name}
                    </option>
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                  <p>ประเภทอุปกรณ์</p>
                  <Select
                    native
                    value={typeEquipment.ID}
                    onChange={handleType}
                    inputProps={{
                      name: "ID",
                    }}
                  >
                    <option aria-label="None" value="">
                      กรุณาเลือกประเภทอุปกรณ์
                    </option>
                    {typeEquipments.map((item: TypeEquipmentsInterface) => (
                      <option value={item.ID} key={item.ID}>
                        {item.Name}
                      </option>
                    ))}
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                  <p>อุปกรณ์</p>
                  <Select
                    native
                    value={repair.RoomEquipmentID}
                    onChange={handleChange}
                    inputProps={{
                      name: "RoomEquipmentID",
                    }}
                  >
                    <option aria-label="None" value="">
                      กรุณาเลือกอุปกรณ์
                    </option>
                    {roomEquipments.map((item: RoomEquipmentsInterface) => (
                      <option value={item.ID} key={item.ID}>
                        {item.Equipment.Name}
                      </option>
                    ))}
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                  <p>หมายเหตุ</p>
                  <TextField
                    id="Note"
                    variant="outlined"
                    type="string"
                    size="medium"
                    placeholder="กรุณากรอกหมายเหตุ"
                    value={repair.Note || ""}
                    onChange={handleInputChange}
                />
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl fullWidth variant="outlined">
                  <p>วันที่และเวลา</p>
                  <MuiPickersUtilsProvider utils={DateFnsUtils}>
                    <KeyboardDateTimePicker
                      name="WatchedTime"
                      value={selectedDate}
                      onChange={handleDateChange}
                      label="กรุณาเลือกวันที่และเวลา"
                      minDate={new Date("2018-01-01T00:00")}
                      format="yyyy/MM/dd hh:mm a"
                    />
                  </MuiPickersUtilsProvider>
                </FormControl>
              </Grid>
              <Grid item xs={12}>
                <Button
                    style={{ float: "right" }}
                    onClick={submit}
                    variant="contained"
                    color="primary"
                >
                บันทึก
                </Button>
              </Grid>
            </Grid>
          </Paper>
          <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="15%">
                  ผู้แจ้งซ่อม
                </TableCell>
                <TableCell align="center" width="15%">
                  เลขที่ห้องพัก
                </TableCell>
                <TableCell align="center" width="15%">
                  อุปกรณ์
                </TableCell>
                <TableCell align="center" width="20%">
                  หมายเหตุ
                </TableCell>
                <TableCell align="center" width="25%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {repairs.map((item: RepairInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Tenant.Name}</TableCell>
                  <TableCell align="center">{item.Rental.Room.ROOM_NUMBER}</TableCell>
                  <TableCell align="center">{item.RoomEquipment.Equipment.Name}</TableCell>
                  <TableCell align="center">{item.Note}</TableCell>
                  <TableCell align="center">{format((new Date(item.AddedTime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        </Container>
      );
}

export default CreateRepair;
