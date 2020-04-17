# Super-Alarm
Go project that generates the fake Super Alarm that will be incorporated into the HASS server.\
It is a vulnerable device that is designed to be exploited in a CTF.\
\
The python script "generateTaffic" will act as the "victim" in this CTF.\
The script first sends a GET request ( to \<alarm address\>\/getCookie ) that will generate a JWT token and set it as a client cookie.\
Then the script will send another GET request every minute to the address ( \<alarm address\>\/toggle ) with the JWT token cookie.\
That is the traffic the hackers must replay. If they do so successfully they will receive the flag.\
\
The cookie generated only lasts 5 minutes, so the python script must be re-called every 5 minutes.
