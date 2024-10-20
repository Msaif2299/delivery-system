# Delivery system [UNDER DEVELOPMENT]

A basic delivery system that uses REST API to communicate with an unbuilt frontend system. Made for educational and showcase purposes.
Uses Docker for deployment and MySQL for databases.

## Things to add

The system involves tracking vehicles, managing deliveries, and reacting to real-time traffic updates, while utilizing various technologies like HTTP, WebSockets, Kafka, SQL, NoSQL, and caching.

1. ### User Authentication and Authorization
   Register/Login: Users (e.g., fleet managers, drivers) can create accounts, log in, and access the system.\
   Role-based Access Control (RBAC): Different permissions for fleet managers, drivers, and admins.\
   Fleet Manager: Can manage vehicles, deliveries, routes, and drivers.\
   Driver: Can view their assigned tasks and send real-time status updates.\
   Token-based Authentication: Implement JWT or OAuth for session management and API access control.\
2. ### Vehicle Management
   CRUD Operations for Vehicles:\
   ✅ Create, Read, Update, Delete vehicles from the fleet.\
   ✅ Store vehicle details like type, make, model, capacity, status, etc.\
   ✅ Assign Drivers to Vehicles: Associate a vehicle with a driver for a given period.\
   Real-time Telemetry: Receive and store real-time vehicle data like location, speed, fuel level, tire pressure, engine status via WebSocket or Kafka.\
   Vehicle Status Management: Set and track vehicle statuses (e.g., idle, in-transit, maintenance).\
   Vehicle Availability: Track vehicle availability based on maintenance schedules or task assignments.\
3. ### Driver Management
   CRUD Operations for Drivers:\
   Create, Read, Update, Delete driver profiles.\
   Manage driver information such as license number, contact details, and status (e.g., active, unavailable).\
   Driver Assignment: Assign a driver to a delivery task or vehicle.\
   Driver Route View: Allow drivers to view their assigned routes and task details.\
   Driver Notifications: Send task updates or route changes via WebSocket or messaging queues (Kafka).\
4. ### Delivery Task Management
   Task Creation: Create delivery tasks with detailed information (e.g., pickup/dropoff locations, time windows, cargo details).\
   Task Assignment: Assign tasks to vehicles and drivers.\
   Real-time Task Updates:\
   Track the progress of tasks (pending, en route, completed, canceled).\
   Drivers update task status through a mobile app or GPS device.\
   Task Tracking: Provide real-time tracking URLs for customers to monitor deliveries.\
   Task Priority and Scheduling: Assign priority levels to tasks and schedule deliveries based on priority, vehicle capacity, and driver availability.\
   Task History: Store completed tasks for reporting and analytics.\
5. ### Route Management
   Route Planning: Calculate optimal routes for vehicles based on task locations, real-time traffic data, and road conditions.\
   Dynamic Route Updates: Modify routes in real-time based on traffic incidents (e.g., accidents, congestion) and notify drivers via WebSocket/Kafka.\
   Route Optimization: Use algorithms to dynamically reassign vehicles and optimize routes for efficiency, reducing travel time and fuel consumption.\
   Traffic-aware Routing: Integrate real-time traffic data (e.g., road closures, accidents) into route planning.\
   Route Tracking and Logging: Store actual routes taken by vehicles for future analysis and optimization.\
6. ### Traffic Data Integration
   Real-time Traffic Updates: Continuously receive and process traffic data from external sources (e.g., sensors, traffic APIs) using WebSockets/Kafka.\
   Traffic Incident Management:\
   Detect accidents, congestion, or road closures.\
   Automatically reroute vehicles based on real-time traffic events.\
   Incident Alerts: Notify fleet managers and drivers about traffic incidents affecting their routes.\
   Historical Traffic Data: Store historical traffic patterns to improve future route predictions and task scheduling.\
7. ### Monitoring and Telemetry Dashboard
   Fleet Overview: Provide a real-time dashboard for fleet managers to monitor all active vehicles, tasks, and their statuses.\
   Live Vehicle Tracking: Display vehicle locations on a map with details like speed, route, and ETA.\
   Task Tracking: View all active and upcoming tasks, including assigned vehicles and drivers.\
   Traffic Visualizations: Show traffic conditions (e.g., congestion levels, accidents) on the map in real-time.\
   Alerts and Notifications: Display alerts for important events (e.g., vehicle breakdowns, traffic incidents).\
8. ### Geofencing and Alerts
   Geofence Setup: Create virtual boundaries (geofences) around specific areas (e.g., warehouses, customer locations).\
   Geofence Alerts: Send alerts when a vehicle enters or exits a geofenced area.\
   Idle Alerts: Notify fleet managers when a vehicle has been idle for too long in one location.\
   Speeding Alerts: Trigger notifications when a driver exceeds speed limits.\
9. ### Reporting and Analytics
   Vehicle Performance Reports: Generate reports on vehicle usage, mileage, fuel efficiency, and maintenance history.\
   Driver Performance Reports: Track driver behavior (e.g., speeding, sudden braking) and delivery success rates.\
   Task Completion Reports: Analyze delivery times, delays, and task completion rates.\
   Traffic Pattern Analysis: Analyze historical traffic data to find bottlenecks or common delay points.\
   Cost Efficiency Reports: Analyze fuel consumption, delivery times, and route efficiency for cost-saving insights.\
10. ### Caching and Performance Optimization
    Caching Frequently Accessed Data: Implement caching (e.g., Redis) for frequently accessed data like vehicle locations or traffic updates.\
    Rate Limiting: Apply rate limiting to API endpoints to prevent overloading of the system.\
    Load Balancing: Distribute network traffic across multiple instances of the application to ensure scalability and high availability.\
    Data Sharding: For large datasets (e.g., historical traffic data or telemetry logs), implement sharding for improved database performance.\
11. ### WebSocket/Kafka Integration for Real-time Communication
    Real-time Vehicle Updates: Use WebSockets to push real-time updates of vehicle location, speed, and status to the dashboard and driver apps.\
    Task and Route Notifications: Notify drivers of new or updated tasks/routes using WebSockets or Kafka.\
    Traffic Data Broadcasting: Stream real-time traffic incident data (e.g., accidents, congestion) using Kafka to ensure all relevant systems receive timely updates.\
12. ### NoSQL for High-Volume Data
    Store Real-time Telemetry Data: Use NoSQL databases (e.g., MongoDB, Cassandra) for storing high-volume, real-time data such as vehicle telemetry or traffic data.\
    Unstructured Data Storage: Store unstructured data like sensor readings or traffic camera images for later analysis.\
13. ### Backup and Disaster Recovery
    Database Backup: Implement automated backups for the SQL database (e.g., MySQL) and NoSQL databases (e.g., MongoDB).\
    Failover Mechanism: Set up failover mechanisms to ensure high availability during outages (e.g., using database replication or container orchestration).\
14. ### Maintenance and Health Monitoring
    Vehicle Maintenance Scheduler: Track maintenance schedules and alert fleet managers when vehicles require servicing.\
    System Health Monitoring: Monitor server and application health (e.g., via Prometheus and Grafana) to detect performance issues or downtime.\
    Alerting System: Notify administrators when critical issues occur, such as database downtime, high latency, or network issues.\
