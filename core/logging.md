Why to log?


Spot bugs in the applications code 
Discover performance problems
Do post-mortem analysis of outages and security incidents 

What to log? 

The timestamp for when a event occurred, or a log was generated 
Log levels such as debug , error or info 
Contextual data

- log package 
The Go standard library has a build-in log package that provides most basic logging features
While ti does not have log levels(such as debug, warning or error) it still provides everything you need to get a basic logging strategy set up 

Great for local development when getting fast feedback is more important than generating rich, structured logs 

- Loggin frameworks
Loggin framework helps in standardizing the log data
It's easier to read and understand the log data 
-> popular loggging frameworks 
glog and logrus 
logrus is better maintained and used in popular projects like docker 

