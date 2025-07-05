### App Monitoring 4 golden signals:
* Latency
* Traffic
* Errors
* Saturation (how much system pct is used)

### Monitoring vs Evalution
* Monitoring is continuous process, operational level
* Evaluation long-term process done sometimes, business level to check if useful, sustainable and profitable.

### Monitoring vs Observing
* Monitoring: collecting, aggregating, watching
* Observing: analyzing, recognizing, understanding

### Telemetry vs Monitoring vs Observability
* Telemetry - the continuous collection of metrics, logs, and traces,  
    forms the foundational data pipeline that powers both monitoring and observability.
* Monitoring - uses telemetry to track known conditions and trigger alerts  
    based on predefined metrics, providing a reactive approach to system health.
* Observability - builds on telemetry and monitoring, enabling deeper analysis  
    and correlation of data to diagnose unknown issues and understand complex system behavior.

## Syntetic Monitoring
Behavioral testing, but during continious runtime. 
Sctript pretends to be a user and continiously runs workflow operations.
Used for SLA.

This can helps us to monitor:
* Availability
* Response Time
* Downtime
* Faults

Syntetic Monitoring Tools:
* Datadog, Sematext, SolarWinds Pingdom, App Dynamics, SmartBear AlertSite,
* Dynatrace, New Relic, Size 24x7, AlertBot, Uptrends.

## APM (App perf monitoring)
Functions of APM:
* observation of components
* vis and alerts
* detection of anomalies (ML)
* distributed tracing
* dependency mapping

Telemetry monitored:
* resource utilization
* server logs
* network traffic
* performance metrics

Visualization tools: Kibana, Elastic Search, SysDig, Splunk, Cloud Solution

## Alerts
* Metric Alert
* Log Alert
* Activity Log Alert
* Smart Detection Alert

Alerting tools: Bosun, Cabot, StatsAgg.

## Visualization
Vis tools are powerful not just for presenting data, but for telling a story with it.  
In data storytelling the goal is to turn new numbers into a compelling understandable  
narrative that informs decision making.