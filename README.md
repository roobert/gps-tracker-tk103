# gps-tracker-tk103

A receiver and web-ui for data from TK103 based GPS trackers.

## Build

```
git clone git@github.com:roobert/gps-tracker-tk103
cd gps-tracker-tk103
make
```

## Run

Web UI/API
```
./gps-tracker-tk103-ui
```

Data receiver
```
./gps-tracker-tk103-receiver
```

## Example Output

```
$ ./gps-tracker-tk103-receiver
Listening on 0.0.0.0:9000
2017-10-15 20:43:01 <- (handshake) ##,imei:111222333444555,A;
2017-10-15 20:43:01 -> (handshake) LOAD\r\n
2017-10-15 20:43:16 <- (data) imei:111222333444555,tracker,171016034315,,F,194311.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:43:46 <- (data) imei:111222333444555,tracker,171016034345,,F,194341.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:44:01 <- (ping) 111222333444555;
2017-10-15 20:44:01 -> (pong) OK\r\n
2017-10-15 20:44:16 <- (data) imei:111222333444555,tracker,171016034415,,F,194411.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:44:46 <- (data) imei:111222333444555,tracker,171016034445,,F,194441.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:45:01 <- (ping) 111222333444555;
2017-10-15 20:45:01 -> (pong) OK\r\n
2017-10-15 20:45:16 <- (data) imei:111222333444555,tracker,171016034515,,F,194511.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:45:46 <- (data) imei:111222333444555,tracker,171016034545,,F,194541.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:46:01 <- (closing connection) EOF
2017-10-15 20:46:02 <- (handshake) ##,imei:111222333444555,A;
2017-10-15 20:46:02 -> (handshake) LOAD\r\n
2017-10-15 20:46:16 <- (data) imei:111222333444555,tracker,171016034616,,F,194611.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:46:46 <- (data) imei:111222333444555,tracker,171016034646,,F,194641.000,A,1111.2222,N,11111.2222,W,0.00,0;
2017-10-15 20:47:02 <- (ping) 111222333444555;
2017-10-15 20:47:02 -> (pong) OK\r\n
```

## Configure GPS Unit

```
# initialize device
begin123456

# phone device 10 times to set phone number as master number

# add phone number with country code to ensure international incoming calls will be authorized
admin111111 00<number>

# change default password
password123456 111111

# set wap gateway
apn111111 ...

# set wap credentials (optional)
up111111 <username> <password>

# set address for data receiver
adminip111111 <ip> <port>

# enable GPRS, TCP, enable heartbeat
gprs111111,0,0

# always send updates every 30 seconds, regardless of probe states
fix030s030s***n111111

# set timezone to GMT
time zone111111 0
```
